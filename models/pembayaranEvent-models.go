package models

import (
	"gorm.io/gorm"
)

type PembayaranEvent struct {
	gorm.Model
	User_ID            string `json:"user_id" form:"user_id"`
	Event_ID           string `json:"event_id" form:"event_id"`
	KodePembayaran     string `json:"kode_pembayaran" form:"kode_pembayaran"`
	BuktiPembayaran    string `json:"bukti_pembayaran" form:"bukti_pembayaran"`
	Status             string `json:"status" form:"status"`
	Qty                int64  `json:"qty" form:"qty"`
	Total             int64  `json:"total" form:"total"`
	PromoID            int64  `json:"promo_id" form:"promo_id"`
	MetodePembayaranID int64  `json:"metode_pembayaran_id" form:"metode_pembayaran"`
	MetodePembayaran   MetodePembayaran
}
