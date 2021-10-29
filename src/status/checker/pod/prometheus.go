package pod

import (
	"fmt"
	"strings"

	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"
	"github.com/deployKubernetesInCHINA/dkc-command/src/status/checker/pkg/context"
)

func PrometheusChecker(ctx context.Context) {

	for _, v := range []string{"app=prometheus", "app=kube-prometheus-stack-operator", "app=prometheus-node-exporter",
		"app=alertmanager", "app=alertmanager-webhook-dingtalk", "app.kubernetes.io/name=grafana", "app.kubernetes.io/name=thanos"} {
		if ctx.OmCount == 1 && v == "app.kubernetes.io/name=thanos" {
			continue
		}
		out, err := ctx.Client.CombinedOutput(fmt.Sprintf("kubectl get pods -n monitoring -l %s  -o json", v))
		if err != nil {
			log.Log.Errorln(err.Error())
			return
		}
		PodIsRunning(ctx, strings.Split(v, "=")[1], out)
	}

	return
}
