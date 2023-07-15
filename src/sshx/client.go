package sshx

import (
	"errors"
	"log"
	"strings"
)

var (
	// 单机命令
	runCMD = "sudo docker run -d --restart=always --privileged --net=host --name stars.client -p 7725:7725/tcp xianwei2022/stars.client:latest"
)

func IsInstallClient(username, password, host string, port int) (flag bool, err error) {
	connect, err := GetConnect(username, password, host, port, []string{})
	if err != nil {
		return false, err
	}
	if installDocker, installDockerErr := connect.IsInstallDocker(); installDockerErr != nil {
		return false, installDockerErr
	} else if !installDocker {
		return false, errors.New("没有安装Docker")
	}
	commands, err := connect.SendCommands("sudo docker ps -a")
	if err != nil {
		return false, err
	}
	for _, line := range strings.Split(string(commands), "\n") {
		if strings.Contains(line, "xianwei2022/stars.client:latest") || strings.Contains(line, "stars.client") {
			flag = true
			break
		}
	}
	return
}

func IsRunClient(username, password, host string, port int) (flag bool, err error) {
	connect, err := GetConnect(username, password, host, port, []string{})
	if err != nil {
		return false, err
	}
	if installDocker, installDockerErr := connect.IsInstallDocker(); installDockerErr != nil {
		return false, installDockerErr
	} else if !installDocker {
		return false, errors.New("没有安装Docker")
	}
	commands, err := connect.SendCommands("sudo docker ps -a")
	if err != nil {
		return false, err
	}
	for _, line := range strings.Split(string(commands), "\n") {
		if strings.Contains(line, "xianwei2022/stars.client:latest") {
			if !strings.Contains(line, "Exited") {
				flag = true
			}
			break
		}
	}
	return
}
func InstallClient(username, password, host string, port int) error {
	connect, err := GetConnect(username, password, host, port, []string{})
	if err != nil {
		return err
	}
	if installDocker, installDockerErr := connect.IsInstallDocker(); installDockerErr != nil {
		return installDockerErr
	} else if !installDocker {
		return errors.New("没有安装Docker")
	}

	xxx, insmodErr := connect.SendCommands("sudo /sbin/insmod /lib/modules/tun.ko")
	if insmodErr != nil {
		if strings.Contains(string(xxx), "insmod: ERROR: could not insert module") && strings.Contains(string(xxx), "File exists") {
			log.Println("已加载驱动!")
		} else {
			log.Println(string(xxx))
			return insmodErr
		}
	}

	// 1：加载驱动、2：删除自启动脚本、3：创建自启动脚本、4：设置自启动脚本权限
	_, installTupErr := connect.SendCommands(
		"sudo rm -rf /usr/local/etc/rc.d/tun.sh",
		"sudo bash -c \"echo 'insmod /lib/modules/tun.ko' >> /usr/local/etc/rc.d/tun.sh\"",
		"sudo chmod a+x /usr/local/etc/rc.d/tun.sh")
	if installTupErr != nil {
		return installTupErr
	}

	_, commandsErr := connect.SendCommands("sudo docker pull xianwei2022/stars.client:latest", runCMD)
	if commandsErr != nil {
		return commandsErr
	}
	return nil
}

func UnInstallClient(username, password, host string, port int) error {
	connect, err := GetConnect(username, password, host, port, []string{})
	if err != nil {
		return err
	}
	if installDocker, installDockerErr := connect.IsInstallDocker(); installDockerErr != nil {
		return installDockerErr
	} else if !installDocker {
		return errors.New("没有安装Docker")
	}
	_, commandsErr := connect.SendCommands("sudo docker stop stars.client", "sudo docker rm stars.client")
	if commandsErr != nil {
		return commandsErr
	}
	return nil
}

func StopClient(connect *Connection) error {
	if installDocker, installDockerErr := connect.IsInstallDocker(); installDockerErr != nil {
		return installDockerErr
	} else if !installDocker {
		return errors.New("没有安装Docker")
	}

	_, commandsErr := connect.SendCommands("sudo docker stop stars.client")
	if commandsErr != nil {
		return commandsErr
	}
	return nil
}
func StartClient(connect *Connection) error {
	if installDocker, installDockerErr := connect.IsInstallDocker(); installDockerErr != nil {
		return installDockerErr
	} else if !installDocker {
		return errors.New("没有安装Docker")
	}

	_, commandsErr := connect.SendCommands("sudo docker start stars.client")
	if commandsErr != nil {
		return commandsErr
	}
	return nil
}

func ReStartClient(connect *Connection) error {
	err := StopClient(connect)
	if err != nil {
		return err
	}
	return StartClient(connect)
}
