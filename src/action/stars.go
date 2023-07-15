package action

import (
	"encoding/json"
	"gitee.com/Myzhang/stars.synology/src/sshx"
	"github.com/ld-2022/jsonx"
	"net/http"
)

func Status(writer http.ResponseWriter, request *http.Request) {
	result := jsonx.NewJSONObject()
	sshUsername := request.FormValue("ssh_username")
	sshPassword := request.FormValue("ssh_password")
	sshPort := request.FormValue("ssh_port")
	if sshPort == "" {
		sshPort = "22"
	}
	connectNew, err := sshx.GetConnectNew(sshUsername, sshPassword, "127.0.0.1", sshPort, []string{})
	if err != nil {
		result.Put("status", err.Error())
	} else {
		if connectNew.IsInstallStars() {
			result.Put("status", "已安装")
		} else {
			result.Put("status", "没有安装")
		}
	}
	writeJSON(writer, result)
}

func Install(writer http.ResponseWriter, request *http.Request) {
	result := jsonx.NewJSONObject()
	sshUsername := request.FormValue("ssh_username")
	sshPassword := request.FormValue("ssh_password")
	sshPort := request.FormValue("ssh_port")
	if sshPort == "" {
		sshPort = "22"
	}
	connectNew, err := sshx.GetConnectNew(sshUsername, sshPassword, "127.0.0.1", sshPort, []string{})
	if err != nil {
		result.Put("status", err.Error())
	} else {
		if connectNew.InstallStars() {
			result.Put("status", "安装成功")
		} else {
			result.Put("status", "安装失败")
		}
	}
	writeJSON(writer, result)
}

// 卸载
func Uninstall(writer http.ResponseWriter, request *http.Request) {
	result := jsonx.NewJSONObject()
	sshUsername := request.FormValue("ssh_username")
	sshPassword := request.FormValue("ssh_password")
	sshPort := request.FormValue("ssh_port")
	if sshPort == "" {
		sshPort = "22"
	}
	connectNew, err := sshx.GetConnectNew(sshUsername, sshPassword, "127.0.0.1", sshPort, []string{})
	if err != nil {
		result.Put("status", err.Error())
	} else {
		if connectNew.UnInstallStars() {
			result.Put("status", "卸载成功")
		} else {
			result.Put("status", "卸载失败")
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
