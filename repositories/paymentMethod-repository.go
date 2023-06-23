package repositories

import (
	"capstone/models"

	"gorm.io/gorm"
)

type PaymentMethodRepository interface {
	GetPaymentMethodRepository(id int64) (*models.MetodePembayaran, error)
}

type paymentMethodRepository struct {
	DB *gorm.DB
}

func NewPaymentMethodRepository(DB *gorm.DB) PaymentMethodRepository {
	return &paymentMethodRepository{
		DB: DB,
	}
}

func (p *paymentMethodRepository) GetPaymentMethodRepository(id int64) (*models.MetodePembayaran, error) {
	var paymentMethod *models.MetodePembayaran

	if err := p.DB.Where("ID = ?", id).Take(&paymentMethod).Error; err != nil {
		return nil, err
	}

	return paymentMethod, nil
}
