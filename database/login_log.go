package database

type ApiLoginLog struct {
	ID         int    `json:"id" gorm:"column:id"`
	IP         string `json:"ip" gorm:"column:ip"`
	Method     string `json:"method" gorm:"column:method"`
	LoginDate  string `json:"loginDate" gorm:"column:login_date"`
	Count      int    `json:"count" gorm:"column:count"`
	Type       uint8  `gorm:"column:type" json:"type"`
	CreateTime string `json:"createTime" gorm:"column:create_time"`
	UpdateTime string `json:"updateTime" gorm:"column:update_time"`
}
