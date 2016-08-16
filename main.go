package main

import (
	"distribute_file_system/dao"

	_ "distribute_file_system/routers"

	"github.com/astaxie/beego"
)

func main() {

	dao.InitDB()
	beego.Run()
}
