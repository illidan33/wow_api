package macro

import (
	"github.com/gin-gonic/gin"
	"github.com/illidan33/wow_api/global"
	"github.com/illidan33/wow_api/modules"
	"net/http"
)

func Index(c *gin.Context) {
	modules.CreateLoginLog(c, "MacroIndex")

	c.HTML(http.StatusOK, "macro_index.html", gin.H{
		"apiPage": "home",
	})
}

func ByHandIndex(c *gin.Context) {
	modules.CreateLoginLog(c, "MacroByHandIndex")

	c.HTML(http.StatusOK, "macro_by_hand.html", gin.H{
		"apiPage": "title-macro-byhand",
	})
}

func CreateIndex(c *gin.Context) {
	modules.CreateLoginLog(c, "MacroCreateIndex")

	c.HTML(http.StatusOK, "macro_create.html", gin.H{
		"apiPage": "title-macro-create",
	})
}

func CtSequenceIndex(c *gin.Context) {
	modules.CreateLoginLog(c, "MacroCtSequenceIndex")

	c.HTML(http.StatusOK, "macro_ct_sequence.html", gin.H{
		"apiPage": "title-macro-ctsequence",
	})
}

func InfoIndex(c *gin.Context) {
	modules.CreateLoginLog(c, "MacroInfoIndex")

	c.HTML(http.StatusOK, "macro_info.html", gin.H{
		"apiPage": "title-macro-info",
	})
}

func ListIndex(c *gin.Context) {
	modules.CreateLoginLog(c, "MacroListIndex")

	c.HTML(http.StatusOK, "macro_list.html", gin.H{
		"apiPage": "title-macro-list",
	})
}

func ShareIndex(c *gin.Context) {
	modules.CreateLoginLog(c, "MacroShareIndex")

	c.HTML(http.StatusOK, "macro_share.html", gin.H{
		"apiPage": "title-macro-share",
	})
}

func VerifyIndex(c *gin.Context) {
	modules.CreateLoginLog(c, "MacroVerifyIndex")

	code, err := c.Cookie("code")
	if err != nil {
		c.HTML(http.StatusNotFound, "404.html", gin.H{})
		return
	}
	if code == "" || code != global.Config.VerifyCode {
		c.HTML(http.StatusNotFound, "404.html", gin.H{})
		return
	}

	c.HTML(http.StatusOK, "macro_verify.html", gin.H{
		"apiPage": "title-macro-verify",
	})
}
