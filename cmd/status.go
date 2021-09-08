package cmd

import (
	"github.com/deployKubernetesInCHINA/dkc-command/src/config"
	"github.com/deployKubernetesInCHINA/dkc-command/src/status"

	"github.com/spf13/cobra"
)

func newStatusCMD() *cobra.Command {
	var statusCmd = &cobra.Command{
		Use:   "status",
		Short: "检查k8s集群状态",
		//Long:  `All software has versions. This is Scheduler's`,
		//TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {
			config.InitStatus()
			status.StatusBase()
		},
	}

	var statusK8sCmd = &cobra.Command{
		Use:   "k8s",
		Short: "检查k8s集群状态",
		//Long:  `All software has versions. This is Scheduler's`,
		//TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {
			config.InitStatus()
			status.StatusK8s()
		},
	}

	statusCmd.AddCommand(statusK8sCmd)
	return statusCmd
}
