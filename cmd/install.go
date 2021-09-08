package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/deployKubernetesInCHINA/dkc-command/src/config"
	"github.com/deployKubernetesInCHINA/dkc-command/src/install"
)

var (
	//useDocker bool
	extraArgs string
)

func newInstallCMD() *cobra.Command {
	var installCmd = &cobra.Command{
		Use:   "install",
		Short: "默认使用本地ansible 2.9.18安装k8s",
		Long:
		`默认使用本地ansible 2.9.18安装k8s`,
		TraverseChildren: true,
		//Run: func(cmd *cobra.Command, args []string) {
		//	pkg.Help("install")
		//},
	}

	//installCmd.PersistentFlags().BoolVar(&useDocker, "use-docker", false, "强制使用docker进行安装")

	viper.BindPFlags(installCmd.PersistentFlags())

	var k8sCommand = &cobra.Command{
		Use:   "k8s",
		Short: "默认使用本地ansible 2.9.18安装k8s集群",
		Long:
		`默认使用本地ansible 2.9.18安装k8s集群`,
		//TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {
			config.InitInstall()
			install.InstallK8s()
		},
	}
	k8sCommand.PersistentFlags().StringVarP(&extraArgs, "extra-args", "e", "", "将参数传输给cluster-offline.yml剧本,比如指定--extra-args \"--tags elk\"")
	viper.BindPFlags(k8sCommand.PersistentFlags())

	installCmd.AddCommand(k8sCommand)
	return installCmd
}
