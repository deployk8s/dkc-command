package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/viper"

	"github.com/deployKubernetesInCHINA/dkc-command/src/download"
	"github.com/spf13/cobra"
)

var (
	mirror  string
	onlyOne string
	md5     bool
)
var (
	user         string
	password     string
	fool         bool
	chartRepo    string
	generateOnly bool
	templateDir  string
	opsappsOnly  bool
)

func newDownloadCMD() *cobra.Command {
	var downloadCmd = &cobra.Command{
		Use:              "download",
		Short:            "下载离线文件, 包括docker/ansible/k8s相关安装包",
		TraverseChildren: true,
		Example:          fmt.Sprintf("%s download cache", os.Args[0]),
		//Run: func(cmd *cobra.Command, args []string) {
		//	download.Download()
		//},
	}
	//cobra.OnInitialize(initConfig)
	downloadCmd.PersistentFlags().StringVarP(&onlyOne, "only-one", "o", "", "只下载一个离线文件")

	var downloadCacheCmd = &cobra.Command{
		Use:   "cache",
		Short: "下载离线文件, 包括docker/ansible/k8s相关安装包和ansible剧本",
		//Long:  `All software has versions. This is Scheduler's`,
		TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {
			download.Download()
		},
	}
	downloadCacheCmd.PersistentFlags().StringVarP(&mirror, "mirror", "m", "http://packages.bizconf.cn/packages/offline/", "下载地址 例: http://packages.bizconf.cn/packages/offline/, http://bizconf:bizconf@10.184.101.14:8000")
	downloadCacheCmd.PersistentFlags().BoolVarP(&md5, "check-md5", "c", true, "下载之后检查MD5, 要求:本机可用内存在4G以上.")
	//downloadCacheCmd.PersistentFlags().BoolVarP(&cache, "use-cache", "u", false, "说明: 1.当离线文件压缩包存在时,不进行下载 2. 不存在时,进行下载.之后不删除离线文件压缩包")


	viper.BindPFlags(downloadCacheCmd.PersistentFlags())
	viper.BindPFlags(downloadCmd.PersistentFlags())

	downloadCmd.AddCommand(downloadCacheCmd)
	return downloadCmd
}
