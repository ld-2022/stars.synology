package command

import (
	"bufio"
	"fmt"
	"github.com/labstack/gommon/log"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	err := Run(func(line string, err error) {
		fmt.Println(line)
	}, "ping", "192.168.1.1")
	if err != nil {
		t.Error(err)
	}
}
func TestStart(t *testing.T) {
	startTime := time.Now().UnixMilli()

	c, err := Start(func(line string, err error) {
		fmt.Println(line)
		if err != nil {
			fmt.Println(err)
		}
	}, "ping", "192.168.1.1")
	err = c.Wait()
	endTime := time.Now().UnixMilli() - startTime
	fmt.Println("耗时:", endTime)
	if err != nil {
		t.Error(err)
	}
}

func TestCommand_Run(t *testing.T) {
	err, message := sudoExec(`echo 1sss23 > /var/root/test.txt && ls /var/root/`, "123456789..")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}

// Handler type for handling each line of the output
type Handler func(line string, err error)

// sudoExec is a function that executes a command as a super user using sudo.
func sudoExec(cmdStr string, sudoPassword string) (error, string) {
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
