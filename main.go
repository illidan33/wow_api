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
	"os"
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
	if os.Getenv("GOPATH") == "" {
		global.Config.ApiRootPath = "/data/go/src/github.com/illidan33/wow_api"
	}

	WowApi = gin.New()
	routers.Api = WowApi.Group("/api")
	routers.Macro = WowApi.Group("/macro")
	routers.New()
	public.New(WowApi)

	WowApi.GET("/", Index)
	WowApi.Run(fmt.Sprintf("%s:%d", global.Config.ListenHost, global.Config.ListenPort))
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"apiPage": "home",
	})
}
