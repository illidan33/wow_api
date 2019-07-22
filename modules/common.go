package modules

import (
	"github.com/gin-gonic/gin"
	"github.com/illidan33/wow_api/global"
	"net/http"
)

func Return(c *gin.Context, code int32, resp interface{}) {
	if e, ok := resp.(error); ok {
		global.Config.Log.Error(e)
		if e.Error() == "record not found" {
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  "Interner Error",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  e.Error(),
			})
		}

	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"data": resp,
		})
	}
}

func IsNotFound(err error) bool {
	if err.Error() == "record not found" {
		return true
	}
	return false
}
