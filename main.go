package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/illidan33/wow_tools/global"
	"github.com/illidan33/wow_tools/modules"
	"github.com/illidan33/wow_tools/public"
	"github.com/illidan33/wow_tools/routers"
	"github.com/illidan33/wow_tools/routers/index"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
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
	routers.Chart = WowApi.Group("/chart")
	routers.Chart.Use(index.AuthMiddleware)
	routers.Auth = WowApi.Group("/auth")
	routers.Auth.Use(index.AuthMiddleware)

	routers.Api = WowApi.Group("/api")
	routers.Macro = WowApi.Group("/macro")
	routers.MacroOld60 = WowApi.Group("/macro60")
	routers.New()
	public.New(WowApi)

	WowApi.GET("/", index.Index)
	WowApi.NoRoute(index.Index)
	WowApi.Run(fmt.Sprintf("%s:%d", global.Config.ListenHost, global.Config.ListenPort))
}
