package pod

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

type PodStatus struct {
	ApiVersion string `json:"apiVersion"`
	Items      []Item
}
type Metadata struct {
	Labels Label  `json:"labels"`
	Name   string `json:"name"`
}
type Label struct {
	App  string `json:"app"`
	App2 string `json:"app.kubernetes.io/name"`
}
type Item struct {
	Kind     string   `json:"kind"`
	Status   Status   `json:"status"`
	Metadata Metadata `json:"metadata"`
}

type Status struct {
	ContainerStatuses []ContainerStatus `json:"containerStatuses"`
	Phase             string            `json:"phase"`
}

type ContainerStatus struct {
	Name    string `json:"name"`
	Ready   bool   `json:"ready"`
	Started bool   `json:"started"`
}

func PodIsRunning(ctx context.Context, appName, outString string) {
	outBytes := bytes.NewBufferString(outString)
	var ps PodStatus
	err := json.Unmarshal(outBytes.Bytes(), &ps)
	if err != nil {
		log.Log.Errorln(err.Error())
		return
	}
	//var app string
	var podName []string
	var phase []string
	var containerNames []string
	var containerStatuses []string

	// 判断是否有数据返回

	if len(ps.Items) == 0 {
		appName = pkg.Red(appName + " (not installed)")
		r := []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: appName, Span: 5},
		}
		ctx.Out.Store(appName, r)
		//ctx.Table.Body.Cells = append(ctx.Table.Body.Cells, r)
		return
	}
	show := config.Kconfig.Show // 正常不显示
	for _, itemV := range ps.Items {
		//if itemV.Metadata.Labels.App == "" {
		//	app = itemV.Metadata.Labels.App2
		//}else{
		//	app = itemV.Metadata.Labels.App
		//}
		podName = append(podName, itemV.Metadata.Name)
		if itemV.Status.Phase != "Running" {
			phase = append(phase, pkg.Red(itemV.Status.Phase))
			show = true
		} else {
			phase = append(phase, itemV.Status.Phase)
		}

		for i, v := range itemV.Status.ContainerStatuses {
			var ready string
			if v.Ready {
				ready = "true"
			} else {
				ready = pkg.Red("false")
				show = true
			}
			containerNames = append(containerNames, v.Name)
			containerStatuses = append(containerStatuses, ready)
			//从i > 0开始 podName 和phase 补充\n
			if i > 0 {
				podName = append(podName, "")
				phase = append(phase, "")
			}
		}

	}
	if show {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: appName},
			{Align: simpletable.AlignLeft, Text: strings.Join(podName, "\n")},
			{Align: simpletable.AlignLeft, Text: strings.Join(phase, "\n")},
			{Align: simpletable.AlignLeft, Text: strings.Join(containerNames, "\n")},
			{Align: simpletable.AlignLeft, Text: strings.Join(containerStatuses, "\n")},
		}
		ctx.Out.Store(appName, r)
		//ctx.Table.Body.Cells = append(ctx.Table.Body.Cells, r)
	}
	return
}
