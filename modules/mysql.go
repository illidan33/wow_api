package modules

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/illidan33/wow_api/global"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var DbConn *gorm.DB

func init() {
	var err error
	DbConn, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8&parseTime=True&loc=Local", global.Config.DbUser, global.Config.DbPwd, global.Config.DbName))
	if err != nil {
		panic(err)
	}
	DbConn.SingularTable(true)
	DbConn.SetLogger(global.Log)
	if global.Config.LogLevel == logrus.DebugLevel {
		DbConn.LogMode(true)
	}
}
