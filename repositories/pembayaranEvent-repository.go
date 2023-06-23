package repositories

import (
	"capstone/models"

	"gorm.io/gorm"
)

type PembayaranEventRepository interface {
	CreatePembayaranEvent(PembayaranEvent models.PembayaranEvent) (*models.PembayaranEvent, error)
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
