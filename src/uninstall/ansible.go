package uninstall

import (
	"io"
	"os"
	"os/exec"

	"github.com/deployKubernetesInCHINA/dkc-command/src"
	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"
)

//ansible-playbook -i kubespray/inventory/templates/1master-3worker-3mongo/hosts.yaml kubespray/upgrade-cluster.yml --tags master -v
func runUseAnsible(cmds []*exec.Cmd) {
	if !src.CheckYes("Uninstall " + os.Args[2]+" with local ansible 2.9.18 ") {
		os.Exit(0)
	}
	logfile, err := os.OpenFile("./uninstall.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Log.Fatal(err.Error())
	}
	defer logfile.Close()
	for _, v := range cmds {
		v.Stdout = io.MultiWriter(os.Stdout, logfile)
		v.Stderr = io.MultiWriter(os.Stdout, logfile)
		if err := v.Run(); err != nil {
			log.Log.Fatal(err.Error())
		}
	}

}
