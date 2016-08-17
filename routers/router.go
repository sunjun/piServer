package routers

import (
	"github.com/sunjun/piServer/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/device",
			beego.NSInclude(
				&controllers.DeviceController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
