package main

import (
	"net/http"
)

func main() {
	// 静态文件服务器，用于提供前端的HTML、CSS和JavaScript文件。
	fs := http.FileServer(http.Dir("ui"))

	// 注册路由
	http.Handle("/", fs)

	// 启动Web服务器
	http.ListenAndServe(":8081", nil)
}
