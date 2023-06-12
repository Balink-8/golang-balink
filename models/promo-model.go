package models

import "gorm.io/gorm"

type Promo struct {
	gorm.Model
	Kode			string `json:"kode" form:"kode"`
	PotonganHarga	int64 `json:"potongan_harga" form:"Potongan_Harga"`
}
