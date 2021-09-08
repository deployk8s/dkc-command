package remove_node

import (
	"fmt"
	"strings"

	"github.com/TwinProduction/go-color"

	"github.com/deployKubernetesInCHINA/dkc-command/src"
	"github.com/deployKubernetesInCHINA/dkc-command/src/config"
	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/inventory"
	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"
	"github.com/deployKubernetesInCHINA/dkc-command/src/prepare"
)

func RemoveNode() {
	config.InitRemoveNode()
	inventory.CopyInventory()

	if err := checkNode(); err != nil {
		log.Log.Fatal(err.Error())
	}
	if !config.Kconfig.OnlyUpdateInventory {

		if !config.Kconfig.UseDocker {
			if !src.RunAnsible {
				if src.IsSupported() {
					prepare.InstallAnsible()
				} else {
					log.Log.Fatal(color.Ize(color.Red, "Before run command [remove-node], you need install ansible 2.9.18 first"))
				}
			}
			removeUseAnsible()
		} else {
			log.Log.Fatal("can not support command --use-docker")
		}
	}
	updateInventory()
	//update inventory file
}
func checkNode() error {
	i := inventory.NewInventory()
	if _, ok := i.All.Hosts[config.Kconfig.Hostname]; !ok {
		return fmt.Errorf("hostname %s is not exist.", config.Kconfig.Hostname)
	}
	// is master
	for _, v := range i.GetMaster() {
		if strings.Compare(v, config.Kconfig.Hostname) == 0 {
			return fmt.Errorf("hostname %s is k8s master, cannot remove.", config.Kconfig.Hostname)
		}
	}

	flag := false
	for _, v := range i.GetNode() {
		if strings.Compare(v, config.Kconfig.Hostname) == 0 {
			flag = true
		}
	}
	if !flag {
		return fmt.Errorf("hostname %s is not k8s node, cannot remove.", config.Kconfig.Hostname)
	}
	return nil
}
