package repositories

import (
	"capstone/models"

	"gorm.io/gorm"
)

type PembayaranProdukRepository interface {
	CreatePembayaranProduk(PembayaranProduk models.PembayaranProduk) (*models.PembayaranProduk, error)
	UpdateRepository(pembayaranProduk models.PembayaranProduk) error
	GetPembayaranProdukRepository(id int) (*models.PembayaranProduk, error)
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
func (p *pembayaranProdukRepository) UpdateRepository(pembayaranProduk models.PembayaranProduk) error {

	err := p.DB.Preload("Keranjang").Preload("BuktiPembayaran").Updates(&pembayaranProduk).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *pembayaranProdukRepository) GetPembayaranProdukRepository(id int) (*models.PembayaranProduk, error) {
	var PembayaranProduk *models.PembayaranProduk

	if err := p.DB.Preload("Keranjang").Preload("MetodePembayaran").Where("id = ?", id).First(&PembayaranProduk).Error; err != nil {
		return nil, err
	}

	return PembayaranProduk, nil
}
