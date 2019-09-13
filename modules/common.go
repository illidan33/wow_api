package modules

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/illidan33/wow_api/database"
	"github.com/illidan33/wow_api/global"
	"net/http"
)

func Return(c *gin.Context, code int32, resp interface{}) {
	if e, ok := resp.(error); ok {
		global.Config.Log.Error(e)
		if e.Error() == "record not found" {
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  "Interner Error",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  e.Error(),
			})
		}

	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"data": resp,
		})
	}
}

func IsNotFound(err error) bool {
	if err.Error() == "record not found" {
		return true
	}
	return false
}

// 快速创建
type MacroSequence struct {
	// 名称
	SkillName string `json:"skillName"`
	// 优先级
	Level int `json:"level"`
	// 冷却时间(秒*100)
	Cooldown int `json:"cooldown"`
	// 剩余时间
	CDTime int `json:"-"`
}

func CreateSequence(temps []MacroSequence) (macroText []string, maxTime int) {
	for _, value := range temps {
		if maxTime == 0 {
			maxTime = value.Cooldown
		} else {
			if value.Cooldown > maxTime {
				maxTime = value.Cooldown
			}
		}
	}

	for i := 0; i < maxTime; i++ {
		coolIndex := 0
		coolLevel := 0
		for j, value := range temps {
			if value.CDTime != 0 {
				continue
			}
			if coolLevel == 0 {
				coolLevel = value.Level
				coolIndex = j
			} else {
				if coolLevel > value.Level {
					coolLevel = value.Level
					coolIndex = j
				}
			}
		}
		if coolLevel != 0 {
			macroText = append(macroText, temps[coolIndex].SkillName)
			temps[coolIndex].CDTime = temps[coolIndex].Cooldown
		}
		// CD时间减1
		for k, value := range temps {
			if value.CDTime != 0 {
				temps[k].CDTime -= 1
			}
		}
	}

	return macroText, maxTime
}

func CreateLoginLog(c *gin.Context, html string) {
	go UpdateOrCreateLog(c.ClientIP(), html)
}

func GetApiByParentID(tableType string, parentID string) (interface{}, error) {
	switch tableType {
	case "title-api":
		apiList := make([]database.ApiWow, 0)
		err := DbConn.Where("parent_id = ?", parentID).Find(&apiList).Error
		if err != nil {
			return nil, err
		}
		return apiList, nil
	case "title-macro":
		apiList := make([]database.ApiMacro, 0)
		err := DbConn.Where("parent_id = ?", parentID).Find(&apiList).Error
		if err != nil {
			return nil, err
		}
		return apiList, nil
	case "title-event":
		apiList := make([]database.ApiEvent, 0)
		err := DbConn.Where("parent_id = ?", 0).Find(&apiList).Error
		if err != nil {
			return nil, err
		}
		return apiList, nil
	case "title-widget":
		apiList := make([]database.ApiWidget, 0)
		err := DbConn.Where("parent_id = ?", parentID).Find(&apiList).Error
		if err != nil {
			return nil, err
		}
		return apiList, nil
	case "title-widgetHandler":
		apiList := make([]database.ApiWidgetHandler, 0)
		err := DbConn.Where("parent_id = ?", parentID).Find(&apiList).Error
		if err != nil {
			return nil, err
		}
		return apiList, nil
	default:
		return nil, errors.New("no such type")
	}
}

func GetApiByID(tableType string, id string) (interface{}, error) {
	switch tableType {
	case "title-wow-api":
		api := database.ApiWow{}
		err := DbConn.Where("id = ?", id).Find(&api).Error
		if err != nil {
			return nil, err
		}
		return api, nil
	case "title-wow-macro":
		api := database.ApiMacro{}
		err := DbConn.Where("id = ?", id).Find(&api).Error
		if err != nil {
			return nil, err
		}
		return api, nil
	case "title-wow-event":
		api := database.ApiEvent{}
		err := DbConn.Where("id = ?", id).Find(&api).Error
		if err != nil {
			return nil, err
		}
		return api, nil
	default:
		return nil, errors.New("no such type")
	}
}

type QueryFilter struct {
	QueryString string
	QueryParams []interface{}
}

func (query *QueryFilter) FilterParams(fieldName string, fieldValue interface{}) {
	if query.QueryString == "" {
		query.QueryString = fmt.Sprintf("%s = ?", fieldName)
	} else {
		query.QueryString = fmt.Sprintf("%s AND %s = ?", query.QueryString, fieldName)
	}
	query.QueryParams = append(query.QueryParams, fieldValue)
}
