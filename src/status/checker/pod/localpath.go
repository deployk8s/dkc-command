package pod

import "github.com/deployKubernetesInCHINA/dkc-command/src/status/checker/pkg/context"

func LocalPathChecker(ctx context.Context) {

	out, err := ctx.Client.CombinedOutput("kubectl get pods -n local-path-storage -l app=local-path-provisioner -o json")
	if err != nil {
		return
	}
	PodIsRunning(ctx, "local-path-provisioner", out)
}
