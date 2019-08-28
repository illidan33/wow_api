package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/illidan33/wow_api/routers/api"
	"github.com/illidan33/wow_api/routers/macro"
)

var Api *gin.RouterGroup
var Macro *gin.RouterGroup

func New() {
	// Api
	Api.GET("/", api.HomeIndex)
	Api.GET("/view/:name", api.ApiIndex)
	//Api.GET("/event", api.EventIndex)
	//Api.GET("/macro", api.MacroIndex)
	Api.GET("/detail/:id", api.DetailIndex)

	Api.GET("/list", api.ApiList)
	Api.POST("/apiUnverify", api.SaveUnverifyApi)

	// Macro
	Macro.GET("/", macro.Index)
	Macro.GET("/byHandIndex", macro.ByHandIndex)
	Macro.GET("/precreateIndex", macro.PreCreateIndex)
	Macro.GET("/ctSequenceIndex", macro.CtSequenceIndex)
	Macro.GET("/infoIndex", macro.InfoIndex)
	Macro.GET("/listIndex", macro.ListIndex)
	Macro.GET("/shareIndex", macro.ShareIndex)
	Macro.GET("/verifyIndex", macro.VerifyIndex)

	Macro.GET("/preCreate", macro.PreCreate)
	Macro.POST("/createSequence", macro.CreateSequence)
	Macro.GET("/macroList", macro.MacroList)
	Macro.POST("/", macro.CreateMacro)

}
