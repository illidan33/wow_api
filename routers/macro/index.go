package macro

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/illidan33/wow_api/global"
	"github.com/illidan33/wow_api/modules"
	"net/http"
)

func ViewIndex(c *gin.Context) {
	name := c.Param("name")

	var htmlName string
	switch name {
	case "byHand":
		htmlName = "macro_by_hand.html"
	case "precreate":
		htmlName = "macro_precreate.html"
	case "ctSequence":
		htmlName = "macro_ct_sequence.html"
	case "info":
		htmlName = "macro_info.html"
	case "list":
		htmlName = "macro_list.html"
	case "share":
		htmlName = "macro_share.html"
	case "verify":
		code, _ := c.Cookie("token")
		if code == "" || code != global.Config.VerifyCode {
			c.HTML(http.StatusUnauthorized, "404.html", nil)
			return
		}
		htmlName = "macro_verify.html"
	case "home":
		htmlName = "macro_index.html"
	default:
		name = "home"
		htmlName = "macro_index.html"
	}
	modules.CreateLoginLog(c, htmlName)

	c.HTML(http.StatusOK, htmlName, gin.H{"apiPage": fmt.Sprintf("title-%s", name),})
}
