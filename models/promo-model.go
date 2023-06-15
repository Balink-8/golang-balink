package models

import "gorm.io/gorm"

type Promo struct {
	gorm.Model
	Nama			string `json:"nama" form:"nama"`
	Deskripsi		string `json:"deskripsi" form:"deskripsi"`
	Kode			string `json:"kode" form:"kode"`
	PotonganHarga	int64 `json:"potongan_harga" form:"Potongan_Harga"`
}
