package config

import (
	"gitee.com/Myzhang/stars.synology/src/action"
	"github.com/labstack/echo/v4"
	"strings"
)

func Bind(host, port string) {
	e := echo.New()
	static(e)
	e.GET("/open/v1/status", action.Status)
	err := e.Start(strings.Join([]string{host, port}, ":"))
	if err != nil {
		e.Logger.Fatal(err)
	}
}

func static(e *echo.Echo) {
	e.Static("/", "resources")
}
