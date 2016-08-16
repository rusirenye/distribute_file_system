package dao

import(
	"database/sql"
)

func GetNodeList(host string) {
	o := dao.GetOrmer()
	sql := `select * from node where host = ` + host
}
