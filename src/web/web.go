package web

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"

	"github.com/deployKubernetesInCHINA/dkc-command/src/config"
)

func Web() {
	config.InitWeb()
	//download.UpdateMirror()
	beego.BConfig.Listen.HTTPPort = config.Kconfig.Port
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		//允许访问所有源
		AllowAllOrigins: true,
		//可选参数"GET", "POST", "PUT", "DELETE", "OPTIONS" (*为所有)
		//其中Options跨域复杂请求预检
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		//指的是允许的Header的种类
		AllowHeaders: []string{"Content-Type","*"},
		//公开的HTTP标头列表
		ExposeHeaders: []string{"Content-Length"},
		//如果设置，则允许共享身份验证凭据，例如cookie
		AllowCredentials: true,
	}))
	beego.SetStaticPath("/images","static/images")
	beego.SetStaticPath("/css","static/css")
	beego.SetStaticPath("/js","static/js")
	beego.SetStaticPath("/fonts","static/fonts")
	beego.Run()
}
