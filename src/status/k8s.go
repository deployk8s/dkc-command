package status

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/TwinProduction/go-color"

	"github.com/deployKubernetesInCHINA/dkc-command/src"
	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/inventory"
	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"
)

func newEtcdTable() *src.Table {
	return src.NewTable("Member", "Healthy")
}

type address struct {
	Address string `json:"address"`
	Type    string `json:"type"`
}
type condition struct {
	Status  string `json:"status"`
	Type    string `json:"type"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

type status struct {
	Addresses  []address   `json:"addresses"`
	Conditions []condition `json:"conditions"`
}

func showK8sStatus(i *inventory.Inventory) {
	for _, v := range i.Hosts {
		if v.IsMaster {
			instance := inventory.NewHostInstance(v)
			err := instance.Connect()
			if err != nil {
				log.Log.Error(v.Hostname, "connect err.")
				continue
			} else {
				defer instance.Close()
				// get nodes
				log.Log.Info("K8S Cluster:")
				instance.Run("kubectl get nodes -o wide")

				//conditions
				log.Log.Info("Nodes status:")
				var st []status
				out, _ := instance.CombinedOutput("kubectl get nodes -o jsonpath='{\"[\"}{range .items[*]}{\"{\"}{\"\\\"addresses\\\":\"}{.status.addresses}{\",\"}{\"\\\"conditions\\\":\"}{.status.conditions}{\"}\"}{\",\"}{end}{\"]\"}'")
				out = strings.Replace(out, "},]", "}]", 1) //修正json格式
				outBytes := bytes.NewBufferString(out)
				json.Unmarshal(outBytes.Bytes(), &st)
				for _, v := range st {
					flag := false
					hostname := ""
					var strs []string
					for _, va := range v.Addresses {
						if va.Type == "Hostname" {
							hostname = va.Address
						}
					}
					for _, vc := range v.Conditions {
						for _, vt := range []string{"NetworkUnavailable", "MemoryPressure", "DiskPressure", "PIDPressure"} {
							if vc.Type == vt && vc.Status == "True" {
								strs = append(strs, fmt.Sprintf("    %s", vc.Type))
								flag = true
							}
						}
						for _, vt := range []string{"Ready"} {
							if vc.Type == vt && vc.Status == "False" {
								strs = append(strs, fmt.Sprintf("    Node not %s", vc.Type))
								flag = true
							}
						}
					}
					if !flag {
						fmt.Println("  ",hostname, "healthy")
					} else {
						fmt.Println(color.Ize(color.Red,fmt.Sprintf("  %s %s",hostname, "is unhealthy")))
						for _, s := range strs {
							fmt.Println(s)
						}
					}
				}
				// helm release status
				log.Log.Info("Helm release List:")
				instance.Run("helm ls --all-namespaces")
				// pod failed status
				log.Log.Info("Not Running Pods List:")
				instance.Run("kubectl get pods --all-namespaces |grep -v Running")
				break
			}
		}
	}
	//etcd status
	//node status
	//pod status
}
