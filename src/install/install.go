package install

import (
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/TwinProduction/go-color"

	"github.com/deployKubernetesInCHINA/dkc-command/src"
	"github.com/deployKubernetesInCHINA/dkc-command/src/config"
	"github.com/deployKubernetesInCHINA/dkc-command/src/download"
	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/inventory"
	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"
	"github.com/deployKubernetesInCHINA/dkc-command/src/prepare"
)

func InstallAll() {
	inventory.CopyInventory()
	for _, v := range config.TarFiles {
		//tarName := strings.Split(v, ".")[0]
		if _, ok := os.Stat(v); ok != nil {
			log.Log.Info(v, " 不存在, 跳过解压")
			continue
		}
		err := download.Untar(v, path.Dir(v))
		if err != nil {
			log.Log.Fatal("解压 ", v, " 失败")
		}
	}
	var cmds []*exec.Cmd
	if !config.Kconfig.UseDocker {
		if !src.RunAnsible {
			if src.IsSupported() {
				prepare.InstallAnsible()
			} else {
				log.Log.Fatal(color.Ize(color.Red, "Before run command [install], you need install ansible 2.9.18 first"))
			}
		}

		var args []string
		args = append(args, "-i")
		args = append(args, config.InventoryPath)
		args = append(args, "k8s-offline-install/kubespray/prepare.yml")
		args = append(args, "-v")
		cmd := exec.Command("ansible-playbook", args...)

		log.Log.Println(" Run CMD: ", cmd.String())
		cmds = append(cmds, cmd)
		args = []string{}
		args = append(args, "-i")
		args = append(args, config.InventoryPath)
		args = append(args, "k8s-offline-install/kubespray/cluster-offline.yml")
		args = append(args, "-v")
		cmd = exec.Command("ansible-playbook", args...)
		//cmd = exec.Command("ansible-playbook", "-i", config.InventoryPath, "k8s-offline-install/kubespray/cluster-offline.yml", "-v", config.Kconfig.ExtraArgs)
		log.Log.Println(" Run CMD: ", cmd.String())
		cmds = append(cmds, cmd)

		i := inventory.NewInventory()
		if v, ok := i.All.Children["rs1"]; ok {
			if _, _ok := v.Vars["mongodb_enabled"]; _ok {
					args = []string{}
					args = append(args, "-i")
					args = append(args, config.InventoryPath)
					args = append(args, "k8s-offline-install/kubespray/mongodb-install.yml")
					args = append(args, "-v")
					cmd = exec.Command("ansible-playbook", args...)
					//cmd = exec.Command("ansible-playbook", "-i", config.InventoryPath, "k8s-offline-install/kubespray/cluster-offline.yml", "-v", config.Kconfig.ExtraArgs)
					log.Log.Println(" Run CMD: ", cmd.String())
					cmds = append(cmds, cmd)
			}
		}
		if src.CheckYes("Install k8s/mongo with local ansible 2.9.18 ") {
			runUseAnsible(cmds...)
		}
	} else {
		var cmdStr string
		cmdStr = "ansible-playbook -i inventory/hosts.yaml k8s-offline-install/kubespray/prepare.yml -v "
		cmdStr = strings.Join([]string{cmdStr, "ansible-playbook -i inventory/hosts.yaml k8s-offline-install/kubespray/cluster-offline.yml -v " + config.Kconfig.ExtraArgs}, ";")
		runUseDocker(cmdStr)
	}
}
