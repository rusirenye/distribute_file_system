package controllers

import (
	"distribute_file_system/dao"
	"distribute_file_system/log"
	"distribute_file_system/models"
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
)

type NodeController struct {
	beego.Controller
}

type NodeRespond struct {
	Error      string        `json:"errno"`
	Error_info string        `json:"errmsg"`
	Data       []models.Node `json:"data"`
}

func (nc *NodeController) Get() {
	var host models.Host
	host.Ip = nc.GetString("host_ip")
	nodelist, err1 := dao.GetNodeList(host)

	var err string = "0"
	var err_info string = ""

	if err1 != nil {
		err = "2001"
		err_info = "node select occur problem"
		log.Errorf(err_info)

	}
	if nodelist == nil {
		err = "2002"
		err_info = "node list is empty"
		log.Debugf(err_info)
	}
	resObject := &NodeRespond{Error: err, Error_info: err_info}
	resObject.Data = nodelist[:]

	resStr, _ := json.Marshal(resObject)
	log.Infof(string(resStr))
	nc.Data["json"] = strings.ToLower(string(resStr))
	nc.ServeJSON()

}
func (nc *NodeController) Post() {
	nc.Get()
}
