package api

import (
	"github.com/gin-gonic/gin"
	"github.com/illidan33/wow_api/modules"
	"net/http"
)

func HomeIndex(c *gin.Context) {
	modules.CreateLoginLog(c, "ApiHomeIndex")

	c.HTML(http.StatusOK, "api_index.html", gin.H{
		"apiPage": "home",
	})
}

func ApiIndex(c *gin.Context) {
	modules.CreateLoginLog(c, "ApiApiIndex")

	c.HTML(http.StatusOK, "api_api.html", gin.H{
		"apiPage": "title-wow-api",
	})
}

func DetailIndex(c *gin.Context) {
	modules.CreateLoginLog(c, "ApiDetailIndex")

	apiType := c.Query("type")
	id := c.Param("id")
	api, err := modules.GetApiByID(apiType, id)
	if err != nil {
		c.HTML(http.StatusNotFound, "404.html", gin.H{})
		return
	}

	c.HTML(http.StatusOK, "api_detail.html", gin.H{
		"api": api,
	})
}

func EventIndex(c *gin.Context) {
	modules.CreateLoginLog(c, "ApiEventIndex")

	c.HTML(http.StatusOK, "api_event.html", gin.H{
		"apiPage": "title-wow-event",
	})
}

func MacroIndex(c *gin.Context) {
	modules.CreateLoginLog(c, "ApiMacroIndex")

	c.HTML(http.StatusOK, "api_macro.html", gin.H{
		"apiPage": "title-wow-macro",
	})
}
