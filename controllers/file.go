package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"strconv"
	"strings"
	"time"

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

func (fc *FileController) Uploadfile() {
	var err string = "0"
	var err_info string = ""

	log.Infof("uploading file")
	file, head, err1 := fc.Ctx.Request.FormFile("file")
	if err1 != nil {
		err = "3003"
		err_info = "file upload failed "
		resObject := &FileRespond{Error: err, Error_info: err_info}
		resStr, _ := json.Marshal(resObject)
		log.Infof(string(resStr))
		fc.Data["json"] = strings.ToLower(string(resStr))
		fc.ServeJSON()
		return
	}
	defer func() {
		file.Close()
	}()
	// file name
	fileName := head.Filename
	// file id
	fileId := fmt.Sprintf("%d", time.Now().Unix())
	// save file
	saveDir := "./static/files/"
	savePath := path.Join(saveDir, fileId)
	log.Infof("uploading file: " + fileId)
	fc.SaveToFile("file", savePath)
	// create tmp directory
	tempFolder, err0 := ioutil.TempDir("./static/files/", "")
	if err0 != nil {
		err_info = "can't create directory :" + saveDir + "_tmp with mode 0777"
		log.Fatal(err_info)

		resObject := &FileRespond{Error: err, Error_info: err_info}
		resStr, _ := json.Marshal(resObject)
		fc.Data["json"] = strings.ToLower(string(resStr))
		fc.ServeJSON()
		return
	}
	time_now := time.Now()
	upload_time := fmt.Sprintf("%d-%s-%02d %02d:%02d:%02d", time_now.Year(), time_now.Month().String(), time_now.Day(), time_now.Hour(), time_now.Minute(), time_now.Second())
	log.Infof(upload_time + " " + fileName + "  " + tempFolder)
	fileData, _ := ioutil.ReadAll(file)

	dataLen := len(fileData)
	blockNum := 0
	if isTrue := (dataLen%(1024*1024*2) == 0); isTrue {
		blockNum = dataLen / (1024 * 1024 * 2)
	} else {
		blockNum = dataLen/(1024*1024*2) + 1
	}
	for i := 0; i < blockNum; i += 1 {
		ioutil.WriteFile(path.Join(tempFolder, strconv.Itoa(i)), fileData[(i*1024*1024*2):((i+1)*1024*1024*2)], 777)
	}
	fc.Data["json"] = "{file:1}"
	fc.ServeJSON()
}
