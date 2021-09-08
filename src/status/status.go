package status

import (
	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/inventory"
)

func StatusBase(){
	i := inventory.NewInventory()
	showK8sStatus(i)
}

func StatusK8s() {

	i := inventory.NewInventory()
	showK8sStatus(i)
}

