package macro

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/illidan33/wow_api/modules"
	"net/http"
)

func Index(c *gin.Context) {
	modules.CreateLoginLog(c, "MacroIndex")

	c.HTML(http.StatusOK, "macro_index.html", gin.H{
		"apiPage": "home",
	})
}

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
		htmlName = "macro_verify.html"
	default:
		htmlName = ""
	}
	if htmlName == "" {
		c.HTML(http.StatusNotFound, "404.html", nil)
		return
	}
	modules.CreateLoginLog(c, name)

	c.HTML(http.StatusOK, htmlName, gin.H{"apiPage": fmt.Sprintf("title-%s", name),})
}
