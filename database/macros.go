package database

import "time"

type Macros struct {
	JsonMacro
	MasteryID    int64 `gorm:"column:mastery_id" json:"masteryId"`
	ProfessionID int64 `gorm:"column:profession_id" json:"professionId"`
	IsVerify     uint8 `gorm:"column:is_verify" json:"isVerify"`
}

type JsonMacro struct {
	ID         int64     `gorm:"column:id" json:"id"`
	Title      string    `gorm:"column:title" json:"title"`
	Macro      string    `gorm:"column:macro" json:"macro"`
	UpdateTime time.Time `gorm:"column:updatetime" json:"updatetime"`
	Author     string    `gorm:"column:author" json:"author"`
}
