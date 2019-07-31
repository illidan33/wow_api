package public

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/illidan33/wow_api/global"
	"net/http"
)

func New(router *gin.Engine) {
	rootPath := fmt.Sprintf("%s/public/", global.Config.ApiRootPath)

	// 设置页面路径
	router.LoadHTMLGlob(rootPath + "html/*/*")

	// 设置静态资源
	router.Static("/js", rootPath+"js")
	router.Static("/css", rootPath+"css")
	//router.Static("/img", rootPath+"img")
	router.StaticFile("/favicon.ico", rootPath+"favicon.ico")
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "404 Not Found",
		})
	})
}
