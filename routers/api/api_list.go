package api

import (
	"github.com/gin-gonic/gin"
	"github.com/illidan33/wow_api/global"
	"github.com/illidan33/wow_api/modules"
	"net/http"
)

func ApiList(c *gin.Context) {
	c.Request.ParseForm()
	pid := c.DefaultQuery("pid", "0")
	tableType := c.DefaultQuery("type", "title-api")

	global.Config.Log.Debugf("ApiList, pid: %s, type: %s.", pid, tableType)

	wowApis, err := modules.GetApiByParentID(tableType, pid)
	if err != nil {
		modules.Return(c, 500, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"list": wowApis,
	})
}
