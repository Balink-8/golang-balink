package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Keranjang_ID     string `json:"keranjang_id" form:"keranjang_id"`
	Deskripsi        string `json:"deskripsi" form:"deskripsi"`
	Status_Transaksi string `json:"status_transaksi" form:"status_transaksi"`
}

// Transaction product model
type TransactionProduct struct {
	ID         uint64 `gorm:"primary_key;auto_increment"`
	ProductID  uint64 `gorm:"not null"`
	CustomerID uint64 `gorm:"not null"`
	Quantity   int    `gorm:"not null"`
}

// Ticket model
type Ticket struct {
	ID         uint64 `gorm:"primary_key;auto_increment"`
	TicketID   uint64 `gorm:"not null"`
	CustomerID uint64 `gorm:"not null"`
	Status     string `gorm:"not null"`
	Products   []TransactionProduct `gorm:"foreignkey:TicketID"`
}
