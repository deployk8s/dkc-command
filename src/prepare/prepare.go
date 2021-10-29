package prepare

import (
	"fmt"
	"net"
	"time"

	"github.com/TwinProduction/go-color"
	"github.com/deployKubernetesInCHINA/dkc-command/src"
	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/inventory"
	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"
)

func Show() {
	inventory.CopyInventory()
	i := inventory.NewInventory()
	//t := src.NewTable("Hostname", "Master", "Node", "Mongo", "Log","Monitor","Track", "SG", "IP", "Release(尝试连接服务器)")
	t := src.NewTable("Hostname", "Master", "Node", "OM", "IP", "Release(尝试连接服务器)")
	embarkData(t, i.Hosts)
	t.Show()
	//worker节点 80/443端口检测
	fmt.Println()
	log.Log.Info("node节点80/443 端口检测:")
	for _, v := range i.Hosts {
		if v.IsNode {
			for _, port := range []string{":80", ":443"} {
				_, err := net.DialTimeout("tcp", v.Ip+port, 5*time.Second)
				if err != nil {
					log.Log.Debug(err.Error())
					fmt.Printf("  %s/%s 端口 %s 未被占用, 测试通过\n", v.Hostname, v.Ip, port)
				} else {
					fmt.Println(color.Ize(color.Red, fmt.Sprintf("  %s/%s 端口 %s 被占用,安装可能会失败", v.Hostname, v.Ip, port)))
				}
			}
		}
	}

	//test nas
	fmt.Println()
	if i.All.Variables.NfsType == "internal" {
		if v, ok := i.All.Children["k8s_cluster"].Vars["csi_driver_nfs_enabled"]; ok && v.(bool) {
			//检测端口
			_, err := net.DialTimeout("tcp", i.All.Children["k8s_cluster"].Vars["csi_driver_nfs_server"].(string)+":2049", 5*time.Second)
			if err != nil {
				log.Log.Debug(err.Error())
				log.Log.Infoln("Internal NAS端口检测: 成功")
			} else {
				log.Log.Errorln(color.Ize(color.Red, "Internal NAS端口检测: 失败, 端口2049已被占用,安装可能会失败"))
			}
		}
	} else if i.All.Variables.NfsType == "external" {
		if v, ok := i.All.Children["k8s_cluster"].Vars["csi_driver_nfs_enabled"]; ok && v.(bool) {
			//检测端口
			_, err := net.DialTimeout("tcp", i.All.Children["k8s_cluster"].Vars["csi_driver_nfs_server"].(string)+":2049", 5*time.Second)
			if err != nil {
				log.Log.Debug(err.Error())
				log.Log.Errorln(color.Ize(color.Red, "NAS连接测试: 失败"))
			} else {
				log.Log.Infoln("External NAS连接测试: 成功")
			}
		}
	}
	// test aliyun mongo
	fmt.Println()

}
