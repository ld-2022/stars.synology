package sshx

import (
	"fmt"
	"testing"
)
import _ "golang.org/x/crypto/ssh"

func TestInstallStars(t *testing.T) {
	connect, err := GetConnect("zxw", "QQai,.00", "192.168.100.107", 22, []string{})
	if err != nil {
		t.Error(err)
		return
	}
	if connect.InstallStars() {
		commands, err := connect.SendCommands("sudo stars start")
		if err != nil {
			return
		}
		fmt.Println(string(commands))
		fmt.Println("安装成功")
	} else {
		fmt.Println("安装失败")
	}
}

func TestUnInstallStars(t *testing.T) {
	connect, err := GetConnect("zxw", "QQai,.00", "192.168.100.107", 22, []string{})
	if err != nil {
		t.Error(err)
		return
	}
	if connect.UnInstallStars() {
		fmt.Println("安装成功")
	} else {
		fmt.Println("安装失败")
	}
}
func TestIsRunStars(t *testing.T) {
	connect, err := GetConnect("zxw", "QQai,.00", "192.168.100.107", 22, []string{})
	if err != nil {
		t.Error(err)
		return
	}
	if connect.IsRunStars() {
		fmt.Println("正在运行")
	} else {
		fmt.Println("未运行")
	}
}

func TestStarsIsInstall(t *testing.T) {
	connect, err := GetConnect("zxw", "QQai,.00", "192.168.100.107", 22, []string{})
	if err != nil {
		t.Error(err)
		return
	}
	if connect.IsInstallStars() {
		fmt.Println("已安装")
	} else {
		fmt.Println("未安装")
	}
}
func TestA(t *testing.T) {
	//	connect, err := GetConnect("root", "qqai,.00", "132.232.201.73", 50001, []string{})
	connect, err := GetConnect("zxw", "QQai,.00", "192.168.100.107", 22, []string{})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(connect.IsInstallDocker())
}
