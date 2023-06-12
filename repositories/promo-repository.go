package repositories

import (
	"capstone/models"

	"gorm.io/gorm"
)

type PromoRepository interface {
	GetPromosRepository(page int, limit int, order string) ([]*models.Promo, int, error)
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

func (p *promoRepository) GetPromosRepository(page int, limit int, order string) ([]*models.Promo, int, error) {
	var Promos []*models.Promo
	var totalData int64

	if err := p.DB.Model(&models.Promo{}).Count(&totalData).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	query := p.DB.Offset(offset).Limit(limit)

	switch order {
	case "asc":
		query = query.Order("ID ASC")
	case "desc":
		query = query.Order("ID DESC")
	}

	if err := query.Find(&Promos).Error; err != nil {
		return nil, 0, err
	}

	return Promos, int(totalData), nil
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

	err = p.DB.Where("ID = ?", id).Updates(models.Promo{Kode: PromoBody.Kode, PotonganHarga: PromoBody.PotonganHarga}).Error
	if err != nil {
		return nil, err
	}

	Promo.Kode = PromoBody.Kode
	Promo.PotonganHarga = PromoBody.PotonganHarga

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
