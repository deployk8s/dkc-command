package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/deployKubernetesInCHINA/dkc-command/src/add_node"
	"github.com/deployKubernetesInCHINA/dkc-command/src/remove_node"
)

var (
	hostname              string
	ip                    string
	only_update_inventory bool
	//resetNode             bool
)

func newNodeCMD() *cobra.Command {
	var command = &cobra.Command{
		Use:   "node",
		Short: "增/删k8s node节点",
		//Long:  `All software has versions. This is Scheduler's`,
		//TraverseChildren: true,
		//Run: func(cmd *cobra.Command, args []string) {
		//	add_node.AddNode()
		//},
	}
	var commandAdd = &cobra.Command{
		Use:   "add",
		Short: "增加k8s node节点,不支持批量增加",
		Long:  fmt.Sprintf(`增加k8s node节点,不支持批量增加. 示例: %s node add --hostname dkc-worker-2 --ip 192.168.57.4`, os.Args[0]),
		//TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {
			add_node.AddNode()
		},
	}
	command.PersistentFlags().StringVar(&hostname, "hostname", "", "增加节点的hostname, 例: dkc-worker-2")
	command.PersistentFlags().BoolVarP(&only_update_inventory, "only-update-inventory", "o", false, "只更新inventory/hosts.yaml, 不进行k8s add/del node实际操作")
	commandAdd.PersistentFlags().StringVar(&ip, "ip", "", "增加节点的ip地址, 例: 192.168.57.4")

	var commandRemove = &cobra.Command{
		Use:   "del",
		Short: "删除k8s node节点,不支持批量删除",
		Long:  fmt.Sprintf(`删除k8s node节点,不支持批量删除. 示例: %s node del --hostname dkc-worker-2`, os.Args[0]),
		//TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {
			remove_node.RemoveNode()
		},
	}
	//commandRemove.PersistentFlags().BoolVarP(&resetNode, "remove-data", "r", true, "清除删除节点上的安装文件")

	viper.BindPFlags(commandAdd.PersistentFlags())
	viper.BindPFlags(commandRemove.PersistentFlags())
	viper.BindPFlags(command.PersistentFlags())
	//cobra.OnInitialize(initConfig)
	command.AddCommand(commandAdd, commandRemove)
	return command
}
