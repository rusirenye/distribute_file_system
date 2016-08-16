package controllers

import (
	//	"encoding/json"

	//	"distribute_file_system/dao"
	//	"distribute_file_system/log"
	//	"distribute_file_system/models"

	"github.com/astaxie/beego"
)

type NodeController struct {
	beego.Controller
}

func (nc *NodeController) Post{
	host := nc.GetString("host")
	nodelist, err := dao.GetNodeList(host)
}
