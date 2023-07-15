package action

import (
	"gitee.com/Myzhang/stars.synology/src/command"
	"net/http"
)

func Status(writer http.ResponseWriter, request *http.Request) {
	sshPassword := request.URL.Query().Get("ssh-password")
	err, s := command.SudoExec("stars", sshPassword)
	if err != nil {
		writer.Write([]byte(err.Error() + s))
	}
	writer.Write([]byte(s))
}
