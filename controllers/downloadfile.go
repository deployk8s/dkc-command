package controllers

import (
	"net/url"
	"path"

	"github.com/astaxie/beego"

	"github.com/deployKubernetesInCHINA/dkc-command/src/config"
	"github.com/deployKubernetesInCHINA/dkc-command/src/download"
)

type DownloadFileController struct {
	beego.Controller
}

func (c *DownloadFileController) Get() {
	id := c.Ctx.Input.Param(":id")
	fileUrl, _ := url.Parse(config.Kconfig.Mirror)
	fileUrl.Path = id
	filePath := path.Join(config.Kconfig.DownloadDir, id)
	res := NewObjectResp(0, "success", nil)
	if err := download.Downloadfile(fileUrl.String(), filePath); err != nil {
		res.Set(-1, err.Error(), nil)
	}
	c.Data["json"] = &res
	c.ServeJSON()
	//c.TplName = "index.html"
}

func (c *DownloadFileController) Delete() {
	id := c.Ctx.Input.Param(":id")

	bs := NewBaseResp(0, "success")
	if err := download.Delete(id); err != nil {
		bs.Set(-1, err.Error())
	}
	c.Data["json"] = &bs
	c.ServeJSON()
	//c.TplName = "index.html"
}
