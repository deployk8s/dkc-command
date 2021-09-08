package cmd

import (
	"github.com/deployKubernetesInCHINA/dkc-command/src/config"
	"github.com/deployKubernetesInCHINA/dkc-command/src/uninstall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	uninstalluseDocker bool
)

func newUninstallCMD() *cobra.Command {
	var uninstallCmd = &cobra.Command{
		Use:   "uninstall",
		Short: "默认使用本地ansible 2.9.18删除k8s集群",
		//Long:  `All software has versions. This is Scheduler's`,
		TraverseChildren: true,
		//Run: func(cmd *cobra.Command, args []string) {
		//	pkg.Help("uninstall")
		//},
	}

	uninstallCmd.PersistentFlags().BoolVar(&uninstalluseDocker, "use-docker", false, "强制使用docker删除k8s/mongo")

	viper.BindPFlag("uninstall.use-docker", uninstallCmd.PersistentFlags().Lookup("use-docker"))

	var uninstallk8sCmd = &cobra.Command{
		Use:   "k8s",
		Short: "默认使用本地ansible 2.9.18删除k8s",
		//Long:  `All software has versions. This is Scheduler's`,
		//TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {
			config.InitUninstall()
			uninstall.UninstallK8s()
		},
	}


	uninstallCmd.AddCommand(uninstallk8sCmd)
	return uninstallCmd
}
