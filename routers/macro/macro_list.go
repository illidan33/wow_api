package macro

import (
	"github.com/gin-gonic/gin"
	"github.com/illidan33/wow_api/database"
	"github.com/illidan33/wow_api/modules"
)

func MacroList(c *gin.Context) {
	isVerify := c.Query("isVerify")
	pid := c.Query("professionId")
	id := c.Query("id")

	query := modules.QueryFilter{}
	if id != "" {
		query.FilterParams("id", id)
	}
	if isVerify != "" {
		query.FilterParams("is_verify", isVerify)
	}
	if pid != "" && pid != "0" {
		query.FilterParams("profession_id", pid)
	}

	if query.QueryString == "" {
		modules.Return(c, 500, "bad request")
		return
	}

	macros := make([]database.Macros, 0)
	err := modules.DbConn.Where(query.QueryString, query.QueryParams...).Find(&macros).Error
	if err != nil {
		modules.Return(c, 500, err)
	} else {
		modules.Return(c, 0, macros)
	}
}
