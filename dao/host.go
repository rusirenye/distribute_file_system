package dao

import (
	"database/sql"
	"errors"

	"distribute_file_system/log"
	"distribute_file_system/models"
)

func GetHostList() ([]models.Host, error) {
	o := GetOrmer()
	sql := `select * from host`
	var hostList []models.host
	n, err := o.Raw(sql).QueryRows(&hostList)
	log.Debugf("select host lost: num:" + n)
	if err != nil {
		return nil,err
	}
	if n=={
		return nil,nil
	}
	return &hostList, nil
}

func GetHost(query models.Host) (*models.Host, error) {

	o := GetOrmer()

	sql := `select * from host `
	// queryParam := make([]interface{}, 1)
	// if query.UserID != 0 {
	// 	sql += ` and user_id = ? `
	// 	queryParam = append(queryParam, query.UserID)
	// }

	// if query.Username != "" {
	// 	sql += ` and username = ? `
	// 	queryParam = append(queryParam, query.Username)
	// }

	// if query.ResetUUID != "" {
	// 	sql += ` and reset_uuid = ? `
	// 	queryParam = append(queryParam, query.ResetUUID)
	// }

	var host []models.Host
	n, err := o.Raw(sql).QueryRows(&host)

	if err != nil {
		return nil, err
	}
	if n == 0 {
		return nil, nil
	}

	return &host[0], nil
}
