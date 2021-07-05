package database

import "time"

type Macros struct {
	ID           int64     `gorm:"column:id" json:"id"`
	UpdateTime   time.Time `gorm:"column:updatetime" json:"updatetime"`
	IsVerify     uint8     `gorm:"column:is_verify" json:"isVerify"`
	MasteryID    int64     `gorm:"column:mastery_id" json:"masteryId"`
	ProfessionID int64     `gorm:"column:profession_id" json:"professionId"`

	SimpleMacro
}

type SimpleMacro struct {
	Title  string `gorm:"column:title" json:"title"`
	Macro  string `gorm:"column:macro" json:"macro"`
	Author string `gorm:"column:author" json:"author"`
}

type Profession struct {
	PID     int64 `gorm:"column:pid" json:"pid"`
	Version int8  `gorm:"column:version" json:"version"`
	SimpleProfession
}

type SimpleProfession struct {
	ID   int64  `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}
