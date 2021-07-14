package modules

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/illidan33/wow_api/database"
	"github.com/illidan33/wow_api/global"
	"net/http"
	"strings"
)

func Return(c *gin.Context, code int32, resp interface{}) {
	if e, ok := resp.(error); ok {
		global.Log.Error(e)
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

func ReturnPage(c *gin.Context, code int32, pageNo int64, pageSize int64, resp interface{}) {
	if e, ok := resp.(error); ok {
		global.Log.Error(e)
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
		if pageSize == -1 {
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"data": resp,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":   code,
				"pageNo": pageNo,
				"data":   resp,
			})
		}
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

func CreateLoginLog(c *gin.Context, html string, t uint8) {
	go UpdateOrCreateLog(c.ClientIP(), html, t)
}

// 获取项目的子项目
func GetApiByParentID(tableType string, parentID string) (interface{}, error) {
	simples := make([]database.SimpleApiItem, 0)
	t := GetApiTypeByTableType(tableType)
	apiList := make([]database.ApiItem, 0)

	err := DbConn.Where("parent_id = ? and type = ?", parentID, t).Find(&apiList).Error
	if err != nil {
		return nil, err
	}
	for _, api := range apiList {
		simple := database.SimpleApiItem{
			ID:     api.ApiID,
			Name:   api.Name,
			NameCn: api.NameCn,
			Desc:   api.Desc,
		}
		simples = append(simples, simple)
	}
	return simples, nil
}

// 转换前端的tableType为api类型
func GetApiTypeByTableType(tbType string) (t uint8) {
	switch tbType {
	case "api":
		t = 1
	case "macro":
		t = 3
	case "event":
		t = 2
	case "widget":
		t = 4
	case "widgetHandler":
		t = 5
	default:
		t = 0
	}
	return
}
func GetTbTypeByApiType(t uint8) (tbType string) {
	switch t {
	case 1:
		tbType = "api"
	case 3:
		tbType = "macro"
	case 2:
		t = 2
		tbType = "event"
	case 4:
		tbType = "widget"
	case 5:
		tbType = "widgetHandler"
	default:
		tbType = ""
	}
	return
}

// 获取详情
func GetApiByID(id string) (simpleApi database.SimpleApiItem, err error) {
	api := database.ApiItem{}
	err = DbConn.Where("api_id = ?", id).Find(&api).Error
	if err != nil {
		return
	}

	simpleApi = database.SimpleApiItem{
		ID:     api.ApiID,
		Name:   api.Name,
		NameCn: api.NameCn,
		Desc:   api.Desc,
	}
	return
}

func GetApiListBySearchText(s string) (simpleApis []database.SearchApiItem, err error) {
	apis := make([]database.ApiItem, 0)

	err = DbConn.Where("parent_id!=? and type != ? and name like ?", 0, 2, "%"+s+"%").Find(&apis).Error
	if err != nil {
		return
	}
	for _, api := range apis {
		simpleApi := database.SearchApiItem{
			SimpleApiItem: database.SimpleApiItem{
				ID:     api.ApiID,
				Name:   api.Name,
				NameCn: api.NameCn,
				Desc:   api.Desc,
			},
			Type: GetTbTypeByApiType(api.Type),
		}
		simpleApis = append(simpleApis, simpleApi)
	}
	events := make([]database.ApiItem, 0)
	err = DbConn.Where("type = ? and name like ?", 2, "%"+s+"%").Find(&events).Error
	if err != nil {
		return
	}
	for _, api := range events {
		simpleApi := database.SearchApiItem{
			SimpleApiItem: database.SimpleApiItem{
				ID:     api.ApiID,
				Name:   api.Name,
				NameCn: api.NameCn,
				Desc:   api.Desc,
			},
			Type: "event",
		}
		simpleApis = append(simpleApis, simpleApi)
	}
	return
}

func GetApiDetailUrlByID(tableType string, name string) (url string) {
	switch tableType {
	case "api":
		url = fmt.Sprintf("https://wow.gamepedia.com/API_%s", name)
	case "macro":
		url = fmt.Sprintf("https://wow.gamepedia.com/MACRO_%s", name)
	case "event":
		url = fmt.Sprintf("https://wow.gamepedia.com/%s", name)
	case "widget":
		nameArr := strings.Split(name, "：")
		if len(nameArr) > 1 {
			url = fmt.Sprintf("https://wow.gamepedia.com/API_%s_%s", nameArr[0], nameArr[1])
		} else {
			nameArr := strings.Split(name, ":")
			if len(nameArr) > 1 {
				url = fmt.Sprintf("https://wow.gamepedia.com/API_%s_%s", nameArr[0], nameArr[1])
			} else {
				url = fmt.Sprintf("https://wow.gamepedia.com/API_%s", name)
			}
		}
	case "widgetHandler":
		url = fmt.Sprintf("https://wow.gamepedia.com/UIHANDLER_%s", name)
	default:
		url = ""
	}
	return
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
