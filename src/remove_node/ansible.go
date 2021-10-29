package remove_node

import (
	"io"
	"os"
	"os/exec"

	"github.com/deployKubernetesInCHINA/dkc-command/src"
	"github.com/deployKubernetesInCHINA/dkc-command/src/config"
	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/inventory"
	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"
)

func removeUseAnsible() {

	var cmds []*exec.Cmd
	resetNode := "reset_nodes=false"
	removal := "allow_ungraceful_removal=true"

	//判断被删除节点的可连通性
	i := inventory.NewInventory()
	for _, v := range i.Hosts {
		if v.Hostname == config.Kconfig.Hostname {
			instance := inventory.NewHostInstance(v)
			err := instance.Connect()
			if err == nil {
				resetNode = "reset_nodes=true"
				removal = "allow_ungraceful_removal=false"
			}
			break
		}
	}
	cmd := exec.Command("ansible-playbook", "-i", config.InventoryPath, "kubespray/remove-node.yml", "-e", "node="+config.Kconfig.Hostname, "-e", resetNode, "-e", removal, "-e", "delete_nodes_confirmation=yes", "-v")
	log.Log.Info(" Run CMD: ", cmd.String())
	cmds = append(cmds, cmd)

	logfile, err := os.OpenFile("./node-del.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Log.Fatal(err.Error())
	}
	if config.Kconfig.Yes || src.CheckYes("Remove k8s node with local ansible 2.9.18 ") {

		for _, v := range cmds {
			v.Stdout = io.MultiWriter(os.Stdout, logfile)
			v.Stderr = io.MultiWriter(os.Stdout, logfile)
			if err := v.Run(); err != nil {
				log.Log.Fatal(err.Error())
			}
		}
	}
}
