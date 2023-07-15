package main

import (
	"fmt"
	"gitee.com/Myzhang/stars.synology/src/action"
	"gitee.com/Myzhang/stars.synology/src/path"
	"net/http"
	"strings"
)

func main() {
	Bind("0.0.0.0", "8081")
}

func Bind(host, port string) {

	fmt.Println("当前运行目录:", path.GetCurrentRunDir())
	fmt.Println("当前运行文件:", http.Dir(path.GetCurrentRunDir()+"/resources/"))
	http.Handle("/", http.FileServer(http.Dir(path.GetCurrentRunDir()+"/resources/")))
	http.HandleFunc("/open/v1/status", action.Status)
	http.HandleFunc("/open/v1/install", action.Install)
	http.HandleFunc("/open/v1/uninstall", action.Uninstall)

	fmt.Printf("Starting server at port 8081\n")
	if err := http.ListenAndServe(strings.Join([]string{host, port}, ":"), nil); err != nil {
		panic(err)
	}
}
