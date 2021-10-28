package context

import (
	"sync"

	"github.com/deployKubernetesInCHINA/dkc-command/src"
	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/inventory"
)

type Context struct {
	Client    inventory.HostInstance
	Inventory inventory.Inventory
	Table     *src.Table
	OmCount   int
	Out       *sync.Map
}
