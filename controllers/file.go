package controllers

import (
	"encoding/json"
	//	"fmt"
	//	"os"
	"strconv"
	"strings"
	//	"time"

	"distribute_file_system/dao"
	"distribute_file_system/log"
	"distribute_file_system/models"

	"github.com/astaxie/beego"
)

type FileController struct {
	beego.Controller
}
type Files struct {
	hostlist []models.File
}
type FileRespond struct {
	Error      string        `json:"errno"`
	Error_info string        `json:"errmsg"`
	Data       []models.File `json:"data"`
}

// Get return file
func (fc *FileController) Get() {
	//email := cc.GetString("email")
	var file models.File
	file.FileId = fc.GetString("file_id")
	file.Name = fc.GetString("name")
	file.Size, _ = strconv.ParseInt(fc.GetString("size"), 10, 64)
	file.CreateTime = fc.GetString("create_time")
	file.UpdateTime = fc.GetString("update_time")
	file.UploadTime = fc.GetString("upload_time")
	file.Health = fc.GetString("health")

	fileList, err1 := dao.GetFileList(file)
	var err string = "0"
	var err_info string = ""

	if err1 != nil {
		err = "3001"
		err_info = "file select occur problem"
		log.Errorf(err_info)

	}
	if fileList == nil {
		err = "3002"
		err_info = "file list is empty"
		log.Debugf(err_info)
	}
	resObject := &FileRespond{Error: err, Error_info: err_info}
	resObject.Data = fileList[:]
	//	for _, h := range hostlist {

	//	}

	resStr, _ := json.Marshal(resObject)
	log.Infof(string(resStr))
	fc.Data["json"] = strings.ToLower(string(resStr))
	fc.ServeJSON()
}

//func (fc *FileController) uploadFile() {
//	var err string = "0"
//	var err_info string = ""
//	file, head, err1 := fc.GetFile("file")
//	if err1 != nil {
//		err = "3003"
//		err_info = "file upload failed "
//		resObject := &FileRespond{Error: err, Error_info: err_info}
//		resStr, _ := json.Marshal(resObject)
//		log.Infof(string(resStr))
//		fc.Data["json"] = strings.ToLower(string(resStr))
//		fc.ServeJSON()
//		return
//	}
//	// file name
//	name := head.Filename
//	// file id
//	file_id := time.Unix().String()
//	saveDir := "./static/files/"
//	savePath := saveDir + file_id
//	file.Close()
//	// save file
//	fc.SaveToFile("file", savePath)
//	// create tmp directory
//	os.
//	err = os.Mkdir(savePath+"_tmp", 0777)
//	if err != nil {
//		err_info = "can't create directory :" + saveDir + "_tmp with mode 0777"
//		log.Fatal(err_info)
//		resStr, _ := json.Marshal(resObject)
//		resObject := &FileRespond{Error: err, Error_info: err_info}
//		fc.Data["json"] = strings.ToLower(string(resStr))
//		fc.ServeJSON()
//		return
//	}
//	time_now := time.Now()
//	upload_time := fmt.Sprint("%d-%2d-%2d %2d:%2d:%2d", time_now.Year(), time_now.Month().String(), time_now.Day(), time_now.Hour(), time_now.Minute(), time_now.Second())
//}
