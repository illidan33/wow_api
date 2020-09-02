package macro

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/illidan33/wow_tools/global"
	"github.com/illidan33/wow_tools/modules"
	"net/http"
)

func ViewIndex(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		name = "home"
	}
	if name == "verify" {
		code, _ := c.Cookie("token")
		if code == "" || code != global.Config.VerifyCode {
			c.HTML(http.StatusUnauthorized, "404.html", nil)
			return
		}
	}

	htmlName := fmt.Sprintf("macro_%s.html", name)
	modules.CreateLoginLog(c, htmlName, 1)

	c.HTML(http.StatusOK, htmlName, gin.H{"apiPage": fmt.Sprintf("%s", name),})
}
