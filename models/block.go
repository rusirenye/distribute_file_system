package models

type Block struct {
	Id       int    `orm:"column(id)" json:"id"`
	FileId   string `orm:"column(file_id)" json:"file_id"`
	BlockNum int64  `orm:"column(block_num)" json:"block_num"`
	Size     int64  `orm:"column(block_size)" json:"block_size"`
	NodeIp   string `orm:"column(node_ip)" json:"node_ip"`
	Health   string `orm:"column(health)" json:"health"`
}

type FileBlockInfo struct {
	Id         int      `json:"id"`
	FileId     string   `json:"file_id"`
	BlockNum   int64    `json:"block_num"`
	Size       int64    `json:"block_size"`
	NodeIpList []string `json:"node_ip_list"`
	Health     string   `json:"health"`
}
