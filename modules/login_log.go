package modules

import (
	"github.com/illidan33/sql-builder"
	"time"
)

// 创建登录日志
func CreateLog(ip string, method string) error {
	now := time.Now()
	t := now.Format("2006-01-02 15:04:05")

	log := LoginLog{
		ID:         0,
		IP:         ip,
		Method:     method,
		LoginDate:  now.Format("2006-01-02"),
		Count:      1,
		CreateTime: t,
		UpdateTime: t,
	}
	builder := sql_builder.Insert("api_login_log")
	builder.InsertByStruct(log)

	conn := GetDbConn()
	_, err := conn.Exec(builder.String(), builder.Args()...)
	if err != nil {
		return err
	}

	return nil
}

func UpdateLog(id int, count int) error {
	builder := sql_builder.Update("api_login_log")
	builder.WhereEq("id", id)
	builder.UpdateSet("count", count)

	conn := GetDbConn()
	_, err := conn.Exec(builder.String(), builder.Args()...)
	if err != nil {
		return err
	}

	return nil
}

func GetLog(ip string, method string, date string) (LoginForGet, error) {
	builder := sql_builder.Select("api_login_log")
	builder.WhereEq("ip", ip)
	builder.WhereEq("method", method)
	builder.WhereEq("login_date", date)
	builder.SetSearchFields([]string{"id", "ip", "method", "login_date", "count"})

	log := LoginForGet{}
	conn := GetDbConn()
	err := conn.Get(&log, builder.String(), builder.Args()...)
	if err != nil {
		return LoginForGet{}, err
	}
	return log, nil
}
