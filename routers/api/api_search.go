package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/illidan33/wow_tools/modules"
	"net/http"
)

func ApiSearch(c *gin.Context) {
	s := c.DefaultQuery("s", "")

	if len(s) <= 3 {
		modules.Return(c, 500, errors.New("关键词过短"))
		return
	}

	apis, err := modules.GetApiListBySearchText(s)
	if err != nil {
		modules.Return(c, 500, err)
		return
	}

	if len(apis) > 30 {
		modules.Return(c, 500, errors.New("关键词过短，搜索超时！"))
		return
	}

	modules.CreateLoginLog(c, fmt.Sprintf("api_search_%s", s), 2)
	modules.Return(c, http.StatusOK, apis)
}
