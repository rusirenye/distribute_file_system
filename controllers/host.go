package controllers

import (
	"encoding/json"
	"strings"

	"distribute_file_system/dao"
	"distribute_file_system/log"
	"distribute_file_system/models"

	"github.com/astaxie/beego"
)

type HostController struct {
	beego.Controller
}
type Hosts struct {
	hostlist []models.Host
}
type Respond struct {
	Error      string        `json:"errno"`
	Error_info string        `json:"errmsg"`
	Data       []models.Host `json:"data"`
}

// Get return host list
func (c *HostController) Get() {
	//email := cc.GetString("email")
	hostlist, err1 := dao.GetHostList()
	var err string = "0"
	var err_info string = ""

	if err1 != nil {
		err = "1001"
		err_info = "host select occur problem"
		log.Errorf(err_info)

	}
	if hostlist == nil {
		err = "1002"
		err_info = "host list is empty"
		log.Debugf(err_info)
	}
	resObject := &Respond{Error: err, Error_info: err_info}
	resObject.Data = hostlist[:]
	//	for _, h := range hostlist {

	//	}

	resStr, _ := json.Marshal(resObject)
	log.Infof(string(resStr))
	c.Data["json"] = strings.ToLower(string(resStr))
	c.ServeJSON()
}
