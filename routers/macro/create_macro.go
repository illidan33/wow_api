package macro

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/illidan33/wow_tools/database"
	"github.com/illidan33/wow_tools/global"
	"github.com/illidan33/wow_tools/modules"
	"time"
)

func CreateMacro(c *gin.Context) {
	macro := database.Macros{}
	err := c.BindJSON(&macro)
	if err != nil {
		modules.Return(c, 500, errors.New("params is error"))
		return
	}
	global.Config.Log.Debugf("CreateSequence req: %+v", macro)

	macro.UpdateTime = time.Now()
	macro.IsVerify = 2
	err = modules.DbConn.Create(&macro).Error
	if err != nil {
		modules.Return(c, 500, "params is error")
		return
	}

	modules.CreateLoginLog(c, "macro_share", 2)
	modules.Return(c, 0, "ok")
}
