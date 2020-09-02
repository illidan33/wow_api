package macro

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/illidan33/wow_tools/modules"
	"net/http"
)

type CombineSkillsReq struct {
	Mousehp string `json:"mousehp"`
	Mousehm string `json:"mousehm"`
	Focus   string `json:"focus"`
	Tarhp   string `json:"tarhp"`
	Tarhm   string `json:"tarhm"`
	Shift   string `json:"shift"`
	Alt     string `json:"alt"`
	Ctrl    string `json:"ctrl"`
	Player  string `json:"player"`
	Def     string `json:"default"`
}

func CombineSkills(c *gin.Context) {
	req := CombineSkillsReq{}
	var buf bytes.Buffer
	buf.WriteString("#showtooltip\r\n")

	err := c.Bind(&req)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
		})
		return
	}
	if req.Mousehp != "" {
		buf.WriteString("/cast [@mouseover,help]")
		buf.WriteString(req.Mousehp)
		buf.WriteString("\r\n")
	}
	if req.Mousehm != "" {
		buf.WriteString("/cast [@mouseover,harm]")
		buf.WriteString(req.Mousehm)
		buf.WriteString("\r\n")
	}
	if req.Focus != "" {
		buf.WriteString("/cast [@focus,nodead]")
		buf.WriteString(req.Focus)
		buf.WriteString("\r\n")
	}
	if req.Tarhp != "" {
		buf.WriteString("/cast [@target,help]")
		buf.WriteString(req.Tarhp)
		buf.WriteString("\r\n")
	}
	if req.Tarhm != "" {
		buf.WriteString("/cast [@target,harm]")
		buf.WriteString(req.Tarhm)
		buf.WriteString("\r\n")
	}
	if req.Shift != "" {
		buf.WriteString("/cast [mod:shift]")
		buf.WriteString(req.Shift)
		buf.WriteString("\r\n")
	}
	if req.Alt != "" {
		buf.WriteString("/cast [mod:alt]")
		buf.WriteString(req.Alt)
		buf.WriteString("\r\n")
	}
	if req.Ctrl != "" {
		buf.WriteString("/cast [mod:ctrl]")
		buf.WriteString(req.Ctrl)
		buf.WriteString("\r\n")
	}
	if req.Player != "" {
		buf.WriteString("/cast [@player]")
		buf.WriteString(req.Player)
		buf.WriteString("\r\n")
	}
	if req.Def != "" {
		buf.WriteString("/cast ")
		buf.WriteString(req.Def)
		buf.WriteString("\r\n")
	}

	modules.CreateLoginLog(c, "macro_combineSkills", 2)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": buf.String(),
	})
}
