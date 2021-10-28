package pod

import (
	"fmt"
	"strings"

	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"
	"github.com/deployKubernetesInCHINA/dkc-command/src/status/checker/pkg/context"
)

func KubeChecker(ctx context.Context) {

	for _, v := range []string{"component=kube-apiserver", "component=kube-controller-manager", "component=kube-scheduler", "component=etcd",
		"k8s-app=kube-dns", "k8s-app=dns-autoscaler", "k8s-app=kube-proxy", "k8s-app=kube-nginx", "k8s-app=calico-node",
		"k8s-app=calico-kube-controllers", "k8s-app=nodelocaldns", "app.kubernetes.io/name=metrics-server"} {
		out, err := ctx.Client.CombinedOutput(fmt.Sprintf("kubectl get pods -n kube-system -l %s  -o json", v))
		if err != nil {
			log.Log.Errorln(err.Error())
			return
		}

		PodIsRunning(ctx, strings.Split(v, "=")[1], out)
	}

	return
}
