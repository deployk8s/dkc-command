package checker

import (
	"sort"
	"sync"

	"github.com/alexeyco/simpletable"

	"github.com/deployKubernetesInCHINA/dkc-command/src"
	"github.com/deployKubernetesInCHINA/dkc-command/src/status/checker/configuration"
	"github.com/deployKubernetesInCHINA/dkc-command/src/status/checker/pkg/context"
)

type ConfigurationCheckers map[string]checkerFunc

func NewConfigurationCheckers() *ConfigurationCheckers {
	cc := ConfigurationCheckers{}
	cc["storageclass"] = configuration.StorageclassChecker
	return &cc
}

func (cc ConfigurationCheckers) Run(ctx context.Context) {
	var sg sync.WaitGroup
	table := src.NewTable("configuration", "value", "exceptValue", "ready")

	for _, v := range []string{"storageclass"} {
		sg.Add(1)
		go func(key string) {
			defer sg.Done()
			cc[key](ctx)
		}(v)
	}

	sg.Wait()

	//sort
	var keys []string
	ctx.Out.Range(func(k, v interface{}) bool {
		keys = append(keys, k.(string))
		return true
	})
	sort.Strings(keys)
	for _, v := range keys {
		cell, ok := ctx.Out.LoadAndDelete(v)
		if ok {
			table.Body.Cells = append(table.Body.Cells, cell.([]*simpletable.Cell))
		}
	}
	table.Show()
}
