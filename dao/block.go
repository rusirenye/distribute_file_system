package dao

import (
	//	"database/sql"
	"distribute_file_system/log"
	"distribute_file_system/models"

	//"errors"
)

func GetBlockList(block models.Block) ([]models.Block, error) {
	o := GetOrmer()
	sql := `select * from block where 1 = 1 `
	queryParams := make([]interface{}, 1)
	if block.FileId != "" {
		sql += ` and file_id = ? `
		queryParams = append(queryParams, block.FileId)
	}
	var blockList []models.Block
	n, err := o.Raw(sql, queryParams).QueryRows(&blockList)
	log.Debugf("select host lost: num:" + string(n))
	if err != nil {
		return nil, err
	}
	if n == 0 {
		return nil, nil
	}
	return blockList, err
}
