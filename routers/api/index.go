package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"apiPage": "title-wow-index",
	})
}

func ApiIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "wow_api.html", gin.H{
		"apiPage": "title-wow-api",
	})
}

func EventIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "wow_event.html", gin.H{
		"apiPage": "title-wow-event",
	})
}

func MacroIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "wow_macro.html", gin.H{
		"apiPage": "title-wow-macro",
	})
}
