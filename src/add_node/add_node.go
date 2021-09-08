package add_node

import (
	"net"
	"os/exec"

	"github.com/TwinProduction/go-color"

	"github.com/deployKubernetesInCHINA/dkc-command/src"
	"github.com/deployKubernetesInCHINA/dkc-command/src/config"
	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/inventory"
	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"
	"github.com/deployKubernetesInCHINA/dkc-command/src/prepare"
)

func AddNode() {
	config.InitAddNode()
	inventory.CopyInventory()

	if net.ParseIP(config.Kconfig.Ip) == nil {
		log.Log.Fatalf("ip address %s is Invalid", config.Kconfig.Ip)
	}
	updateInventory()

	//是否只进行更新inventory操作
	if !config.Kconfig.OnlyUpdateInventory {

		if !config.Kconfig.UseDocker {
			if !src.RunAnsible {
				if src.IsSupported() {
					prepare.InstallAnsible()
				} else {
					log.Log.Fatal(color.Ize(color.Red, "Before run command [add-node], you need install ansible 2.9.18 first"))
				}
			}
			var cmds []*exec.Cmd
			cmd := exec.Command("ansible-playbook", "-i", config.InventoryPath, "kubespray/passwordless.yml", "-v")
			log.Log.Println(cmd.String())
			cmds = append(cmds, cmd)

			cmd = exec.Command("ansible-playbook", "-i", config.InventoryPath, "kubespray/prepare.yml", "--limit="+config.Kconfig.Hostname, "-v")
			log.Log.Println(cmd.String())
			cmds = append(cmds, cmd)

			cmd = exec.Command("ansible-playbook", "-i", config.InventoryPath, "kubespray/facts.yml", "-v")
			log.Log.Println(cmd.String())
			cmds = append(cmds, cmd)

			cmd = exec.Command("ansible-playbook", "-i", config.InventoryPath, "kubespray/scale.yml", "--tags", "etchosts", "-v")
			log.Log.Println(cmd.String())
			cmds = append(cmds, cmd)

			cmd = exec.Command("ansible-playbook", "-i", config.InventoryPath, "k8s-offline-install/kubespray/scale.yml", "--limit="+config.Kconfig.Hostname, "-v")
			log.Log.Println(cmd.String())
			cmds = append(cmds, cmd)
			cmd = exec.Command("ansible-playbook", "-i", config.InventoryPath, "k8s-offline-install/kubespray/cluster-offline.yml", "--limit="+config.Kconfig.Hostname, "--tags", "download", "-v")
			log.Log.Println(cmd.String())
			cmds = append(cmds, cmd)
			runUseAnsible(cmds...)
		}
		//update inventory file
	}

}
