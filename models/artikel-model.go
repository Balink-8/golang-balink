package models

import "gorm.io/gorm"

type Artikel struct {
	gorm.Model
	Gambar    string `json:"gambar" form:"gambar"`
	Judul     string `json:"judul" form:"judul"`
	Deskripsi string `json:"deskripsi" form:"deskripsi"`
	Image     string
}
