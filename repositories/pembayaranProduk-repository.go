package repositories

import (
	"capstone/models"

	"gorm.io/gorm"
)

type PembayaranProdukRepository interface {
	CreatePembayaranProduk(PembayaranProduk models.PembayaranProduk) (*models.PembayaranProduk, error)
}

type pembayaranProdukRepository struct {
	DB *gorm.DB
}

func NewPembayaranProdukRepository(db *gorm.DB) *pembayaranProdukRepository {
	return &pembayaranProdukRepository{
		DB: db,
	}
}

func (p *pembayaranProdukRepository) CreatePembayaranProduk(pembayaranProduk models.PembayaranProduk) (*models.PembayaranProduk, error) {
	if err := p.DB.Create(&pembayaranProduk).Error; err != nil {
		return nil, err
	}

	return &pembayaranProduk, nil
}
