package controllers

import (
	"encoding/json"
	
	"distribute_file_system/dao"
	"distribute_file_system/log"
	"distribute_file_system/models"
	
	"github.com/astaxie/beego"
)

type HostController struct {
	beego.Controller
}
type Hosts struct {
	hostlist []modles.Host
}
type Respond struct {
	Error      string    `json:"errno"`
	Error_info string `json:"errmsg"`
	Data       Hosts  `json:"data"`
}

// Get return host list
func (c *HostController) Get() {
	//email := cc.GetString("email")
	hostlist, err := dao.GetHostList()
	var err  string = "0"
	var err_info string = ""
	
	if err != nil {
 		err = "1001"
		err_info = "host select occur problem"
		log.Errorf(err_info)
		
	}
	if hostlist == nil {
		err = "1002"
		err_info = "host list is empty"
		log.Debugf(err_info)
	}
	resObject := Respond{
		Error : err,
		Error_info : err_info,
		Data : hostlist
	}
	resStr, _ := json.Marshal(resObject)
	c.Data["json"] = resStr
	c.ServeJSON()
}
