package main

import "gitee.com/Myzhang/stars.synology/config"

func main() {
	config.Bind("0.0.0.0", "8080")
}
