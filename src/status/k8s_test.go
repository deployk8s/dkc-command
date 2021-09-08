package status

import (
	"testing"

	"github.com/deployKubernetesInCHINA/dkc-command/src/config"
	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/inventory"
)

func TestStatusK8s(t *testing.T) {
	config.Kconfig.InventoryFile = "../../inventory/hosts.yaml.bak"
	i := inventory.NewInventory()
	showK8sStatus(i)
}
