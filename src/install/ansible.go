package install

import (
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/TwinProduction/go-color"

	"github.com/deployKubernetesInCHINA/dkc-command/src"
	"github.com/deployKubernetesInCHINA/dkc-command/src/config"
	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"
)

//ansible-playbook -i kubespray/inventory/templates/1master-3worker-3mongo/hosts.yaml kubespray/upgrade-cluster.yml --tags master -v
func runUseAnsible(cmds ...*exec.Cmd) {
	//check sshpass
	// install sshpass

	logfile, err := os.OpenFile("./install.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Log.Fatal(err.Error())
	}
	defer logfile.Close()
	//if config.Kconfig.K8s || config.Kconfig.Mongo {
	var newcmds []*exec.Cmd
	if _, err := exec.LookPath("sshpass"); err != nil {
		var localPath string
		absPath, _ := filepath.Abs(".")
		if src.Release == "7" {
			localPath = filepath.Join(absPath, config.El7DirName, "local/")
		} else if src.Release == "8" {
			localPath = filepath.Join(absPath, config.El8DirName, "local/")
		} else {
			log.Log.Fatal(color.Ize(color.Red, "ansible playbook need sshpass. please install it first."))
		}
		cmd := exec.Command("sh", "-c", "yum install --disablerepo=* -y "+localPath+"/*.rpm")
		log.Log.Info(cmd.String())
		newcmds = append(newcmds, cmd)
	}
	newcmds = append(newcmds, cmds...)
	for _, v := range newcmds {
		v.Stdout = io.MultiWriter(os.Stdout, logfile)
		v.Stderr = io.MultiWriter(os.Stdout, logfile)
		if err := v.Run(); err != nil {
			log.Log.Fatal(err.Error())
		}
	}
}
