package cmd

import (
	"github.com/spf13/cobra"

	"github.com/deployKubernetesInCHINA/dkc-command/src/config"
	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg"
	"github.com/deployKubernetesInCHINA/dkc-command/src/prepare"
)

func newPrepareCMD() *cobra.Command {
	var prepareCmd = &cobra.Command{
		Use:   "prepare",
		Short: "检测部署环境信息",
		//Long:  `All software has versions. This is Scheduler's`,
		//TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {
			pkg.Help("prepare")
		},
	}

	var showCmd = &cobra.Command{
		Use:   "show",
		Short: "展示拓扑图",
		//Long:  `All software has versions. This is Scheduler's`,
		//TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {
			config.InitPrepare()
			prepare.Show()
		},
	}
	prepareCmd.AddCommand(showCmd)
	return prepareCmd
}
