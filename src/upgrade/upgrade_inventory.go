package upgrade

import (
	"bytes"
	"io"
	"os"

	"github.com/deployKubernetesInCHINA/dkc-command/src/config"
	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/inventory"
	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"
	"gopkg.in/yaml.v2"
)

// add nfs, ansible_user; delete localhost
func UpgradeInventory() {
	inv := inventory.NewInventory()
	if local, ok := inv.All.Hosts["localhost"]; ok {
		if inv.All.Variables.AnsibleUser == "" {
			inv.All.Variables.AnsibleUser = local.RemoteMachineUsername
			if inv.All.Variables.AnsibleUser == "root" {
				inv.All.Variables.AnsibleBecome = false
				inv.All.Variables.AnsibleBecomePassword = ""
			} else {
				inv.All.Variables.AnsibleBecome = true
				inv.All.Variables.AnsibleBecomePassword = local.RemoteMachinePassword
			}
		}
		if inv.All.Variables.LoginType == "password" {
			inv.All.Variables.AnsiblePassword = local.RemoteMachinePassword
		} else {
			inv.All.Variables.AnsibleSSHkey = local.RemoteSSHkey
		}
		delete(inv.All.Hosts, "localhost")
	}
	if _, ok := inv.All.Children["nfs"]; !ok {
		inv.All.Children["nfs"] = inventory.ChildrenBaseStruct{}
		if inv.All.Children["k8s_cluster"].Vars["csi_driver_nfs_enabled"].(bool) {
			inv.All.Variables.NfsType = "external"
		} else {
			inv.All.Variables.NfsType = "none"
		}
	}
	out, err := yaml.Marshal(inv)
	if err != nil {
		log.Log.Errorln(err.Error())
	}
	inventory.BackUp(config.Kconfig.InventoryFile)
	newfile, _ := os.Create(config.Kconfig.InventoryFile)
	defer newfile.Close()
	io.Copy(newfile, bytes.NewReader(out))

}
