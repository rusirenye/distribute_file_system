package controllers

import (
	//"distribute_file_system/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["json"] = "{\"ObjectId\":\"" + "132" + "\"}"
	c.ServeJSON()
}
