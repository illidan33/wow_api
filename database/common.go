package database

import "time"

type ApiUnit struct {
	// 主键ID
	ID uint64 `grom:"primary_key;column:id" json:"id"`
	// 父级ID
	ParentID uint64 `grom:"column:parent_id" json:"parent_id"`
	// 英文名称
	Name string `grom:"column:name" json:"name"`
	// 中文名称
	NameCn string `grom:"column:name_cn" json:"name_cn"`
	// 描述
	Desc string `grom:"column:desc" json:"desc"`
	// 删除标识
	Enabled uint8 `grom:"column:enabled" json:"enabled"`
	// 创建时间
	CreateTime time.Time `grom:"column:create_time" json:"createTime"`
	// 更新时间
	UpdateTime time.Time `grom:"column:update_time" json:"updateTime"`
}
