package controllers

import (
	"encoding/json"
	//"strconv"
	"strings"

	"distribute_file_system/dao"
	"distribute_file_system/log"
	"distribute_file_system/models"

	"github.com/astaxie/beego"
)

type BlockController struct {
	beego.Controller
}
type Blocks struct {
	hostlist []models.File
}
type BlockRespond struct {
	Error      string                 `json:"errno"`
	Error_info string                 `json:"errmsg"`
	Data       []models.FileBlockInfo `json:"data"`
}

// Get return block
func (bc *BlockController) GetBlock() {
	type RequestObject struct {
		file_id string
	}
	var block models.Block

	//block = bc.GetString("file_id")
	//	log.Infof(parameters)
	//	jsonByte := []byte(parameters)
	//	var fildIdObject RequestObject

	//	json.Unmarshal(jsonByte, &fildIdObject)
	//	//email := cc.GetString("email")
	//block.FileId = fildIdObject.file_id
	block.FileId = bc.GetString("file_id")
	if block.FileId == "" {
		block.FileId = "1472121418"
	}
	log.Infof(block.FileId + "daffaafadfafasdfaf")
	//	block.BlockNum, _ = strconv.ParseInt(bc.GetString("block_num"), 10, 64)
	//	block.Size, _ = strconv.ParseInt(bc.GetString("size"), 10, 64)
	//	block.NodeIp = bc.GetString("node_id")
	//	block.Health = bc.GetString("health")

	blockList, err1 := dao.GetBlockList(block)
	var err string = "0"
	var err_info string = ""

	if err1 != nil {
		err = "4001"
		err_info = "block select occur problem"
		log.Errorf(err_info)

	}
	if blockList == nil {
		err = "4002"
		err_info = "block list is empty"
		log.Debugf(err_info)
	}
	resObject := &BlockRespond{Error: err, Error_info: err_info}
	resObject.Data = blockList[:]
	//	for _, h := range hostlist {

	//	}

	resStr, _ := json.Marshal(resObject)
	log.Infof(string(resStr))
	bc.Data["json"] = strings.ToLower(string(resStr))
	bc.ServeJSON()
}
