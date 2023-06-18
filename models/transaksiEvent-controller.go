package models

import "gorm.io/gorm"

type TransaksiEvent struct {
	gorm.Model
	KeranjangID uint      `json:"keranjang_id" form:"keranjang_id"`
	Keranjang   Keranjang `gorm:"foreignKey:KeranjangID"`
	Description string    `json:"description" form:"description"`
	Status      string    `gorm:"type:ENUM('unpaid', 'paid', 'canceled')"`
}