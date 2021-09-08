package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	_ "github.com/deployKubernetesInCHINA/dkc-command/routers"
	"github.com/deployKubernetesInCHINA/dkc-command/src/web"
)

var (
	wFile string
	port  int
)

func newWebCMD() *cobra.Command {
	var webCmd = &cobra.Command{
		Use:   "web",
		Short: "开启web服务, 编辑inventory file, 本地访问: http://localhost:5555",
		//Long:  `All software has versions. This is Scheduler's`,
		//TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {
			web.Web()
		},
	}
	//cobra.OnInitialize(initConfig)
	webCmd.PersistentFlags().IntVarP(&port, "port", "p", 5555, "端口")
	viper.BindPFlag("port", webCmd.PersistentFlags().Lookup("port"))

	return webCmd
}
