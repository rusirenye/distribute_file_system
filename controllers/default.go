package controllers

import (
	"distribute_file_system/dao"
	"distribute_file_system/log"
	"distribute_file_system/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	h, err := dao.GetHost(models.Host{Ip: "1"})
	if err != nil {
		log.Errorf("no host")
	}
	c.Data["json"] = "{\"ObjectId\":\"" + h.Ip + "\"}"
	c.ServeJSON()
}
