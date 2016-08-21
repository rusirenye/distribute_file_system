package models

type Block struct {
	Id       int    `orm:"column(id)" json:"id"`
	FileId   string `orm:"column(file_id)" json:"file_id"`
	BlockNum int64  `orm:"column(block_num)" json:"block_num"`
	Size     int64  `orm:"column(block_size)" json:"block_size"`
	NodeIp   string `orm:"column(node_ip)" json:"node_ip"`
	Health   string `orm:"column(health)" json:"health"`
}
