package models

import "gorm.io/gorm"

type Transaksi struct {
	gorm.Model
	Keranjang_ID		string `json:"keranjang_id" form:"keranjang_id"`	
	Deskripsi			string `json:"deskripsi" form:"deskripsi"`
	Status_Transaksi 	string `json:"Status_Transaksi" form:"Status_Transaksi"`
}