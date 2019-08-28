package database

type ApiEvent struct {
	ApiUnit
}

type ApiMacro struct {
	ApiUnit
}

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

type ApiWow struct {
	ApiUnit
}

type ApiWidget struct {
	ApiUnit
}

type ApiWidgetHandler struct {
	ApiUnit
}
