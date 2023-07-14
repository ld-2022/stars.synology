package action

import (
	"gitee.com/Myzhang/stars.synology/src/command"
	"github.com/labstack/echo"
)

func Status(c echo.Context) error {
	sshPassword := c.QueryParam("ssh-password")
	err, s := command.SudoExec("stars", sshPassword)
	if err != nil {
		return c.JSON(200, err.Error())
	}
	return c.JSON(200, s)
}
