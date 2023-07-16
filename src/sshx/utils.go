package sshx

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"strings"
)

func (conn *Connection) InstallStars() bool {
	commands, err := conn.SendCommands("curl -O https://download.tbytm.com/stars/releases/shell/nas-linux-install.sh && sudo setsid sh nas-linux-install.sh")
	if err != nil {
		return false
	}
	fmt.Println(string(commands))
	return true
}

func (conn *Connection) UnInstallStars() bool {
	commands, err := conn.SendCommands("curl -O https://download.tbytm.com/stars/releases/shell/nas-shell-uninstall.sh && sudo setsid sh nas-shell-uninstall.sh")
	if err != nil {
		return false
	}
	fmt.Println(string(commands))
	return true
}

func (conn *Connection) IsRunStars() bool {
	// 判断文件是否存在/opt/stars/stars
	commands, err := conn.SendCommands("ps aux | grep -v grep | grep -w \"stars\"")
	if err != nil {
		return false
	}
	if commands == nil || len(commands) == 0 {
		return false
	}
	return true
}
func (conn *Connection) IsInstallStars() bool {
	// 判断文件是否存在/opt/stars/stars
	commands, err := conn.SendCommands("sudo ls /opt/stars/stars")
	if err != nil {
		return false
	}
	if !strings.Contains(string(commands), "/opt/stars/stars") {
		return false
	}
	return true
}
func (conn *Connection) IsInstallDocker() (flag bool, err error) {
	//Docker version 20.10.22, build 3a2c30b
	commands, err := conn.SendCommands("sudo docker -v")
	if err != nil {
		if exitErr, ok := err.(*ssh.ExitError); ok && exitErr.ExitStatus() == 127 {
			return false, nil
		}
		return false, err
	}
	if strings.Contains(strings.TrimSpace(string(commands)), "Docker version") {
		flag = true
	}
	return
}

func (conn *Connection) IsUbuntu() (flag bool, err error) {
	commands, err := conn.SendCommands("lsb_release -a 2> /dev/null | grep \"Distributor ID:\" | cut -d \":\" -f2")
	if err != nil {
		return false, err
	}
	if strings.TrimSpace(string(commands)) == "Ubuntu" {
		flag = true
	}
	return
}

func (conn *Connection) IsCentOS() (flag bool, err error) {
	commands, err := conn.SendCommands("lsb_release -a 2> /dev/null | grep \"Distributor ID:\" | cut -d \":\" -f2")
	if err != nil {
		return false, err
	}
	if strings.TrimSpace(string(commands)) == "CentOS" {
		flag = true
	}
	return
}
