package main

import "gitee.com/Myzhang/stars.synology/src/config"

func main() {
	config.Bind("0.0.0.0", "8080")
}
