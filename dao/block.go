package dao

import (
	//	"database/sql"
	"fmt"

	"distribute_file_system/log"
	"distribute_file_system/models"

	//"errors"
)

func GetBlockList(block models.Block) ([]models.FileBlockInfo, error) {
	type NumT struct {
		Num int `orm:"column(num)"`
	}
	type IpListT struct {
		IpList string `orm:"column(node_ip)"`
	}
	o := GetOrmer()
	sql := `select max(block_num) as "num" from block where 1 = 1 `
	queryParams := make([]interface{}, 1)
	if block.FileId != "" {
		sql += ` and file_id = ? `
		queryParams = append(queryParams, block.FileId)
	}
	var maxNum []NumT
	n, err := o.Raw(sql, queryParams).QueryRows(&maxNum)
	log.Infof(fmt.Sprintf("max block num : %d", maxNum[0].Num))

	sql = `select * from block where file_id = ? and block_num = ?`
	var fileBlockInfList []models.FileBlockInfo
	var fileBlockInf models.FileBlockInfo
	for i := 0; i <= maxNum[0].Num; i++ {
		var blockInf []models.Block
		n, err = o.Raw(sql, block.FileId, i).QueryRows(&blockInf)
		var ips []string
		for _, b := range blockInf {
			ips = append(ips, b.NodeIp)
		}
		fileBlockInf.FileId = blockInf[0].FileId
		fileBlockInf.BlockNum = blockInf[0].BlockNum
		fileBlockInf.Size = blockInf[0].Size
		fileBlockInf.Health = blockInf[0].Health
		fileBlockInf.NodeIpList = ips
		fileBlockInfList = append(fileBlockInfList, fileBlockInf)
	}

	//	FileId     string   `json:"file_id"`
	//	BlockNum   int64    `json:"block_num"`
	//	Size       int64    `json:"block_size"`
	//	NodeIpList []string `json:"node_ip_list"`
	//	Health     string   `json:"health"`
	if err != nil {
		return nil, err
	}
	if n == 0 {
		return nil, nil
	}
	log.Infof("ok")
	return fileBlockInfList, err
}
func GetBlockList1(block models.Block) ([]models.Block, error) {
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
	sql := `select * from node where ip not in (SELECT node_ip from block WHERE block_num=? and file_id= ?)
`
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
	log.Infof("s")
	sql := `insert into block (file_id, block_num, block_size, node_ip, health) values (? ,? ,? ,?,"1") `
	p, err := o.Raw(sql).Prepare()
	if err != nil {
		log.Infof("1")
		return err
	}
	defer p.Close()

	r, err := p.Exec(fileId, blockNum, blockSize, nodeIp)
	log.Infof("11")
	if err != nil {
		log.Infof("2")
		return err
	}
	log.Infof("22")
	if n, err := r.RowsAffected(); n == 0 {
		log.Infof("3")
		return err
	}
	log.Infof("33")
	return nil
}
