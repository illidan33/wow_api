package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/illidan33/wow_api/modules"
	"net/http"
)

func ApiIndex(c *gin.Context) {
	name := c.Param("name")

	var htmlName string
	switch name {
	case "api":
		htmlName = "api_api.html"
	case "event":
		htmlName = "api_event.html"
	case "macro":
		htmlName = "api_macro.html"
	case "widget":
		htmlName = "api_widget.html"
	case "widgetHandler":
		htmlName = "api_widget_handler.html"
	case "home":
		htmlName = "api_index.html"
	default:
		name = "home"
		htmlName = "api_index.html"
	}
	modules.CreateLoginLog(c, htmlName)

	c.HTML(http.StatusOK, htmlName, gin.H{
		"apiPage": fmt.Sprintf("title-%s", name),
	})
}

func DetailIndex(c *gin.Context) {
	modules.CreateLoginLog(c, "api_detail.html")

	apiType := c.Query("type")
	id := c.Param("id")
	api, err := modules.GetApiByID(apiType, id)
	if err != nil {
		c.HTML(http.StatusNotFound, "404.html", gin.H{})
		return
	}

	c.HTML(http.StatusOK, "api_detail.html", gin.H{
		"api":  api,
		"type": apiType,
	})
}
