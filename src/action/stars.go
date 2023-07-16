package action

import (
	"encoding/json"
	"errors"
	"gitee.com/Myzhang/stars.synology/src/sshx"
	"net/http"
	"strings"
)

var defHost = "192.168.100.107"

func GetConnect(request *http.Request) (conn *sshx.Connection, err error) {
	sshUsername := request.FormValue("ssh_username")
	sshPassword := request.FormValue("ssh_password")
	sshPort := request.FormValue("ssh_port")
	if sshPort == "" {
		sshPort = "22"
	}
	connect, err := sshx.GetConnectNew(sshUsername, sshPassword, defHost, sshPort, []string{})
	if err != nil {
		errStr := err.Error()
		if strings.Contains(errStr, "ssh: unable to authenticate") {
			errStr = "用户名或密码错误"
		} else if strings.Contains(errStr, "connect: connection refused") {
			errStr = "检查：控制面板->终端机和SNMP->终端机->启用SSH功能或端口号是否正确"
		}
		return nil, errors.New(errStr)
	}
	return connect, err
}

func Status(writer http.ResponseWriter, request *http.Request) {
	result := new(R)
	connectNew, err := GetConnect(request)
	if err != nil {
		result.Code = 500
		result.Message = err.Error()
	} else {
		defer connectNew.Close()
		if connectNew.IsInstallStars() {
			result.Code = 0
			if connectNew.IsRunStars() {
				result.Message = "运行中"
			} else {
				result.Message = "未运行"
			}
		} else {
			result.Code = 500
			result.Message = "未安装"
		}
	}
	writeJSON(writer, result)
}

func Install(writer http.ResponseWriter, request *http.Request) {
	result := new(R)
	connectNew, err := GetConnect(request)
	if err != nil {
		result.Code = 500
		result.Message = err.Error()
	} else {
		defer connectNew.Close()
		if connectNew.InstallStars() {
			result.Code = 0
			result.Message = "安装成功"
		} else {
			result.Code = 500
			result.Message = "安装失败"
		}
	}
	writeJSON(writer, result)
}

// 卸载
func Uninstall(writer http.ResponseWriter, request *http.Request) {
	result := new(R)
	connectNew, err := GetConnect(request)
	if err != nil {
		result.Code = 500
		result.Message = err.Error()
	} else {
		defer connectNew.Close()
		if connectNew.UnInstallStars() {
			result.Code = 0
			result.Message = "卸载成功"
		} else {
			result.Code = 500
			result.Message = "卸载失败"
		}
	}
	writeJSON(writer, result)
}

func writeJSON(w http.ResponseWriter, data interface{}) {
	// 设置 Content-Type 为 "application/json"
	w.Header().Set("Content-Type", "application/json")

	// 将数据编码为 json 并写入响应
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		// 如果有错误，设置状态代码为 http.StatusInternalServerError (500)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
