package models

import "gorm.io/gorm"

type Keranjang struct {
	gorm.Model
	User_ID   string `json:"user_id" form:"user_id"`
	Produk_ID string `json:"produk_id" form:"produk_id"`
	Jumlah    int64  `json:"jumlah" form:"jumlah"`
}
