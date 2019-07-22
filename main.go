package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/illidan33/wow_api/global"
	"github.com/illidan33/wow_api/public"
	"github.com/illidan33/wow_api/routers"
	"github.com/illidan33/wow_api/routers/api"
	"github.com/jinzhu/gorm"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"net/http"
)

var (
	WowApi *gin.Engine
	DB     *gorm.DB
)

func NewLogger() *logrus.Logger {
	if global.Config.Log != nil {
		return global.Config.Log
	}

	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  global.Config.LogPath,
		logrus.ErrorLevel: global.Config.LogPath,
		logrus.WarnLevel:  global.Config.LogPath,
	}
	global.Config.Log = logrus.New()
	global.Config.Log.Hooks.Add(lfshook.NewHook(
		pathMap,
		&logrus.JSONFormatter{},
	))
	return global.Config.Log
}

func init() {
	global.Config.Log = NewLogger()
}

func main() {
	WowApi = gin.New()
	routers.Api = WowApi.Group("/api")
	routers.New()

	routers.Macro = WowApi.Group("/macro")
	public.New(WowApi)

	WowApi.GET("/", api.Index)
	WowApi.Run(fmt.Sprintf("127.0.0.1:%d", global.Config.ListenPort))
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"apiPage": "home",
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
