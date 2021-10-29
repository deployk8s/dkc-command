package pod

import (
	"fmt"

	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"
	"github.com/deployKubernetesInCHINA/dkc-command/src/status/checker/pkg/context"
)

func ElkChecker(ctx context.Context) {

	for _, v := range []string{"elasticsearch-master", "elasticsearch-data", "filebeat-filebeat", "kibana", "logstash-logstash"} {
		out, err := ctx.Client.CombinedOutput(fmt.Sprintf("kubectl get pods -n logging -l app=%s  -o json", v))
		if err != nil {
			log.Log.Errorln(err.Error())
			return
		}

		PodIsRunning(ctx, v, out)
	}

	return
}
