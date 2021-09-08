package controllers

import (
	"github.com/astaxie/beego"

	"github.com/deployKubernetesInCHINA/dkc-command/src/download"
)

type FileInfoController struct {
	beego.Controller
}

func (c *FileInfoController) Get() {

	data := download.GetFileInfo()
	res := NewListResp(0, "success", data)

	c.Data["json"] = &res
	c.ServeJSON()
	//c.TplName = "index.html"
}

