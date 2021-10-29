package add_node

import (
	"io"
	"os"
	"os/exec"

	"github.com/deployKubernetesInCHINA/dkc-command/src"
	"github.com/deployKubernetesInCHINA/dkc-command/src/config"
	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"
)

func runUseAnsible(cmds ...*exec.Cmd) {

	if config.Kconfig.Yes || src.CheckYes("Add k8s node use local ansible.") {
		logfile, err := os.OpenFile("./node-add.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Log.Fatal(err.Error())
		}

		for _, v := range cmds {
			log.Log.Println(" Run CMD: ", v.String())
			v.Stdout = io.MultiWriter(os.Stdout, logfile)
			v.Stderr = io.MultiWriter(os.Stdout, logfile)
			if err := v.Run(); err != nil {
				log.Log.Fatal(err.Error())
			}
		}
	}
}
