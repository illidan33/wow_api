package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/illidan33/wow_api/routers/api"
)

var Api *gin.RouterGroup
var Macro *gin.RouterGroup

func New() {
	Api.GET("/", api.ApiIndex)
	Api.GET("/event", api.EventIndex)
	Api.GET("/macro", api.MacroIndex)

	Api.GET("/detail/:id", api.ApiDetail)
	Api.GET("/list", api.ApiList)
	Api.POST("/apiUnverify", api.SaveUnverifyApi)

}
