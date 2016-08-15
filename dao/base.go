package dao

import (
	"net"
	//"os"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" //register mysql driver

	"distribute_file_system/log"
	"sync"
)

var addr string
var port string
var pwd string
var user string
var db string

func init() {
	addr = "127.0.0.1"
	port = "3306"
	pwd = "abc123"
	user = "root"
	db = "dfs"
}

// InitDB initializes the database
func InitDB() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	log.Debugf("init db")
	dbConnectionStr := user + ":" + pwd + "@tcp(" + addr + ":" + port + ")/dfs"

	ch := make(chan int, 1)
	go func() {
		var err error
		var c net.Conn
		for {
			c, err = net.DialTimeout("tcp", addr+":"+port, 20*time.Second)
			if err == nil {
				c.Close()
				ch <- 1
			} else {
				log.Errorf("failed to connect to db, retry after 2 seconds :%v", err)
				time.Sleep(2 * time.Second)
			}
		}
	}()
	select {
	case <-ch:
	case <-time.After(60 * time.Second):
		panic("Failed to connect to DB after 60 seconds")
	}
	err := orm.RegisterDataBase("default", "mysql", dbConnectionStr)
	if err != nil {
		panic(err)
	}
}

var globalOrm orm.Ormer
var once sync.Once

// GetOrmer return global unique ormer, using singleton
func GetOrmer() orm.Ormer {
	once.Do(func() {
		globalOrm = orm.NewOrm()
	})
	return globalOrm
}
