package models

import (
	"gorm.io/gorm"
)

type PembayaranEvent struct {
	gorm.Model
	UserID             string `json:"user_id" form:"keranjang_id"`
	EventID            string `json:"event_id" form:"event_id"`
	Event              Event
	Status             string `json:"status" form:"status"`
	Qty                int    `json:"qty" form:"qty"`
	Total             int `json:"total" form:"total"`
	PromoID            int64  `json:"promo_id" form:"promo_id"`
	MetodePembayaranID int64  `json:"metode_pembayaran_id" form:"metode_pembayaran"`
	BuktiPembayaran    string `json:"bukti_pembayaran" form:"bukti_pembayaran"`
	MetodePembayaran   MetodePembayaran
}
