package routers

import (
	"github.com/astaxie/beego"

	"github.com/deployKubernetesInCHINA/dkc-command/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/files", &controllers.FileInfoController{})
	beego.Router("/files/?:id", &controllers.DownloadFileController{})
	beego.Router("/inventory", &controllers.InventoryController{})
}
