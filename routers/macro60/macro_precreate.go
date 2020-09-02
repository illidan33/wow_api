package macro60

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/illidan33/wow_tools/database"
	"github.com/illidan33/wow_tools/global"
	"github.com/illidan33/wow_tools/modules"
)

func PreCreate(c *gin.Context) {
	modules.CreateLoginLog(c, "macro60_precreate", 2)

	name := c.Query("macro")
	if name == "" {
		modules.Return(c, 500, errors.New("params is empty"))
		return
	}
	global.Config.Log.Debugf("MacroPreCreate macro: %s", name)

	macros := make([]database.MacrosOld60, 0)
	err := modules.DbConn.Where("macro like ? and is_verify = 1", fmt.Sprintf("%%%s%%", name)).Find(&macros).Error
	if err != nil {
		modules.Return(c, 500, err)
	} else {
		modules.Return(c, 0, macros)
	}
}
