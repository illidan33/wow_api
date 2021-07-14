package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/illidan33/wow_api/routers/api"
	"github.com/illidan33/wow_api/routers/index"
	"github.com/illidan33/wow_api/routers/macro"
	"github.com/illidan33/wow_api/routers/macro60"
)

var Chart *gin.RouterGroup
var Auth *gin.RouterGroup
var Api *gin.RouterGroup
var Macro *gin.RouterGroup
var MacroOld60 *gin.RouterGroup

func New() {
	// Chart
	Chart.GET("/", index.ChartIndex)
	Chart.GET("/data", index.GetChartData)

	// Auth
	Auth.PUT("/macro", macro.UpdateMacro)
	Auth.PUT("/macro60", macro60.UpdateMacro)

	// Api
	Api.GET("/", api.ApiIndex)
	Api.GET("/view/:name", api.ApiIndex)
	Api.GET("/view", api.ApiIndex)

	Api.GET("/detail/:id", api.DetailIndex)
	Api.GET("/forgnDetail/:id", api.ForeignDetailIndex)
	Api.GET("/list", api.ApiList)
	Api.GET("/search", api.ApiSearch)
	Api.POST("/apiUnverify", api.SaveUnverifyApi)

	// Macro
	Macro.GET("/", macro.ViewIndex)
	Macro.GET("/view/:name", macro.ViewIndex)
	Macro.GET("/view", macro.ViewIndex)

	Macro.GET("/preCreate", macro.PreCreate)
	Macro.GET("/macroList", macro.MacroList)
	Macro.GET("/professionList", macro.ProfessionList)
	Macro.POST("/createSequence", macro.CreateSequence)
	Macro.POST("/combineSkills", macro.CombineSkills)
	Macro.POST("/", macro.CreateMacro)

	// MacroOld60
	MacroOld60.GET("/", macro60.ViewIndex)
	MacroOld60.GET("/view/:name", macro60.ViewIndex)
	MacroOld60.GET("/view", macro60.ViewIndex)

	MacroOld60.GET("/preCreate", macro60.PreCreate)
	MacroOld60.GET("/macroList", macro60.MacroList)
	MacroOld60.POST("/createSequence", macro60.CreateSequence)
	MacroOld60.POST("/", macro60.CreateMacro)

}
