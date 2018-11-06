package modules

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)

var DbConn *sqlx.DB

func init() {
	DbConnetc()
}

func GetDbConn() *sqlx.DB {
	var err error
	err = DbConn.Ping()
	if err != nil {
		DbConnetc()
	}
	return DbConn
}

func DbConnetc() {
	var err error
	DbConn, err = sqlx.Open("mysql", "test:test@tcp(127.0.0.1:3306)/wow_hong?charset=utf8")
	CheckErr("Connect Database", err)

	DbConn.SetMaxOpenConns(1000)
	DbConn.SetMaxIdleConns(2000)
	err = DbConn.Ping()
	CheckErr("Ping Database", err)
}

func CheckErr(msg string, err error) {
	if err != nil {
		fmt.Fprintf(gin.DefaultWriter, "%s : %+v\n", msg, err)
	}
}

func Debug(msg interface{}) {
	fmt.Fprintf(gin.DefaultWriter, "%s\n", msg)
}