package routers

import (
	"distribute_file_system/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	initHostRouter()
	initNodeRouter()
	initFileRouter()
	initBlockRouter()
}

func initHostRouter() {
	beego.Router("/host", &controllers.HostController{})
}

func initNodeRouter() {
	beego.Router("/node", &controllers.NodeController{})
}

func initFileRouter() {
	beego.Router("/file", &controllers.FileController{})
	//beego.Router("/upload", &controllers.FileController{}, "POST:uploadFile")
}

func initBlockRouter() {
	beego.Router("/block", &controllers.BlockController{})
}
