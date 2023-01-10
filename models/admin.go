package models

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Id       uint   `gorm:"primaryKey" json:"id" xml:"id" form:"id"`
	Username string `json:"username" xml:"username" form:"username" `
	Password string `json:"password" xml:"password" form:"password" `
}
