package models

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	Id           string `json:"id" gorm:"primaryKey" xml:"id" form:"id"`
	FirstNameLao string `json:"firstNameLao" xml:"firstnNameLao" form:"firstNameLao"`
	LastNameLao  string `json:"lastNameLao" xml:"lastNameLao" form:"lastNameLao"`
	FirstNameENG string `json:"firstNameENG" xml:"firstNameENG" form:"firstNameENG"`
	LastNameENG  string `json:"lastNameENG" xml:"lastNameENG" form:"lastNameENG"`
	Class        string `json:"class" xml:"class" form:"class"`
	Tell         string `json:"tell" xml:"tell" form:"tell"`
	Email        string `json:"email" xml:"email" form:"email"`
	Role         string `json:"role" xml:"role" form:"role"`
	Goal         string `json:"goal" xml:"goal" form:"goal"`
	MagicWord    string `json:"magicWord" xml:"magicWord" form:"magicWord"`
	Image        string `json:"image" xml:"image" form:"image"`
}
