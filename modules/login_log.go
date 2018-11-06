package modules

import (
	"github.com/illidan33/sql-builder"
	"time"
)

type LoginLog struct {
	// 访问IP
	IP string `json:"ip" db:"ip"`
	// 访问页面
	Method string `json:"method" db:"method"`
	// 访问时间
	CreateTime string `json:"createTime" db:"createtime"`
}

// 创建登录日志
func CreateLog(ip string, method string) bool {
	log := LoginLog{
		IP:         ip,
		Method:     method,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	builder := sql_builder.Insert("login_log")
	builder.InsertByStruct(log)

	conn := GetDbConn()
	res, err := conn.Exec(builder.String(), builder.Args()...)
	rowNum, err := res.RowsAffected()
	if err != nil {
		CheckErr(builder.String(), err)
		return false
	}

	if rowNum > 0 {
		return true
	}

	return false
}
