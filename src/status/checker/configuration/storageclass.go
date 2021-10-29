package configuration

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alexeyco/simpletable"

	"github.com/deployKubernetesInCHINA/dkc-command/src/config"
	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"
	"github.com/deployKubernetesInCHINA/dkc-command/src/status/checker/pkg"
	"github.com/deployKubernetesInCHINA/dkc-command/src/status/checker/pkg/context"
)

type storageclass struct {
	Items []item
}
type item struct {
	Kind     string
	Metadata metadata
}
type metadata struct {
	Name        string
	Annotations annotation
}
type annotation struct {
	IsDefault string `json:"storageclass.kubernetes.io/is-default-class"`
}

func StorageclassChecker(ctx context.Context) {
	exceptValue := "local-path"
	if v, ok := ctx.Inventory.All.Children["k8s_cluster"].Vars["csi_driver_nfs_enabled"]; ok {
		if v.(bool) {
			exceptValue = "nfs-csi"
		}
	}
	//

	out, err := ctx.Client.CombinedOutput("kubectl get storageclasses.storage.k8s.io -o json")
	if err != nil {
		log.Log.Errorln(err.Error())
		return
	}

	outBytes := bytes.NewBufferString(out)
	var sc storageclass
	err = json.Unmarshal(outBytes.Bytes(), &sc)
	if err != nil {
		log.Log.Errorln(err.Error())
		return
	}
	var value []string
	for _, v := range sc.Items {
		log.Log.Debug("storageclass: ", v.Metadata.Name)
		if v.Metadata.Annotations.IsDefault == "true" {
			value = append(value, v.Metadata.Name)
		}
	}

	if len(value) == 1 && value[0] == exceptValue {
		//正常情况下,开启debug显示
		if config.Kconfig.Show {
			r := []*simpletable.Cell{
				{Align: simpletable.AlignLeft, Text: "storageclass(default)"},
				{Align: simpletable.AlignLeft, Text: exceptValue},
				{Align: simpletable.AlignLeft, Text: exceptValue},
				{Align: simpletable.AlignLeft, Text: "true"},
			}
			ctx.Out.Store("storageclass", r)
		}
	} else {
		//异常情况
		r := []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: "storageclass(default)"},
			{Align: simpletable.AlignLeft, Text: strings.Join(value, ",")},
			{Align: simpletable.AlignLeft, Text: exceptValue},
			{Align: simpletable.AlignLeft, Text: pkg.Red("false")},
		}
		ctx.Out.Store("storageclass", r)
	}

}
