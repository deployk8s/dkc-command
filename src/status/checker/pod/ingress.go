package pod

import (
	"fmt"

	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"
	"github.com/deployKubernetesInCHINA/dkc-command/src/status/checker/pkg/context"
)

func IngressChecker(ctx context.Context) {

	for _, v := range []string{"ingress-nginx"} {
		out, err := ctx.Client.CombinedOutput(fmt.Sprintf("kubectl get pods -n ingress-nginx -l app.kubernetes.io/name=%s  -o json", v))
		if err != nil {
			log.Log.Errorln(err.Error())
			return
		}

		PodIsRunning(ctx, v, out)
	}

	return
}
