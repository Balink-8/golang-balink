package models

import "gorm.io/gorm"

type Promo struct {
	gorm.Model
	Kode			string `json:"kode" form:"kode"`
	Potongan_Harga	string `json:"Potongan_Harga" form:"Potongan_Harga"`
}
