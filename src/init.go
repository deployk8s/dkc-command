package src

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/go-ini/ini"

	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"
)

var RunDocker, RunAnsible, IsWindows, IsDarwin, IsLinux bool
var Release string

func init() {
	if _, err := exec.LookPath("ansible"); err == nil {
		cmd := exec.Command("ansible", "--version")
		if output, err := cmd.Output(); err != nil {
			log.Log.Debug(err.Error())
		} else {
			if strings.Contains(string(output), "ansible 2.9.18") {
				RunAnsible = true
				log.Log.Debug("ansible 2.9.18")
			} else {
				log.Log.Debug(strings.Split(string(output), "\n")[0])
			}
		}
	} else {
		//log.Log.Debug(color.Ize(color.Red, "ansible 2.9.18 not found"))
	}

	if _, err := exec.LookPath("docker"); err == nil {
		cmd := exec.Command("docker", "-v")
		if output, err := cmd.Output(); err != nil {
			log.Log.Debug(err.Error())
		} else {
			RunDocker = true
			log.Log.Debug(string(output))
		}
	} else {
		//log.Log.Debug(color.Ize(color.Red, "docker not found"))
	}
	//if !RunAnsible && !RunDocker {
	//	log.Log.Debug(color.Ize(color.Red, "Warning, command [install] is not supported on this host, "+
	//		"to fix this you need install docker or ansible 2.9.18,"+
	//		" Run: prepare --ansible or --docker"))
	//}

	if runtime.GOOS == "linux" {
		cfg, err := ini.Load("/etc/os-release")
		if err != nil {
			log.Log.Debug("Fail to get os release ", err.Error())
		} else {
			// detect local host is  centos7 or centos8
			ConfigParams := make(map[string]string)
			ConfigParams["ID"] = cfg.Section("").Key("ID").String()
			if ConfigParams["ID"] == "centos" {
				Release = cfg.Section("").Key("VERSION_ID").String()
			}
		}
	}
}
