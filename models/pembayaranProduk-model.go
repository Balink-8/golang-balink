package models

import (
	"gorm.io/gorm"
)

type PembayaranProduk struct {
	gorm.Model
	KeranjangID      int64 `json:"keranjang_id" form:"keranjang_id"`
	Keranjang        Keranjang
	Status           string `json:"status" form:"status"`
	AlamatPengiriman string `json:"alamat_pengiriman" form:"alamat_pengiriman"`
	Pesan            string `json:"pesan" form:"pesan"`
	PromoID          int64  `json:"promo_id" form:"promo_id"`
	MetodePembayaran int64 `json:"metode_pembayaran" form:"metode_pembayaran"`
}
