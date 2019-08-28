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
	Api.GET("/detail/:id", api.DetailIndex)

	Api.GET("/list", api.ApiList)
	Api.POST("/apiUnverify", api.SaveUnverifyApi)

	// Macro
	Macro.GET("/", macro.Index)
	Macro.GET("/view/:name", macro.ViewIndex)
	Macro.GET("/preCreate", macro.PreCreate)
	Macro.GET("/macroList", macro.MacroList)

	Macro.POST("/createSequence", macro.CreateSequence)
	Macro.PUT("/updateMacro", macro.UpdateMacro)
	Macro.POST("/", macro.CreateMacro)

}
