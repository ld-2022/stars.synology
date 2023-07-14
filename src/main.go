package main

import (
	"fmt"
	"gitee.com/Myzhang/stars.synology/src/action"
	"net/http"
	"strings"
)

func main() {
	Bind("0.0.0.0", "8081")
}

func Bind(host, port string) {
	http.Handle("/", http.FileServer(http.Dir("resources/")))
	http.HandleFunc("/open/v1/status", action.Status)
	fmt.Printf("Starting server at port 8081\n")
	if err := http.ListenAndServe(strings.Join([]string{host, port}, ":"), nil); err != nil {
		panic(err)
	}
}
