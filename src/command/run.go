package command

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
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
	cmd.Stderr = os.Stderr
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return err, ""
	}

	stdoutRead := bufio.NewReader(stdoutPipe)
	message := make([]string, 0)
	go readLine(stdoutRead, func(line string, err error) {
		message = append(message, line)
	})

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("cmd.Run() failed with %s\n", err), ""
	}
	return nil, strings.Join(message, "\n")
}
