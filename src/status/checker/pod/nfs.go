package pod

import (
	"fmt"

	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"
	"github.com/deployKubernetesInCHINA/dkc-command/src/status/checker/pkg/context"
)

func NfsChecker(ctx context.Context) {

	for _, v := range []string{"csi-nfs-node", "csi-nfs-controller"} {
		out, err := ctx.Client.CombinedOutput(fmt.Sprintf("kubectl get pods -n nfs -l app=%s  -o json", v))
		if err != nil {
			log.Log.Errorln(err.Error())
			return
		}

		PodIsRunning(ctx, v, out)
	}

	return
}
