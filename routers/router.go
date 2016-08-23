package routers

import (
	"distribute_file_system/controllers"

	"github.com/astaxie/beego"
)

func init() {
	initIndexRouter()
	initHostRouter()
	initNodeRouter()
	initFileRouter()
	initBlockRouter()
}
func initStaticRouter() {
	beego.SetStaticPath("static", "static")
}
func initIndexRouter() {
	beego.Router("/dfs/", &controllers.MainController{})
}

func initHostRouter() {
	beego.Router("/dfs/host", &controllers.HostController{})
}

func initNodeRouter() {
	beego.Router("/dfs/node", &controllers.NodeController{})
}

func initFileRouter() {
	beego.Router("/dfs/file", &controllers.FileController{})
	beego.Router("/dfs/uploadfile", &controllers.FileController{}, "post:Uploadfile")
}

func initBlockRouter() {
	beego.Router("/dfs/block", &controllers.BlockController{})
}
