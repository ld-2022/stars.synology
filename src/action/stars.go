package action

import (
	"encoding/json"
	"gitee.com/Myzhang/stars.synology/src/command"
	"github.com/ld-2022/jsonx"
	"github.com/shirou/gopsutil/process"
	"net/http"
	"os"
	"strings"
)

func Status(writer http.ResponseWriter, request *http.Request) {
	result := jsonx.NewJSONObject()
	if FileExists("/opt/stars/stars") {
		if ProcessExists("stars") {
			result.Put("status", "正在运行")
		} else {
			result.Put("status", "已安装、未启动")

		}
	} else {
		result.Put("status", "没有安装")
	}
	writeJSON(writer, result)
}

func Install(writer http.ResponseWriter, request *http.Request) {
	sshPassword := request.FormValue("ssh-password")
	err, s := command.SudoExec("curl -O https://download.tbytm.com/stars/releases/shell/linux-install.sh && sudo sh linux-install.sh", sshPassword)
	if err != nil {
		writeJSON(writer, jsonx.NewJSONObject().FluentPut("status", "安装失败").FluentPut("msg", err.Error()))
		return // TODO
	}
	writeJSON(writer, jsonx.NewJSONObject().FluentPut("status", "安装成功").FluentPut("msg", s))
}

// 卸载
func Uninstall(writer http.ResponseWriter, request *http.Request) {
	sshPassword := request.FormValue("ssh-password")
	err, s := command.SudoExec("curl -O https://download.tbytm.com/stars/releases/shell/shell-uninstall.sh && sudo sh shell-uninstall.sh", sshPassword)
	if err != nil {
		writeJSON(writer, jsonx.NewJSONObject().FluentPut("status", "安装失败").FluentPut("msg", err.Error()))
		return // TODO
	}
	writeJSON(writer, jsonx.NewJSONObject().FluentPut("status", "安装成功").FluentPut("msg", s))
}

// 判断指定进程是否存在
func ProcessExists(processName string) bool {
	processes, err := process.Processes()
	if err != nil {
		return false
	}
	for _, p := range processes {
		name, nameErr := p.Name()
		if nameErr != nil {
			continue
		}
		if strings.Contains(name, processName) {
			return true
		}
	}
	return false
}

// 判断文件是否存在
func FileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	} else if err != nil {
		return false
	}
	return true
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
