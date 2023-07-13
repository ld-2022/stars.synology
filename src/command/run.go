package command

import (
	"bufio"
	"context"
	"io"
	"log"
	"os/exec"
	"syscall"
)

func readLine(reader *bufio.Reader, call CmdCall) {
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("读取输出出错: ", err)
			break
		}
		if call != nil {
			call(string(line), err)
		}
	}
}
func Command(name string, arg ...string) *exec.Cmd {
	cmd := exec.Command(name, arg...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Foreground: false,
		Setsid:     true,
	}
	return cmd
}

func CommandCtx(ctx context.Context, name string, arg ...string) *exec.Cmd {
	cmd := exec.CommandContext(ctx, name, arg...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Foreground: false,
		Setsid:     true,
	}
	return cmd
}
