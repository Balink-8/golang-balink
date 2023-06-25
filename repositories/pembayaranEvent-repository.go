package repositories

import (
	"capstone/models"

	"gorm.io/gorm"
)

type PembayaranEventRepository interface {
	CreatePembayaranEvent(PembayaranEvent models.PembayaranEvent) (*models.PembayaranEvent, error)
	UpdateRepository(pembayaranEvent models.PembayaranEvent) error
	GetPembayaranEventRepository(id int) (*models.PembayaranEvent, error)
}

type pembayaranEventRepository struct {
	DB *gorm.DB
}

func NewPembayaranEventRepository(db *gorm.DB) *pembayaranEventRepository {
	return &pembayaranEventRepository{
		DB: db,
	}
}

func (p *pembayaranEventRepository) CreatePembayaranEvent(pembayaranEvent models.PembayaranEvent) (*models.PembayaranEvent, error) {
	if err := p.DB.Create(&pembayaranEvent).Error; err != nil {
		return nil, err
	}

	return &pembayaranEvent, nil
}
func (p *pembayaranEventRepository) UpdateRepository(pembayaranEvent models.PembayaranEvent) error {

	err := p.DB.Preload("Keranjang").Preload("BuktiPembayaran").Updates(&pembayaranEvent).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *pembayaranEventRepository) GetPembayaranEventRepository(id int) (*models.PembayaranEvent, error) {
	var PembayaranEvent *models.PembayaranEvent

	if err := p.DB.Preload("Keranjang").Preload("MetodePembayaran").Where("id = ?", id).First(&PembayaranEvent).Error; err != nil {
		return nil, err
	}

	return PembayaranEvent, nil
}
