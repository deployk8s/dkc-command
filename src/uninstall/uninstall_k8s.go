package uninstall

import (
	"os/exec"

	"github.com/TwinProduction/go-color"

	"github.com/deployKubernetesInCHINA/dkc-command/src"
	"github.com/deployKubernetesInCHINA/dkc-command/src/config"
	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/inventory"
	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"
)

func UninstallK8s() {

	//copy hosts.yaml 到inventory目录下
	inventory.CopyInventory()
	if src.RunAnsible && !config.Kconfig.UseDocker {
		cmd := exec.Command("ansible-playbook", "-i", config.InventoryPath, "kubespray/reset.yml", "-e", "reset_confirmation=yes", "-v")
		log.Log.Info(" Run CMD: ", cmd.String())
		runUseAnsible([]*exec.Cmd{cmd})
	} else if src.RunDocker && config.Kconfig.UseDocker {
		cmdStr := "ansible-playbook -i inventory/hosts.yaml kubespray/reset.yml -e reset_confirmation=yes -v"
		log.Log.Info(" Run CMD: ", cmdStr)
		runUseDocker(cmdStr)
	} else {
		log.Log.Fatal(color.Ize(color.Red, "Before run command [uninstall], you need install ansible 2.9.18 first,"))
	}
}
