package models

import "gorm.io/gorm"

type Artikel struct {
	gorm.Model
	Judul     string `json:"judul" form:"judul"`
	Deskripsi string `json:"deskripsi" form:"deskripsi"`
	Image     string `json:"image" from:"image"`
}
