package pod

import (
	"fmt"
	"strings"

	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"
	"github.com/deployKubernetesInCHINA/dkc-command/src/status/checker/pkg/context"
)

func MinioChecker(ctx context.Context) {
	//
	for _, v := range []string{"app=console", "name=minio-operator"} {
		out, err := ctx.Client.CombinedOutput(fmt.Sprintf("kubectl get pods -n minio-operator -l %s  -o json", v))
		if err != nil {
			log.Log.Errorln(err.Error())
			return
		}
		PodIsRunning(ctx, strings.Split(v, "=")[1], out)
	}
	for _, v := range []string{"app=minio", "v1.min.io/console=minio-console"} {
		out, err := ctx.Client.CombinedOutput(fmt.Sprintf("kubectl get pods -n storage -l %s  -o json", v))
		if err != nil {
			log.Log.Errorln(err.Error())
			return
		}
		PodIsRunning(ctx, strings.Split(v, "=")[1], out)
	}

	return
}
