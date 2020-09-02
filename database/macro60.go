package database

import "time"

type MacrosOld60 struct {
	SimpleMacro60
	ID           int64     `gorm:"column:id" json:"id"`
	UpdateTime   time.Time `gorm:"column:updatetime" json:"updatetime"`
	IsVerify     uint8     `gorm:"column:is_verify" json:"isVerify"`
	MasteryID    int64     `gorm:"column:mastery_id" json:"masteryId"`
	ProfessionID int64     `gorm:"column:profession_id" json:"professionId"`
}

type SimpleMacro60 struct {
	Title  string `gorm:"column:title" json:"title"`
	Macro  string `gorm:"column:macro" json:"macro"`
	Author string `gorm:"column:author" json:"author"`
}
