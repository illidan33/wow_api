package macro60

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/illidan33/wow_tools/database"
	"github.com/illidan33/wow_tools/global"
	"github.com/illidan33/wow_tools/modules"
	"github.com/sirupsen/logrus"
	"net/http"
)

func UpdateMacro(c *gin.Context) {
	modules.CreateLoginLog(c, "macro60_update", 2)

	code, err := c.Cookie("token")
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusNotFound, nil)
		return
	}
	if code == "" || code != global.Config.VerifyCode {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}
	macro := database.MacrosOld60{}
	err = c.BindJSON(&macro)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	if macro.ID == 0 {
		logrus.Error(errors.New("id empty"))
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	err = modules.DbConn.Model(&macro).Updates(macro).Error
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
	})
}
