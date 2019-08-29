package macro60

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
		htmlName = "macro60_by_hand.html"
	case "precreate":
		htmlName = "macro60_precreate.html"
	case "ctSequence":
		htmlName = "macro60_ct_sequence.html"
	case "info":
		htmlName = "macro60_info.html"
	case "list":
		htmlName = "macro60_list.html"
	case "share":
		htmlName = "macro60_share.html"
	case "verify":
		code, _ := c.Cookie("token")
		if code == "" || code != global.Config.VerifyCode {
			c.HTML(http.StatusUnauthorized, "404.html", nil)
			return
		}
		htmlName = "macro60_verify.html"
	case "home":
		htmlName = "macro60_index.html"
	default:
		name = "home"
		htmlName = "macro60_index.html"
	}
	modules.CreateLoginLog(c, htmlName)

	c.HTML(http.StatusOK, htmlName, gin.H{"apiPage": fmt.Sprintf("title-%s", name),})
}
