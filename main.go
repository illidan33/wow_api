package main

import (
	"flag"
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/illidan33/wow_api/modules"
	"strconv"
	"time"
	"os"
)

var (
	port int
)

func main() {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = "/test"
	}
	rootPath := fmt.Sprintf("%s/src/github.com/illidan33/wow_api/", gopath);
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
	router.GET("/ApiDetail/:type/:id", ApiDetailHandle)

	// data
	router.POST("/wow", GetApi)
	router.POST("/saveApi", SaveUnverifyApi)

	router.Run(fmt.Sprintf(":%d", port))
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"apiPage": "home",
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
func WidgetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "wow_widget.html", gin.H{
		"apiPage": "title-wow-widget",
	})
}
func WidgetHandlerIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "wow_widget_handler.html", gin.H{
		"apiPage": "title-wow-widget-handler",
	})
}

func ApiDetailHandle(c *gin.Context) {
	tableType := c.Param("type")
	id, _ := strconv.Atoi(c.Param("id"))

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
	wowApi := modules.GetApiByID(table, id)
	c.HTML(http.StatusOK, "edit_box.html", gin.H{
		"api":  wowApi,
		"type": tableType,
	})
}

func GetApi(c *gin.Context) {
	c.Request.ParseForm()
	pid, _ := c.GetPostForm("pid")
	pidd, _ := strconv.Atoi(pid)
	tableType, _ := c.GetPostForm("type")

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
		go CheckLoginLog(ip, table)
	}

	wowApis := modules.GetApiByParentID(table, pidd)
	c.JSON(http.StatusOK, gin.H{
		"list": wowApis,
	})
}

func SaveUnverifyApi(c *gin.Context) {
	c.Request.ParseForm()

	id, _ := strconv.Atoi(c.PostForm("id"))
	api := modules.UnVerifyApi{
		ApiID:      id,
		Type:       c.PostForm("type"),
		Name:       c.PostForm("name"),
		NameCn:     c.PostForm("nameCn"),
		Desc:       c.PostForm("desc"),
		InfoDesc:   c.PostForm("infoDesc"),
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
	}

	err := modules.SaveApiUnverify(api)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
		})
	} else {
		c.JSON(http.StatusOK, gin.H{})
	}
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
