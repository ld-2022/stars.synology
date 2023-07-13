package command

import (
	"bufio"
	"context"
	"errors"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os"
	"os/exec"
)

var (
	cmdList []*exec.Cmd
)

type CmdCall func(line string, err error)
type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}

func Run(call CmdCall, name string, arg ...string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return RunCtx(ctx, call, name, arg...)
}
func RunCtx(ctx context.Context, call CmdCall, name string, arg ...string) error {
	command, err := createCmd(ctx, call, name, arg...)
	if err != nil {
		return err
	}
	return command.Run()
}

func Start(call CmdCall, name string, arg ...string) (*exec.Cmd, error) {
	return StartCtx(context.Background(), call, name, arg...)
}
func StartCtx(ctx context.Context, call CmdCall, name string, arg ...string) (*exec.Cmd, error) {
	command, err := createCmd(ctx, call, name, arg...)
	if err != nil {
		return nil, err
	}
	startErr := command.Start()
	return command, startErr
}
func createCmd(ctx context.Context, call CmdCall, name string, arg ...string) (*exec.Cmd, error) {
	command := CommandCtx(ctx, name, arg...)
	stdoutPipe, err := command.StdoutPipe()
	if err != nil {
		return nil, errors.New("获取标准输出管道失败")
	}
	stderrPipe, err := command.StderrPipe()
	if err != nil {
		return nil, errors.New("获取标准错误管道失败")
	}
	stdoutRead := bufio.NewReader(stdoutPipe)
	stderrRead := bufio.NewReader(stderrPipe)
	go readLine(stdoutRead, call)
	go readLine(stderrRead, call)
	return command, nil
}

func ExistFile(str string) bool {
	if _, err := os.Stat(str); err == nil {
		return true
	}
	return false
}

func Close() {
	for _, cmd := range cmdList {
		if err := cmd.Process.Kill(); err != nil {
			return
		}
		err := cmd.Wait()
		if err != nil {
			return
		}
	}
}
