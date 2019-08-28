package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/illidan33/wow_api/global"
	"github.com/illidan33/wow_api/modules"
	"github.com/illidan33/wow_api/public"
	"github.com/illidan33/wow_api/routers"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"net/http"
)

var (
	WowApi *gin.Engine
	DB     *gorm.DB
)

func main() {
	defer func() {
		modules.DbConn.Close()
	}()
	if global.Config.LogLevel != logrus.DebugLevel {
		gin.SetMode(gin.ReleaseMode)
	}

	WowApi = gin.New()
	routers.Api = WowApi.Group("/api")
	routers.Macro = WowApi.Group("/macro")
	routers.MacroOld60 = WowApi.Group("/macro60")
	routers.New()
	public.New(WowApi)

	WowApi.GET("/", Index)
	WowApi.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, nil)
	})
	WowApi.Run(fmt.Sprintf("%s:%d", global.Config.ListenHost, global.Config.ListenPort))
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"apiPage": "home",
	})
}
