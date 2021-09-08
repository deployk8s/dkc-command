package prepare

import (
	"github.com/deployKubernetesInCHINA/dkc-command/src"
	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/inventory"
)

func Show() {
	inventory.CopyInventory()
	i := inventory.NewInventory()
	//t := src.NewTable("Hostname", "Master", "Node", "Mongo", "Log","Monitor","Track", "SG", "IP", "Release(尝试连接服务器)")
	t := src.NewTable("Hostname", "Master", "Node", "Mongo", "Log", "Monitor", "SG", "IP", "Release(尝试连接服务器)")
	embarkData(t, i.Hosts)
	t.Show()
}
