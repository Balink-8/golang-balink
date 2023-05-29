package services

import (
	"capstone/models"
	"capstone/repositories"
)

type PromoService interface {
	GetPromosService() ([]*models.Promo, error)
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

func (p *promoService) GetPromosService() ([]*models.Promo, error) {
	Promos, err := p.PromoR.GetPromosRepository()
	if err != nil {
		return nil, err
	}

	return Promos, nil
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
