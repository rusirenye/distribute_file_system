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

func UpdateNodeStatus(node models.Node) error {
	o := GetOrmer()
	queryParams := make([]interface{}, 1)
	sql := ` update node set brandwidth = ?,brandwidth_used = ?,
		     disk = ?,disk_used = ?, cpu = ?,cpu_used = ?,memory = ?,memory_used = ?,
			 host_ip = ?, health = ? `
	queryParams=append(queryParams,node.Brandwidth,node.BrandwidthUsed,node.Disk,node.DiskUsed,node.Cpu,node.CpuUsed
	            ,node.Memory,node.MemoryUsed,node.HostIp,node.Health)
	r, err := o.Raw(sql, queryParams).Exec()
	if err != nil {
		return err
	}
	if _, err := r.RowsAffected(); err != nil {
		return err
	}
	return nil
}

// get node standard status value 
func GetStandardNodeList() ([]models.Node, error) {
	o := GetOrmer()
	sql := ` select * from node_s `
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
