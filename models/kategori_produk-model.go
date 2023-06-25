package models

import (
	"gorm.io/gorm"
)

type KategoriProduk struct {
	gorm.Model
	Nama      string `json:"nama" form:"nama"`
	Deskripsi string `json:"deskripsi" form:"deskripsi"`
}
