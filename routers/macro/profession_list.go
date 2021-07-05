package macro

import (
	"github.com/gin-gonic/gin"
	"github.com/illidan33/wow_tools/database"
	"github.com/illidan33/wow_tools/modules"
)

func ProfessionList(c *gin.Context) {
	modules.CreateLoginLog(c, "macro_professionList", 2)

	var err error
	version := c.DefaultQuery("v", "0")
	pid := c.DefaultQuery("pid", "0")

	query := modules.QueryFilter{}
	query.FilterParams("version", version)
	query.FilterParams("pid", pid)

	if query.QueryString == "" {
		modules.Return(c, 500, "bad request")
		return
	}

	professions := make([]database.Profession, 0)
	err = modules.DbConn.Where(query.QueryString, query.QueryParams...).Find(&professions).Error
	if err != nil {
		modules.Return(c, 500, err)
	} else {
		jsonList := make([]database.SimpleProfession, 0)
		for _, v := range professions {
			jsonList = append(jsonList, v.SimpleProfession)
		}
		modules.ReturnPage(c, 0, 0, -1, jsonList)
	}
}
