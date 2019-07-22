package api

import (
	"github.com/gin-gonic/gin"
	"github.com/illidan33/wow_api/modules"
	"net/http"
	"strconv"
)

func ApiList(c *gin.Context) {
	c.Request.ParseForm()
	pid := c.DefaultQuery("pid", "0")
	pidd, _ := strconv.Atoi(pid)
	tableType := c.DefaultQuery("type", "title-wow-api")

	var table string
	switch tableType {
	case "title-wow-api":
		table = "api_wow"
		break
	case "title-wow-macro":
		table = "api_macro"
		break
	case "title-wow-event":
		table = "api_event"
		break
	case "title-wow-widget":
		table = "api_widget"
		break
	case "title-wow-widget-handler":
		table = "api_widget_handler"
		break
	default:
		table = "api_wow"
		break
	}

	// 记录日志
	if pidd != 0 || table == "api_event" {
		ip := c.ClientIP()
		go modules.UpdateOrCreateLog(ip, table)
	}

	wowApis := modules.GetApiByParentID(table, pidd)
	c.JSON(http.StatusOK, gin.H{
		"list": wowApis,
	})
}
