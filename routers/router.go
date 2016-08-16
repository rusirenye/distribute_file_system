package routers

import (
	"distribute_file_system/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}

func initHostRouter() {
	beego.Router("/host", &controllers.HostController{})
}
