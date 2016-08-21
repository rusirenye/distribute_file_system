package dao

import (
	//	"database/sql"
	"distribute_file_system/log"
	"distribute_file_system/models"

	//"errors"
)

func GetNodeList(host models.Host) ([]models.Node, error) {
	o := GetOrmer()
	sql := `select * from node where 1 = 1 `
	queryParams := make([]interface{}, 1)
	if host.Ip != "" {
		sql += ` and host_ip = ? `
		queryParams = append(queryParams, host.Ip)
	}
	var nodeList []models.Node
	n, err := o.Raw(sql, queryParams).QueryRows(&nodeList)
	log.Debugf("select host lost: num:" + string(n))
	if err != nil {
		return nil, err
	}
	if n == 0 {
		return nil, nil
	}
	return nodeList, err
}
