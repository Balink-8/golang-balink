package models

import "gorm.io/gorm"

type Masalah struct {
	gorm.Model
	User_ID	string `json:"user_id" form:"user_id"`
	Isi		string `json:"isi" form:"isi"`
}