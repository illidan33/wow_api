package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/illidan33/wow_api/routers/api"
	"github.com/illidan33/wow_api/routers/macro"
	"github.com/illidan33/wow_api/routers/macro60"
)

var Api *gin.RouterGroup
var Macro *gin.RouterGroup
var MacroOld60 *gin.RouterGroup

func New() {
	// Api
	Api.GET("/", api.ApiIndex)
	Api.GET("/view/:name", api.ApiIndex)
	Api.GET("/view", api.ApiIndex)
	Api.GET("/detail/:id", api.DetailIndex)

	Api.GET("/list", api.ApiList)
	Api.POST("/apiUnverify", api.SaveUnverifyApi)

	// Macro
	Macro.GET("/", macro.ViewIndex)
	Macro.GET("/view/:name", macro.ViewIndex)
	Macro.GET("/view", macro.ViewIndex)
	Macro.GET("/preCreate", macro.PreCreate)
	Macro.GET("/macroList", macro.MacroList)

	Macro.POST("/createSequence", macro.CreateSequence)
	Macro.PUT("/updateMacro", macro.UpdateMacro)
	Macro.POST("/", macro.CreateMacro)

	// MacroOld60
	MacroOld60.GET("/", macro60.ViewIndex)
	MacroOld60.GET("/view/:name", macro60.ViewIndex)
	MacroOld60.GET("/view", macro60.ViewIndex)
	MacroOld60.GET("/preCreate", macro60.PreCreate)
	MacroOld60.GET("/macroList", macro60.MacroList)

	MacroOld60.POST("/createSequence", macro60.CreateSequence)
	MacroOld60.PUT("/updateMacro", macro60.UpdateMacro)
	MacroOld60.POST("/", macro60.CreateMacro)

}
