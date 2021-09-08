package inventory

import (
	"fmt"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"

	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"
)

func (h *HostInstance) Connect() error {
	var err error
	var config *ssh.ClientConfig
	if h.Client != nil {
		return err
	}

	if h.Host.Password != ""{
		config = &ssh.ClientConfig{
			Timeout:         1 * time.Second, //ssh 连接time out 时间一秒钟, 如果ssh验证错误 会在一秒内返回
			User:            h.Host.Username,
			Auth:            []ssh.AuthMethod{ssh.Password(h.Host.Password)},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以， 但是不够安全
			//HostKeyCallback: hostKeyCallBackFunc(h.Host),
		}
	}else{

		singer,err := ssh.ParsePrivateKey([]byte(strings.TrimSuffix(h.Host.SSHkey,"\r")))
		if err != nil{
			panic(err)
		}
		config = &ssh.ClientConfig{
			Timeout:         1 * time.Second, //ssh 连接time out 时间一秒钟, 如果ssh验证错误 会在一秒内返回
			User:            h.Host.Username,
			Auth:            []ssh.AuthMethod{ssh.PublicKeys(singer)},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以， 但是不够安全
			//HostKeyCallback: hostKeyCallBackFunc(h.Host),
		}
	}
	h.Client, err = ssh.Dial("tcp", h.Host.Ip+":"+fmt.Sprintf("%d", h.Host.Port), config)
	return err
}

func (h *HostInstance) Close() {
	h.Client.Close()
}

func (h *HostInstance) CombinedOutput(cmd string) (string, error) {
	//创建ssh-session
	log.Log.Debug(cmd)
	session, err := h.Client.NewSession()
	if err != nil {
		log.Log.Fatal("创建ssh session 失败", err.Error())
	}
	defer session.Close()
	//执行远程命令

	combo, err := session.CombinedOutput(cmd)

	if err != nil {
		log.Log.Error("远程执行cmd 失败", err.Error())
		return "", err
	}

	return string(combo), err
}
func (h *HostInstance) Run(cmd string) error {
	//创建ssh-session
	log.Log.Debug(cmd)
	session, err := h.Client.NewSession()
	if err != nil {
		log.Log.Fatal("创建ssh session 失败", err.Error())
	}
	defer session.Close()
	//执行远程命令
	session.Stderr = os.Stdout
	session.Stdout = os.Stdout

	return session.Run(cmd)
}
