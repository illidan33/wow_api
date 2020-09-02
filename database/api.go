package database

import "time"

type ApiUnverify struct {
	ID         int    `gorm:"column:id" json:"id"`
	ApiID      int    `json:"apiID" gorm:"column:api_id"`
	Type       string `json:"type" gorm:"column:type"`
	Name       string `json:"name" gorm:"column:name"`
	NameCn     string `json:"nameCn" gorm:"column:name_cn"`
	Desc       string `json:"desc" gorm:"column:desc"`
	InfoDesc   string `json:"infoDesc" gorm:"column:info_desc"`
	CreateTime string `json:"createTime" gorm:"column:create_time"`
	IsHandle   uint8  `json:"isHandle" gorm:"column:is_handle"`
}

type ApiUnit struct {
	// 主键ID
	ID uint64 `grom:"primary_key;column:id" json:"id"`
	// 英文名称
	Name string `grom:"column:name" json:"name"`
	// 中文名称
	NameCn string `grom:"column:name_cn" json:"nameCn"`
	// 描述
	Desc string `grom:"column:desc" json:"desc"`
	// 父级ID
	ParentID uint64 `grom:"column:parent_id" json:"parentId"`
	// 删除标识
	Enabled uint8 `grom:"column:enabled" json:"enabled"`
	// 创建时间
	CreateTime time.Time `grom:"column:create_time" json:"createTime"`
	// 更新时间
	UpdateTime time.Time `grom:"column:update_time" json:"updateTime"`
}

// 合并数据之后
type ApiItem struct {
	ApiUnit
	ApiID int32 `gorm:"column:api_id" json:"apiId"`
	Type  uint8 `gorm:"column:type" json:"type"`
}

type SimpleApiItem struct {
	// api ID
	ID int32 `gorm:"column:id" json:"id"`
	// 英文名称
	Name string `grom:"column:name" json:"name"`
	// 中文名称
	NameCn string `grom:"column:name_cn" json:"nameCn"`
	// 描述
	Desc string `grom:"column:desc" json:"desc"`
}

type SearchApiItem struct {
	SimpleApiItem
	Type string `json:"type"`
}
