package sshx

import (
	"golang.org/x/crypto/ssh"
	"strings"
)

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
