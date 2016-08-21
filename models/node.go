package models

// import (
// //"github.com/astaxie/beego/orm"
// )

type Node struct {
	Id             int     `orm:"column(id)" json:"id"`
	Ip             string  `orm:"column(ip)" json:"ip"`
	Brandwidth     float32 `orm:"column(brandwidth)" json:"brandwidth"`
	BrandwidthUsed float32 `orm:"column(brandwidth_used)" json:"brandwidth_used"`
	Disk           float32 `orm:"column(disk)" json:"disk"`
	DiskUsed       float32 `orm:"column(disk_used)" json:"disk_used"`
	Cpu            float32 `orm:"column(cpu)" json:"cpu"`
	CpuUsed        float32 `orm:"column(cpu_used)" json:"cpu_used"`
	Memory         float32 `orm:"column(memory)" json:"memory"`
	MemoryUsed     float32 `orm:"column(memory_used)" json:"memory_used"`
	HostIp         string  `orm:"column(host_ip)" json:"host_ip"`
	Health         string  `orm:column(health)" json:"health"`
}
