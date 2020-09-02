package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/illidan33/wow_tools/modules"
	"net/http"
)

func ApiIndex(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		name = "home"
	}

	htmlName := fmt.Sprintf("api_%s.html", name)
	modules.CreateLoginLog(c, htmlName, 1)

	c.HTML(http.StatusOK, htmlName, gin.H{
		"apiPage": fmt.Sprintf("%s", name),
	})
}

func DetailIndex(c *gin.Context) {
	html := "api_detail.html"
	modules.CreateLoginLog(c, html, 1)

	apiType := c.Query("type")
	id := c.Param("id")
	api, err := modules.GetApiByID(id)
	if err != nil {
		c.HTML(http.StatusNotFound, "404.html", gin.H{})
		return
	}

	c.HTML(http.StatusOK, html, gin.H{
		"api":  api,
		"type": apiType,
	})
}

func ForeignDetailIndex(c *gin.Context) {
	modules.CreateLoginLog(c, "api_foreign.html", 1)

	apiType := c.Query("type")
	id := c.Param("id")
	api, err := modules.GetApiByID(id)
	if err != nil {
		c.HTML(http.StatusNotFound, "404.html", gin.H{})
		return
	}
	url := modules.GetApiDetailUrlByID(apiType, api.Name)

	c.Redirect(http.StatusFound, url)
}
