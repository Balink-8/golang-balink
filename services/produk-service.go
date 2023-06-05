package services

import (
	"capstone/models"
	"capstone/repositories"
)

type ProdukService interface {
	GetProduksService(page int, limit int) ([]*models.Produk, int, error)
	GetProdukService(id string) (*models.Produk, error)
	CreateService(Produk models.Produk) (*models.Produk, error)
	UpdateService(id string, ProdukBody models.Produk) (*models.Produk, error)
	DeleteService(id string) error
}

type produkService struct {
	ProdukR repositories.ProdukRepository
}

func NewProdukService(ProdukR repositories.ProdukRepository) ProdukService {
	return &produkService{
		ProdukR: ProdukR,
	}
}

func (p *produkService) GetProduksService(page int, limit int) ([]*models.Produk, int, error) {
	Produks, totalData, err := p.ProdukR.GetProduksRepository(page, limit)
	if err != nil {
		return nil, 0, err
	}

	return Produks, totalData, nil
}

func (pr *produkService) GetProdukService(id string) (*models.Produk, error) {
	Produk, err := pr.ProdukR.GetProdukRepository(id)
	if err != nil {
		return nil, err
	}

	return Produk, nil
}

func (pr *produkService) CreateService(Produk models.Produk) (*models.Produk, error) {
	ProdukR, err := pr.ProdukR.CreateRepository(Produk)
	if err != nil {
		return nil, err
	}

	return ProdukR, nil
}

func (pr *produkService) UpdateService(id string, ProdukBody models.Produk) (*models.Produk, error) {
	Produk, err := pr.ProdukR.UpdateRepository(id, ProdukBody)
	if err != nil {
		return Produk, err
	}

	return Produk, nil
}

func (pr *produkService) DeleteService(id string) error {
	err := pr.ProdukR.DeleteRepository(id)
	if err != nil {
		return err
	}

	return nil
}
