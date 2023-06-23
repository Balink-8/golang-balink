package models

import (
	"gorm.io/gorm"
)

type MetodePembayaran struct {
	gorm.Model
	Nama string `json:"nama" form:"nama"`
	VA   int64  `json:"va" form:"va"`
}
