package upgrade

import (
	"io"
	"os"
	"os/exec"

	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"
)

func runUseAnsible(cmds ...*exec.Cmd) {

	logfile, err := os.OpenFile("./upgrade.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Log.Fatal(err.Error())
	}
	for _, v := range cmds {
		v.Stdout = io.MultiWriter(os.Stdout, logfile)
		v.Stderr = io.MultiWriter(os.Stdout, logfile)
		if err := v.Run(); err != nil {
			log.Log.Fatal(err.Error())
		}
	}
}
