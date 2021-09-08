package add_node

import (
	"os"
	"strings"

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
	if v, ok := i.All.Hosts[config.Kconfig.Hostname]; ok {
		if strings.Compare(v.AnsibleHost, config.Kconfig.Ip) != 0 {
			log.Log.Fatalf("hostname %s is exist. and ip address %s conflict", config.Kconfig.Hostname, config.Kconfig.Ip)
		}
	}
	for k, v := range i.All.Hosts {
		if strings.Compare(v.AnsibleHost, config.Kconfig.Ip) == 0 {
			if strings.Compare(k, config.Kconfig.Hostname) != 0 {
				log.Log.Fatalf("Ip %s is exist. and hostname %s is conflict", config.Kconfig.Ip, config.Kconfig.Hostname)
			}
		}
	}

	i.All.Hosts[config.Kconfig.Hostname] = inventory.HostInfo{
		AnsibleHost: config.Kconfig.Ip,
		AnsiblePort: 22,
	}
	i.All.Children["kube_node"].Hosts[config.Kconfig.Hostname] = inventory.Variable{}

	d, _ := yaml.Marshal(i)
	err := os.WriteFile(config.InventoryPath, d, 0644)
	if err != nil {
		log.Log.Fatal(err.Error())
	}
}
