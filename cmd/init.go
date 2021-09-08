package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/deployKubernetesInCHINA/dkc-command/src/config"
)

var log bool
var iFile string

var rootCmd = &cobra.Command{
	Use:     fmt.Sprintf("%s [subcommand]", os.Args[0]),
	Short:   "kubernetes 离线安装助手",
	Version: config.Version,
	// Run: func(cmd *cobra.Command, args []string) {
	//      v, _ := cmd.PersistentFlags().GetBool("version")
	//      fmt.Println(v, args)
	//      // Do Stuff Here
	// },

}

// Execute
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// rootCmd.PersistentFlags().BoolVarP(&version, "version", "v", false, "Print the version number of scheduler")
	rootCmd.PersistentFlags().BoolVar(&log, "debug", false, "show debug logs")
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
	rootCmd.PersistentFlags().StringVarP(&iFile, "inventory-file", "i", config.InventoryPath, "拓扑文件hosts.yaml")
	viper.BindPFlag("inventory-file", rootCmd.PersistentFlags().Lookup("inventory-file"))

	cobra.OnInitialize()
	rootCmd.AddCommand(newDownloadCMD())
	rootCmd.AddCommand(newPrepareCMD())
	rootCmd.AddCommand(newInstallCMD())
	rootCmd.AddCommand(newWebCMD())
	rootCmd.AddCommand(newStatusCMD())
	rootCmd.AddCommand(newUninstallCMD())
	rootCmd.AddCommand(newNodeCMD())
}
