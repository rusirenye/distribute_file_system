package dao

import (
	//	"database/sql"
	"fmt"

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
	log.Debugf("select host lost: num:" + fmt.Sprintf("%s", n))
	if err != nil {
		return nil, err
	}
	if n == 0 {
		return nil, nil
	}
	return blockList, err
}
func GetNodeCandicates(fileId, blockNum string) ([]models.Node, error) {
	o := GetOrmer()
	sql := `SELECT * from node where ip in (SELECT node_ip from block where id not in (SELECT id from block WHERE block_num=? and file_id=?))`
	var nodeList []models.Node
	n, err := o.Raw(sql, blockNum, fileId).QueryRows(&nodeList)
	log.Debugf("select candidate node list num :" + fmt.Sprintf("%s", n))
	if err != nil {
		return nil, err
	}
	if n == 0 {
		return nil, nil
	}
	return nodeList, err
}
func AddBlockToNode(fileId, blockNum, nodeIp string, blockSize int64) error {
	o := GetOrmer()
	sql := `insert into block (file_id, block_num, block_size, block_ip, health) values (?,?,?,?,"1") `
	p, err := o.Raw(sql).Prepare()
	if err != nil {
		return err
	}
	defer p.Close()

	r, err := p.Exec(fileId, blockNum, blockSize, nodeIp)

	if err != nil {
		return err
	}

	if n, err := r.RowsAffected(); n == 0 {
		return err
	}
	return nil
}
