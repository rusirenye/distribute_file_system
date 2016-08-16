package models

// import (
// //"github.com/astaxie/beego/orm"
// )

type Node struct {
	Id         int     `orm:"column(id) json:"id"`
	Ip         string  `orm:"column(ip) json:"ip"`
	Brandwidth float32 `orm:"column(brandwidth) json:"brandwidth"`
	Disk       float32 `orm:"column(disk) json:"disk"`
	Cpu        float32 `orm:"column(cpu) json:"cpu"`
	Memory     float32 `orm:"column(memory) json:"memory"`
	Health     string  `orm:column(health) json:"health"`
}
