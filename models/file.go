package models

type File struct {
	Id         int    `orm:"column(id)"`
	FileId     string `orm:"column(file_id)" json:"file_id"`
	Name       string `orm:"column(name)" json:"name"`
	Size       int64  `orm:"column(size)" json:"size"`
	CreateTime string `orm:"colunm(create_time)" json:"create_time"`
	UpdateTime string `orm:"column(update_time)" json:"update_time"`
	UploadTime string `orm:"column(upload_time)" json:"upload_time"`
	Health     string `orm:"column(health)" json:"health"`
}
