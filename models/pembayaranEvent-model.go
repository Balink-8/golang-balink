package models

import (
	"gorm.io/gorm"
)

type PembayaranEvent struct {
	gorm.Model
	KeranjangID        string `json:"keranjang_id" form:"keranjang_id"`
	KeranjangTiket     Keranjang
	Status             string `json:"status" form:"status"`
	PromoID            int64  `json:"promo_id" form:"promo_id"`
	MetodePembayaranID int64  `json:"metode_pembayaran_id" form:"metode_pembayaran"`
	BuktiPembayaran    string `json:"bukti_pembayaran" form:"bukti_pembayaran"`
	MetodePembayaran   MetodePembayaran
}
