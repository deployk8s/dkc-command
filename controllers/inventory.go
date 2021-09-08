package controllers

import (
	"encoding/json"
	"io/ioutil"

	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/inventory"
	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"

	"github.com/astaxie/beego"
)

type InventoryController struct {
	beego.Controller
}

func (c *InventoryController) Get() {
	var ir inventory.InventoryData
	i := inventory.NewInventory()

	ir.Ingress = i.All.Variables.ExternalDomainName
	ir.Remote.Username = i.RemoteUsername
	ir.Remote.Password = i.RemotePassword
	ir.Hosts = i.Hosts

	res := NewObjectResp(0, "success", ir)

	c.Data["json"] = &res
	c.ServeJSON()
}
func (c *InventoryController) Post() {
	var ir inventory.InventoryData
	res := NewObjectResp(0, "success", nil)

	if d, err := ioutil.ReadAll(c.Ctx.Request.Body); err != nil {
		log.Log.Println(err.Error())
		res.Set(-1, err.Error(), nil)
	} else {
		if err := json.Unmarshal(d, &ir); err != nil {
			res.Set(-1, err.Error(), nil)
		} else {
			err = inventory.SaveFile(ir)
			if err != nil {
				res.Set(-1, err.Error(), nil)
			} else {
				var ir inventory.InventoryData
				i := inventory.NewInventory()

				ir.Ingress = i.All.Variables.ExternalDomainName
				ir.Remote.Username = i.RemoteUsername
				ir.Remote.Password = i.RemotePassword
				ir.Hosts = i.Hosts
				res.Set(0, "success", ir)
			}

		}
	}

	c.Data["json"] = &res
	c.ServeJSON()
}
