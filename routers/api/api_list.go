package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/illidan33/wow_api/global"
	"github.com/illidan33/wow_api/modules"
	"net/http"
)

func ApiList(c *gin.Context) {
	c.Request.ParseForm()
	pid := c.DefaultQuery("pid", "0")
	tableType := c.DefaultQuery("type", "api")

	global.Log.Debugf("ApiList, pid: %s, type: %s.", pid, tableType)

	wowApis, err := modules.GetApiByParentID(tableType, pid)
	if err != nil {
		modules.Return(c, 500, err.Error())
		return
	}

	modules.CreateLoginLog(c, fmt.Sprintf("api_%s_list", tableType), 2)
	c.JSON(http.StatusOK, gin.H{
		"list": wowApis,
	})
}
