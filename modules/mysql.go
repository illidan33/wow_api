package modules

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"strings"
)

var DbConn *sqlx.DB

func init() {
	DbConnetc()
}

func GetDbConn() *sqlx.DB {
	if DbConn == nil {
		DbConnetc()
	}
	return DbConn
}

func DbConnetc() {
	var err error
	DbConn, err = sqlx.Open("mysql", "test:test@tcp(127.0.0.1:3306)/wow_hong?charset=utf8")
	if err != nil {
		CheckErr("Connect Database", err)
	}

	DbConn.SetMaxOpenConns(200)
	DbConn.SetMaxIdleConns(100)
	err = DbConn.Ping()
	if err != nil {
		CheckErr("Ping Database", err)
	}
}

func CheckErr(msg string, err error) {
	fmt.Fprintf(gin.DefaultWriter, "%s : %s\n", msg, err.Error())
}

func Debug(msg interface{}) {
	if msg == nil {
		fmt.Println("nil")
		return
	}
	tp := reflect.TypeOf(msg)
	if tp.Name() == "string" {
		fmt.Fprintf(gin.DefaultWriter, "%s\n", msg)
	} else if strings.Contains(tp.Name(), "int") {
		fmt.Fprintf(gin.DefaultWriter, "%d\n", msg)
	} else if strings.Contains(tp.Name(), "float") {
		fmt.Fprintf(gin.DefaultWriter, "%f\n", msg)
	} else if strings.Contains(tp.Name(), "map") {
		fmt.Fprintf(gin.DefaultWriter, "%+v\n", msg)
	} else {
		fmt.Fprintf(gin.DefaultWriter, "%s\n", msg)
	}
}
