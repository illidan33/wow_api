package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/illidan33/wow_api/database"
	"github.com/illidan33/wow_api/modules"
	"strconv"
	"time"
)

func SaveUnverifyApi(c *gin.Context) {
	c.Request.ParseForm()

	id, _ := strconv.Atoi(c.PostForm("id"))

	if id == 0 {
		modules.Return(c, 500, errors.New("params error"))
		return
	}
	api := database.ApiUnverify{
		ApiID:      id,
		Type:       c.PostForm("type"),
		Name:       c.PostForm("name"),
		NameCn:     c.PostForm("nameCn"),
		Desc:       c.PostForm("desc"),
		InfoDesc:   c.PostForm("infoDesc"),
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
	}

	err := modules.DbConn.Create(&api).Error
	if err != nil {
		modules.Return(c, 500, errors.New("params error"))
	} else {
		modules.Return(c, 0, "ok")
	}
}
