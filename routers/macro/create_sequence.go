package macro

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/illidan33/wow_api/global"
	"github.com/illidan33/wow_api/modules"
	"math"
	"net/http"
	"strings"
)

func CreateSequence(c *gin.Context) {
	sequenceList := make([]modules.MacroSequence, 0)
	err := c.Bind(&sequenceList)
	if err != nil {
		modules.Return(c, 500, errors.New("params is error"))
		return
	}
	global.Config.Log.Debugf("CreateSequence req: %+v", sequenceList)

	for i, value := range sequenceList {
		if value.Cooldown == 0 {
			sequenceList[i].Cooldown = 100
		}
		sequenceList[i].SkillName = strings.Replace(value.SkillName, " ", "", -1)
		sequenceList[i].SkillName = strings.Replace(value.SkillName, "\n", "", -1)
	}

	macros, maxTime := modules.CreateSequence(sequenceList)
	maxTime = int(math.Ceil(float64(maxTime) / 100))
	macroText := fmt.Sprintf("#showtooltip <br>/castsequence reset=%d %s", maxTime, strings.Join(macros, ","))
	c.JSON(http.StatusOK, gin.H{"code": 0, "text": macroText, "desc": fmt.Sprintf("- 最后一次按键【%d】秒后，将重置<br>- 技能按照左侧循环，时间和技能顺序可以自己修改！", maxTime)})
}
