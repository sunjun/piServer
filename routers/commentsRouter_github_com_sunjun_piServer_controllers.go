package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/sunjun/piServer/controllers:DeviceController"] = append(beego.GlobalControllerRouter["github.com/sunjun/piServer/controllers:DeviceController"],
		beego.ControllerComments{
			"Status",
			`/status`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/sunjun/piServer/controllers:DeviceController"] = append(beego.GlobalControllerRouter["github.com/sunjun/piServer/controllers:DeviceController"],
		beego.ControllerComments{
			"Group",
			`/group`,
			[]string{"get"},
			nil})

}
