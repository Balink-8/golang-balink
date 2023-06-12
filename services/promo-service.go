package services

import (
	"capstone/models"
	"capstone/repositories"
)

type PromoService interface {
	GetPromosService(page int, limit int, order string) ([]*models.Promo, int, error)
	GetPromoService(id string) (*models.Promo, error)
	CreateService(Promo models.Promo) (*models.Promo, error)
	UpdateService(id string, PromoBody models.Promo) (*models.Promo, error)
	DeleteService(id string) error
}

type promoService struct {
	PromoR repositories.PromoRepository
}

func NewPromoService(PromoR repositories.PromoRepository) PromoService {
	return &promoService{
		PromoR: PromoR,
	}
}

func (p *promoService) GetPromosService(page int, limit int, order string) ([]*models.Promo, int, error) {
	Promos, totalData, err := p.PromoR.GetPromosRepository(page, limit, order)
	if err != nil {
		return nil, 0, err
	}

	return Promos, totalData, nil
}

func (p *promoService) GetPromoService(id string) (*models.Promo, error) {
	Promo, err := p.PromoR.GetPromoRepository(id)
	if err != nil {
		return nil, err
	}

	return Promo, nil
}

func (p *promoService) CreateService(Promo models.Promo) (*models.Promo, error) {
	PromoR, err := p.PromoR.CreateRepository(Promo)
	if err != nil {
		return nil, err
	}

	return PromoR, nil
}

func (p *promoService) UpdateService(id string, PromoBody models.Promo) (*models.Promo, error) {
	Promo, err := p.PromoR.UpdateRepository(id, PromoBody)
	if err != nil {
		return Promo, err
	}

	return Promo, nil
}

func (p *promoService) DeleteService(id string) error {
	err := p.PromoR.DeleteRepository(id)
	if err != nil {
		return err
	}

	return nil
}
