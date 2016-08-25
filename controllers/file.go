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
	"distribute_file_system/utils"

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
	var err_info string = "upload success"

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
		err = "3004"
		err_info = "can't create directory :" + saveDir + "_tmp with mode 0777"
		log.Fatal(err_info)

		resObject := &FileRespond{Error: err, Error_info: err_info}
		resStr, _ := json.Marshal(resObject)
		fc.Data["json"] = strings.ToLower(string(resStr))
		fc.ServeJSON()
		return
	}

	upload_time := fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
	// upload_time := fmt.Sprintf("%d-%s-%02d %02d:%02d:%02d", time_now.Year(), time_now.Month().String(), time_now.Day(), time_now.Hour(), time_now.Minute(), time_now.Second())
	fileData, _ := ioutil.ReadAll(file)
	dataLen := len(fileData)
	// save file info to db
	var fileInfo models.File // = models.File{Id: 0, FileId: fileId, Name: fileName, Size: int64(dataLen), CreateTime: time_now.String(), UpdateTime: time_now.String(), UploadTime: time_now.String(), Health: "1"}
	fileInfo.FileId = fileId
	fileInfo.Name = fileName
	fileInfo.Size = int64(dataLen)
	fileInfo.CreateTime = upload_time
	fileInfo.UpdateTime = upload_time
	fileInfo.UploadTime = upload_time
	fileInfo.Health = "1"
	dao.AddFile(fileInfo)
	log.Infof(upload_time + " " + fileName + "  " + tempFolder)

	blockNum := 0
	if isTrue := (dataLen%(1024*1024*2) == 0); isTrue {
		blockNum = dataLen / (1024 * 1024 * 2)
	} else {
		blockNum = dataLen/(1024*1024*2) + 1
	}
	i := 0
	for i = 0; i < blockNum-1; i += 1 {
		ioutil.WriteFile(path.Join(tempFolder, strconv.Itoa(i)), fileData[(i*1024*1024*2):((i+1)*1024*1024*2)], 777)
	}
	ioutil.WriteFile(path.Join(tempFolder, strconv.Itoa(i)), fileData[(i*1024*1024*2):dataLen], 777)
	utils.SaveBlockToNode(fileId, tempFolder)
	fc.Data["json"] = "{\"errno\":\"0\",\"errmsg\":\"upload file success\"}"
	fc.ServeJSON()
}
