package macro60

import (
	"github.com/gin-gonic/gin"
	"github.com/illidan33/wow_tools/database"
	"github.com/illidan33/wow_tools/modules"
	"strconv"
)

type MacroListReq struct {
	ID           string `json:"id"`
	IsVerify     string `json:"isVerify"`
	MasteryID    string `json:"masteryId"`
	ProfessionID string `json:"professionId"`
	PageNo       string `json:"pageNo"`
	PageSize     string `json:"pageSize"`
}

func MacroList(c *gin.Context) {
	modules.CreateLoginLog(c, "macro_list", 2)

	var err error
	req := MacroListReq{}
	req.ID = c.DefaultQuery("id", "")
	req.IsVerify = c.DefaultQuery("isVerify", "1")
	req.ProfessionID = c.DefaultQuery("professionId", "0")
	req.MasteryID = c.DefaultQuery("masteryId", "0")
	req.PageNo = c.DefaultQuery("pageNo", "1")
	req.PageSize = c.DefaultQuery("pageSize", "10")

	v := c.DefaultQuery("v", "")
	if v == "v" {
		req.PageSize = "-1"
	}

	query := modules.QueryFilter{}
	if req.ID != "" {
		query.FilterParams("id", req.ID)
	}
	if req.IsVerify != "0" {
		query.FilterParams("is_verify", req.IsVerify)
	}
	if req.ProfessionID != "0" {
		query.FilterParams("profession_id", req.ProfessionID)
	}
	if req.MasteryID != "0" {
		query.FilterParams("mastery_id", req.MasteryID)
	}

	if query.QueryString == "" {
		modules.Return(c, 500, "bad request")
		return
	}

	pageNo, _ := strconv.ParseInt(req.PageNo, 10, 64)
	pageSize, _ := strconv.ParseInt(req.PageSize, 10, 64)

	macros := make([]database.MacrosOld60, 0)
	if req.PageSize != "-1" {
		err = modules.DbConn.Where(query.QueryString, query.QueryParams...).Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&macros).Error
	} else {
		err = modules.DbConn.Where(query.QueryString, query.QueryParams...).Find(&macros).Error
	}
	if err != nil {
		modules.Return(c, 500, err)
	} else {
		if v == "v" {
			modules.ReturnPage(c, 0, pageNo, pageSize, macros)
		} else {
			jsonMacros := make([]database.SimpleMacro60, 0)
			for _, v := range macros {
				jsonMacros = append(jsonMacros, v.SimpleMacro60)
			}
			modules.ReturnPage(c, 0, pageNo, pageSize, jsonMacros)
		}
	}
}
