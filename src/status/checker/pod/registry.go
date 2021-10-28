package pod

import (
	"fmt"
	"strings"

	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"
	"github.com/deployKubernetesInCHINA/dkc-command/src/status/checker/pkg/context"
)

func RegistryChecker(ctx context.Context) {

	for _, v := range []string{"k8s-app=registry", "app.kubernetes.io/name=chartmuseum", "k8s-app=registry-proxy"} {
		out, err := ctx.Client.CombinedOutput(fmt.Sprintf("kubectl get pods -n registry -l %s  -o json", v))
		if err != nil {
			log.Log.Errorln(err.Error())
			return
		}
		PodIsRunning(ctx, strings.Split(v,"=")[1], out)
	}

	return
}
