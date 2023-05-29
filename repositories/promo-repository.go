package repositories

import (
	"capstone/models"

	"gorm.io/gorm"
)

type PromoRepository interface {
	GetPromosRepository() ([]*models.Promo, error)
	GetPromoRepository(id string) (*models.Promo, error)
	CreateRepository(Promo models.Promo) (*models.Promo, error)
	UpdateRepository(id string, PromoBody models.Promo) (*models.Promo, error)
	DeleteRepository(id string) error
}

type promoRepository struct {
	DB *gorm.DB
}

func NewPromoRepository(DB *gorm.DB) PromoRepository {
	return &promoRepository{
		DB: DB,
	}
}

func (p *promoRepository) GetPromosRepository() ([]*models.Promo, error) {
	var Promos []*models.Promo

	if err := p.DB.Find(&Promos).Error; err != nil {
		return nil, err
	}

	return Promos, nil
}

func (p *promoRepository) GetPromoRepository(id string) (*models.Promo, error) {
	var Promo *models.Promo

	if err := p.DB.Where("ID = ?", id).Take(&Promo).Error; err != nil {
		return nil, err
	}

	return Promo, nil
}

func (p *promoRepository) CreateRepository(Promo models.Promo) (*models.Promo, error) {
	if err := p.DB.Save(&Promo).Error; err != nil {
		return nil, err
	}

	return &Promo, nil
}

func (p *promoRepository) UpdateRepository(id string, PromoBody models.Promo) (*models.Promo, error) {
	Promo, err := p.GetPromoRepository(id)
	if err != nil {
		return nil, err
	}

	err = p.DB.Where("ID = ?", id).Updates(models.Promo{Kode: PromoBody.Kode, Potongan_Harga: PromoBody.Potongan_Harga}).Error
	if err != nil {
		return nil, err
	}

	Promo.Kode = PromoBody.Kode
	Promo.Potongan_Harga = PromoBody.Potongan_Harga

	return Promo, nil
}

func (p *promoRepository) DeleteRepository(id string) error {
	_, err := p.GetPromoRepository(id)
	if err != nil {
		return err
	}

	if err := p.DB.Delete(&models.Promo{}, id).Error; err != nil {
		return err
	}

	return nil
}
