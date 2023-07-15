package sshx

import (
	"fmt"
	"testing"
)
import _ "golang.org/x/crypto/ssh"

func TestReInstall(t *testing.T) {
	//UnInstallNode("root", "Cat@02001", "39.96.162.105", 22)
	//InstallNode("root", "Cat@02001", "39.96.162.105", 22, 3344, 1122, "123456789")

	//UnInstallNode("ubuntu", "qqai,.00", "82.157.122.2", 22)
	//InstallNode("ubuntu", "qqai,.00", "82.157.122.2", 22, 3344, 1122, "123456789")
	//
	UnInstallNode("ubuntu", "qqai,.00", "101.42.225.236", 22)
	InstallNode("ubuntu", "qqai,.00", "101.42.225.236", 22, 6262, 1122, "")

}
func TestA(t *testing.T) {
	//	connect, err := GetConnect("root", "qqai,.00", "132.232.201.73", 50001, []string{})
	connect, err := GetConnect("zxw", "1314qqai", "172.1.1.4", 22, []string{})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(connect.IsInstallDocker())
}
func TestIsInstallNode(t *testing.T) {
	isInstallNode, err := IsInstallNode("zxw", "1314qqai", "172.1.1.4", 22)
	if err != nil {
		t.Error(err)
		return
	}
	if isInstallNode {
		fmt.Println("已经安装")
	} else {
		fmt.Println("未安装")
	}
}
func TestInstallNode(t *testing.T) {
	err := InstallNode("zxw", "1314qqai", "172.1.1.4", 22, 10000, 9999, "123456789")
	if err != nil {
		t.Error(err)
	}
}

func TestUnInstallNode(t *testing.T) {
	err := UnInstallNode("zxw", "1314qqai", "172.1.1.4", 22)
	if err != nil {
		t.Error(err)
	}
}
