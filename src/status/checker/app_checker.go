package checker

import (
	"sort"
	"sync"

	"github.com/deployKubernetesInCHINA/dkc-command/src"
	"github.com/deployKubernetesInCHINA/dkc-command/src/status/checker/pkg/context"
	"github.com/alexeyco/simpletable"
	"github.com/deployKubernetesInCHINA/dkc-command/src/status/checker/pod"
)

type Checker interface {
	checker() bool
}

type checkerFunc func(ctx context.Context)
type AppCheckers map[string]checkerFunc

func NewAppCheckers() *AppCheckers {
	ac := AppCheckers{}
	ac["local_path_provisioner_enabled"] = pod.LocalPathChecker
	ac["csi_driver_nfs_enabled"] = pod.NfsChecker
	ac["elk_enabled"] = pod.ElkChecker
	ac["ingress_nginx_enabled"] = pod.IngressChecker
	ac["minio_enabled"] = pod.MinioChecker
	ac["prometheus_stack_enabled"] = pod.PrometheusChecker
	ac["registry_enabled"] = pod.RegistryChecker

	ac["kube"] = pod.KubeChecker
	return &ac
}

func (ac AppCheckers) Run(ctx context.Context) {

	var sg sync.WaitGroup

	table := src.NewTable("app", "pod", "status", "containers", "containerReady")

	//开关
	for _, v := range []string{"csi_driver_nfs_enabled", "elk_enabled", "minio_enabled", "prometheus_stack_enabled", "registry_enabled"} {
		if kv, ok := ctx.Inventory.All.Children["k8s_cluster"].Vars[v]; ok {
			if kv.(bool) {
				sg.Add(1)
				go func(key string) {
					defer sg.Done()
					ac[key](ctx)
				}(v)
			}
		}
	}
	//常开
	for _, v := range []string{"local_path_provisioner_enabled", "ingress_nginx_enabled", "kube"} {
		sg.Add(1)
		go func(key string) {
			defer sg.Done()
			ac[key](ctx)
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
