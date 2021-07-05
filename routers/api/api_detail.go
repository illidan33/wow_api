package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/illidan33/wow_api/database"
	"github.com/illidan33/wow_api/global"
	"github.com/illidan33/wow_api/modules"
)

func ApiDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		modules.Return(c, 500, errors.New("id is empty"))
		return
	}
	global.Config.Log.Debugf("ApiDetailHandle id: %s", id)

	event := database.ApiItem{}
	err := modules.DbConn.Where("id = ?", id).First(&event).Error
	if err != nil {
		modules.Return(c, 500, err)
		return
	}

	global.Config.Log.Debugf("ApiDetailHandle resp: %+v", event)
	modules.Return(c, 0, event)
}
