package upgrade

import (
	"os"
	"os/exec"
	"path"

	"github.com/TwinProduction/go-color"

	"github.com/deployKubernetesInCHINA/dkc-command/src"
	"github.com/deployKubernetesInCHINA/dkc-command/src/config"
	"github.com/deployKubernetesInCHINA/dkc-command/src/download"
	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"
	"github.com/deployKubernetesInCHINA/dkc-command/src/prepare"
)

func UpgradeK8s() {
	for _, v := range config.TarFiles {
		//tarName := strings.Split(v, ".")[0]
		if _, ok := os.Stat(v); ok != nil {
			log.Log.Info(v, " 不存在, 跳过解压")
			continue
		}
		err := download.Untar(v, path.Dir(v))
		if err != nil {
			log.Log.Fatal("解压 ", v, " 失败")
			os.Exit(1)
		}

	}
	if !config.Kconfig.UseDocker {
		if !src.RunAnsible {
			if src.IsSupported() {
				prepare.InstallAnsible()
			} else {
				log.Log.Fatal(color.Ize(color.Red, "Before run command [install], you need install ansible 2.9.18 first"))
			}
		}
		var cmds []*exec.Cmd
		cmd := exec.Command("ansible-playbook", "-i", config.InventoryPath, "kubespray/upgrade-cluster.yml", "-v")
		log.Log.Info(" Run CMD: ", cmd.String())
		cmds = append(cmds, cmd)
		if src.CheckYes("Upgrade k8s version.") {
			runUseAnsible(cmds...)
		}
	} else {
		log.Log.Fatal("can not support command --use-docker")
	}
}
