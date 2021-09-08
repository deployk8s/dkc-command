package remove_node

import (
	"os"

	"gopkg.in/yaml.v2"

	"github.com/deployKubernetesInCHINA/dkc-command/src/config"
	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/inventory"
	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"
)

func updateInventory() {
	//backup inventory
	inventory.BackUp(config.Kconfig.InventoryFile)

	//update hosts.yaml
	i := inventory.NewInventory()
	if _, ok := i.All.Hosts[config.Kconfig.Hostname]; !ok {
		log.Log.Fatalf("hostname %s is not exist.", config.Kconfig.Hostname)
	}

	delete(i.All.Hosts, config.Kconfig.Hostname)

	for k, _ := range i.All.Children {
		delete(i.All.Children[k].Hosts, config.Kconfig.Hostname)
	}
	d, _ := yaml.Marshal(i)
	err := os.WriteFile(config.InventoryPath, d, 0644)
	if err != nil {
		log.Log.Fatal(err.Error())
	}
}
