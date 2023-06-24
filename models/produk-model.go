package models

import (
	"gorm.io/gorm"
)

type Produk struct {
	gorm.Model
	Kategori_ID string `json:"kategori_id" form:"kategori_id"`
	Nama        string `json:"nama" form:"nama"`
	Deskripsi   string `json:"deskripsi" form:"deskripsi"`
	Harga       int64  `json:"harga" form:"harga"`
	Stok        int64  `json:"stok" form:"stok"`
	Image       string `json:"image" from:"image"`
}
