package main

import (
	"distribute_file_system/dao"
	_ "distribute_file_system/routers"
	"distribute_file_system/utils"

	"github.com/astaxie/beego"
)

func main() {

	dao.InitDB()
	utils.InitNodeStorageDir()
	utils.UpdateNodeStatus()
	beego.Run()
}
