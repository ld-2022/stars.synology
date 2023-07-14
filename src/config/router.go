package config

import (
	"embed"
	"gitee.com/Myzhang/stars.synology/action"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"io/fs"
	"net/http"
	"strings"
)

//go:embed resources
var res embed.FS

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
	e.GET("/*", echo.WrapHandler(http.StripPrefix("/", http.FileServer(getFileSystem()))))
}

func getFileSystem() http.FileSystem {
	fsys, err := fs.Sub(res, "resources")
	if err != nil {
		log.Error(err)
	}
	return http.FS(fsys)
}
