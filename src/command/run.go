package command

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
)

type ReadCall func(string, error)

func readLine(reader *bufio.Reader, call ReadCall) {
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

// sudoExec is a function that executes a command as a super user using sudo.
func SudoExec(cmdStr string, sudoPassword string) (error, string) {
	cmd := exec.Command("sudo", "-S", "bash", "-c", cmdStr)
	cmd.Stdin = strings.NewReader(sudoPassword)
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return err, "获取输出管道失败"
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return err, "获取错误管道失败"
	}

	message := make([]string, 0)
	go readLine(bufio.NewReader(stdoutPipe), func(line string, err error) {
		message = append(message, line)
	})
	go readLine(bufio.NewReader(stderrPipe), func(line string, err error) {
		message = append(message, line)
	})

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("cmd.Run() failed with %s\n", err), strings.Join(message, "\n") + "->错误消息"
	}
	return nil, strings.Join(message, "\n") + "->正常消息"
}
