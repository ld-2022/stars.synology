package sshx

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"net"
	"strings"
	"time"
)

type Connection struct {
	*ssh.Client
	password string
}

func GetConnectNew(user, password, host string, port string, cipherList []string) (conn *Connection, err error) {
	if user == "" || password == "" {
		return nil, fmt.Errorf("用户名或密码不能为空")
	}
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		config       ssh.Config
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))
	if len(cipherList) == 0 {
		config = ssh.Config{
			Ciphers: []string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com", "arcfour256", "arcfour128", "aes128-cbc", "3des-cbc", "aes192-cbc", "aes256-cbc"},
		}
	} else {
		config = ssh.Config{
			Ciphers: cipherList,
		}
	}

	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 10 * time.Second,
		Config:  config,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%s", host, port)

	client, err = ssh.Dial("tcp", addr, clientConfig)
	conn = &Connection{Client: client, password: password}
	return
}

func GetConnect(user, password, host string, port int, cipherList []string) (conn *Connection, err error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		config       ssh.Config
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))
	if len(cipherList) == 0 {
		config = ssh.Config{
			Ciphers: []string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com", "arcfour256", "arcfour128", "aes128-cbc", "3des-cbc", "aes192-cbc", "aes256-cbc"},
		}
	} else {
		config = ssh.Config{
			Ciphers: cipherList,
		}
	}

	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 30 * time.Second,
		Config:  config,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	client, err = ssh.Dial("tcp", addr, clientConfig)
	conn = &Connection{Client: client, password: password}
	return
}

func GetSession(user, password, host, key string, port int, cipherList []string) (session *ssh.Session, err error) {
	connect, err := GetConnect(user, password, host, port, cipherList)
	if err != nil {
		return nil, err
	}
	if session, err = connect.NewSession(); err != nil {
		return nil, err
	}
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	if err = session.RequestPty("xterm", 80, 40, modes); err != nil {
		return nil, err
	}
	return session, nil
}

func (conn *Connection) SendCommands(cmdList ...string) ([]byte, error) {
	cmds := []string{"source /etc/profile"}
	cmds = append(cmds, cmdList...)
	session, err := conn.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	err = session.RequestPty("xterm", 80, 40, modes)
	if err != nil {
		return []byte{}, err
	}

	in, err := session.StdinPipe()
	if err != nil {
		return nil, err
	}

	out, err := session.StdoutPipe()
	if err != nil {
		return nil, err
	}

	var output []byte
	rChan := make(chan int, 0)

	go func(in io.WriteCloser, out io.Reader, output *[]byte) {
		var (
			line string
			r    = bufio.NewReader(out)
		)
		for {
			b, err := r.ReadByte()
			if err != nil {
				rChan <- 1
				break
			}

			*output = append(*output, b)

			if b == byte('\n') {
				line = ""
				continue
			}

			line += string(b)
			if (strings.HasPrefix(line, "[sudo] password for ") || strings.HasPrefix(line, "Password: ")) && strings.HasSuffix(line, ": ") {
				_, err = in.Write([]byte(conn.password + "\n"))
				if err != nil {
					rChan <- 1
					break
				}
			}
		}
	}(in, out, &output)

	cmd := strings.Join(cmds, "&& ")
	//fmt.Printf("> %s\n", cmd)
	_, err = session.CombinedOutput(cmd)
	<-rChan
	if err != nil {
		//fmt.Println(string(output))
		return output, err
	}

	return output, nil
}
