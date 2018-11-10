package main

import (
	"flag"
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/illidan33/wow_api/modules"
	"strconv"
	"time"
)

var (
	port int
)

func main() {
	rootPath := "/data/golang/go/src/github.com/illidan33/wow_api/"

	flag.IntVar(&port, "port", 8001, "listen port")
	flag.Parse()

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.LoadHTMLGlob(rootPath + "html/*")

	// 设置静态资源
	router.Static("/js", rootPath+"js")
	router.Static("/css", rootPath+"css")
	router.Static("/img", rootPath+"img")
	router.StaticFile("/favicon.ico", rootPath+"favicon.ico")
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "404, page not exists!",
		})
	})

	// 模块
	router.GET("/", Index)
	router.GET("/Api", ApiIndex)
	router.GET("/Event", EventIndex)
	router.GET("/Macro", MacroIndex)
	router.GET("/Widget", WidgetIndex)
	router.GET("/WidgetHandler", WidgetHandlerIndex)

	// data
	router.POST("/wow", GetApi)

	router.Run(fmt.Sprintf(":%d", port))
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
func WidgetHandlerIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "wow_widget_handler.html", gin.H{})
}

func GetApi(c *gin.Context) {
	c.Request.ParseForm()
	pid, _ := c.GetPostForm("pid")
	pidd, _ := strconv.Atoi(pid)
	tableType, _ := c.GetPostForm("type")

	// 记录日志
	if pidd != 0 || tableType == "Event" {
		ip := c.ClientIP()
		go CheckLoginLog(ip, tableType)
	}

	var table string
	switch tableType {
	case "Api":
		table = "api_wow"
		break
	case "Macro":
		table = "api_macro"
		break
	case "Event":
		table = "api_event"
		break
	case "Widget":
		table = "api_widget"
		break
	case "WidgetHandler":
		table = "api_widget_handler"
		break
	default:
		table = "api_wow"
		break
	}
	wowApis := modules.GetApiByParentID(table, pidd)
	c.JSON(http.StatusOK, gin.H{
		"list": wowApis,
	})
}

func CheckLoginLog(ip string, method string) {
	date := time.Now().Format("2006-01-02")

	log, err := modules.GetLog(ip, method, date)
	if err != nil || log.ID == 0 {
		err = modules.CreateLog(ip, method)
		if err != nil {
			modules.CheckErr("CreateLog", err)
		}
	} else {
		err = modules.UpdateLog(log.ID, log.Count+1)
		if err != nil {
			modules.CheckErr("UpdateLog", err)
		}
	}
}
