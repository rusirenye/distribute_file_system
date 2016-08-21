package dao

import (
	//	"database/sql"
	"distribute_file_system/log"
	"distribute_file_system/models"

	//"errors"
)

func GetFileList(file models.File) ([]models.File, error) {
	o := GetOrmer()
	sql := `select * from file where 1 = 1 `
	queryParams := make([]interface{}, 1)
	if file.FileId != "" {
		sql += ` and file_id = ? `
		queryParams = append(queryParams, file.FileId)
	}
	if file.Name != "" {
		sql += ` and name = ? `
		queryParams = append(queryParams, file.Name)
	}
	if file.Size != 0 {
		sql += ` and size = ? `
		queryParams = append(queryParams, file.Size)
	}
	if file.CreateTime != "" {
		sql += ` and create_time = ? `
		queryParams = append(queryParams, file.CreateTime)
	}
	if file.UpdateTime != "" {
		sql += ` and update_time = ? `
		queryParams = append(queryParams, file.UpdateTime)
	}
	if file.UploadTime != "" {
		sql += ` and upload_time = ? `
		queryParams = append(queryParams, file.UploadTime)
	}
	if file.Health != "" {
		sql += ` and health = ? `
		queryParams = append(queryParams, file.Health)
	}

	var fileList []models.File
	n, err := o.Raw(sql, queryParams).QueryRows(&fileList)
	log.Debugf("select file num:" + string(n))
	if err != nil {
		return nil, err
	}
	if n == 0 {
		return nil, nil
	}
	return fileList, err
}
