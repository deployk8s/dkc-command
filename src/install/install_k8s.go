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

func InstallK8s() {
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
			os.Exit(1)
		}

	}
	var cmds []*exec.Cmd
	var cmd *exec.Cmd
	if !config.Kconfig.UseDocker {
		if !src.RunAnsible {
			if src.IsSupported() {
				prepare.InstallAnsible()
			} else {
				log.Log.Fatal(color.Ize(color.Red, "Before run command [install], you need install ansible 2.9.18 first"))
			}
		}

		extraArgs := strings.Split(config.Kconfig.ExtraArgs, " ")
		var args []string
		//extraArgs 为空, 安装所有; 否则只安装extra
		if config.Kconfig.ExtraArgs == "" {
			args = append(args, "-i")
			args = append(args, config.InventoryPath)
			args = append(args, "kubespray/prepare.yml")
			args = append(args, "-v")
			cmd = exec.Command("ansible-playbook", args...)
			log.Log.Info("Run CMD: ", cmd.String())
			cmds = append(cmds, cmd)
		}

		args = []string{}
		args = append(args, "-i")
		args = append(args, config.InventoryPath)
		args = append(args, "kubespray/cluster-offline.yml")
		args = append(args, "-v")
		for _, v := range extraArgs {
			if v != "" {
				args = append(args, v)
			}
		}
		cmd = exec.Command("ansible-playbook", args...)
		//cmd = exec.Command("ansible-playbook", "-i", config.InventoryPath, "k8s-offline-install/kubespray/cluster-offline.yml", "-v", config.Kconfig.ExtraArgs)
		log.Log.Info("Run CMD: ", cmd.String())
		cmds = append(cmds, cmd)
		if config.Kconfig.Yes || src.CheckYes("Install " + os.Args[2] + " with local ansible 2.9.18 ") {
			runUseAnsible(cmds...)
		}
	} else {
		var cmdStr string
		cmdStr = "ansible-playbook -i inventory/hosts.yaml k8s-offline-install/kubespray/prepare.yml -v "
		cmdStr = strings.Join([]string{cmdStr, "ansible-playbook -i inventory/hosts.yaml k8s-offline-install/kubespray/cluster-offline.yml -v " + config.Kconfig.ExtraArgs}, ";")
		runUseDocker(cmdStr)
	}
}
