package database

import "time"

type MacrosOld60 struct {
	JsonMacroOld60
	MasteryID    int64 `gorm:"column:mastery_id" json:"masteryId"`
	ProfessionID int64 `gorm:"column:profession_id" json:"professionId"`
}

type JsonMacroOld60 struct {
	ID         int64     `gorm:"column:id" json:"id"`
	Title      string    `gorm:"column:title" json:"title"`
	Macro      string    `gorm:"column:macro" json:"macro"`
	UpdateTime time.Time `gorm:"column:updatetime" json:"updatetime"`
	Author     string    `gorm:"column:author" json:"author"`
	IsVerify   uint8     `gorm:"column:is_verify" json:"isVerify"`
}
