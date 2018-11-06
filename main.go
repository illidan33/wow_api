package main

import (
	"flag"
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"gotest/wow_api_list/modules"
	"strconv"
)

var (
	port int
)

func main() {
	flag.IntVar(&port, "port", 8000, "listen port")
	flag.Parse()

	//gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.LoadHTMLGlob("html/*")

	// 设置静态资源
	router.Static("/js", "js")
	router.Static("/css", "css")
	//router.Static("/html", "html")
	router.Static("/img", "img")
	router.StaticFile("/favicon.ico", "./favicon.ico")
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "404, page not exists!",
		})
	})

	router.POST("/log/:method", CreateLoginLog)

	router.GET("/", Index)
	router.GET("/api", ApiIndex)
	router.GET("/event", EventIndex)
	router.GET("/macro", MacroIndex)
	router.GET("/widget", WidgetIndex)

	router.GET("/wowApi", GetWowApi)
	router.GET("/wowEvent", GetWowEvent)
	router.GET("/wowMacro", GetWowMacro)
	router.GET("/wowWidget", GetWowWidget)

	router.Run(fmt.Sprintf(":%d", port))
}

func CreateLoginLog(c *gin.Context) {
	method := c.Param("method")
	// 记录日志
	ip := c.ClientIP()
	go modules.CreateLog(ip, method)

	c.JSON(http.StatusOK, gin.H{})
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
func ApiIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "wow_api.html", gin.H{})
}
func EventIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "wow_event.html", gin.H{})
}
func MacroIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "wow_macro.html", gin.H{})
}
func WidgetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "wow_widget.html", gin.H{})
}

func GetWowApi(c *gin.Context) {
	// 记录日志
	ip := c.ClientIP()
	go modules.CreateLog(ip, "wow_api")

	pid, _ := strconv.Atoi(c.Query("pid"))
	wowApis := modules.GetWowApiByParentID(pid)
	c.JSON(http.StatusOK, gin.H{
		"list": wowApis,
	})
}

func GetWowEvent(c *gin.Context) {
	// 记录日志
	ip := c.ClientIP()
	go modules.CreateLog(ip, "wow_event")

	//pid, _ := strconv.Atoi(c.Query("pid"))
	wowApis := modules.GetWowEventByParentID(0)
	c.JSON(http.StatusOK, gin.H{
		"list": wowApis,
	})
}

func GetWowMacro(c *gin.Context) {
	// 记录日志
	ip := c.ClientIP()
	go modules.CreateLog(ip, "wow_macro")

	pid, _ := strconv.Atoi(c.Query("pid"))
	wowApis := modules.GetWowMacroByParentID(pid)
	c.JSON(http.StatusOK, gin.H{
		"list": wowApis,
	})
}

func GetWowWidget(c *gin.Context) {
	// 记录日志
	ip := c.ClientIP()
	go modules.CreateLog(ip, "wow_widget")

	pid, _ := strconv.Atoi(c.Query("pid"))
	wowApis := modules.GetWowWidgetByParentID(pid)
	c.JSON(http.StatusOK, gin.H{
		"list": wowApis,
	})
}
