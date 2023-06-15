package services

import (
	"capstone/models"
	"capstone/repositories"
)

type KategoriProdukService interface {
	GetKategoriProduksService(page int, limit int, order string, search string) ([]*models.KategoriProduk, int, error)
	GetKategoriProdukService(id string) (*models.KategoriProduk, error)
	CreateService(KategoriProduk models.KategoriProduk) (*models.KategoriProduk, error)
	UpdateService(id string, KategoriProdukBody models.KategoriProduk) (*models.KategoriProduk, error)
	DeleteService(id string) error
}

type kategoriProdukService struct {
	KategoriProdukR repositories.KategoriProdukRepository
}

func NewKategoriProdukService(KategoriProdukR repositories.KategoriProdukRepository) KategoriProdukService {
	return &kategoriProdukService{
		KategoriProdukR: KategoriProdukR,
	}
}

func (k *kategoriProdukService) GetKategoriProduksService(page int, limit int, order string, search string) ([]*models.KategoriProduk, int, error) {
	KategoriProduks, totalData, err := k.KategoriProdukR.GetKategoriProduksRepository(page, limit, order, search)
	if err != nil {
		return nil, 0, err
	}

	return KategoriProduks, totalData, nil
}

func (k *kategoriProdukService) GetKategoriProdukService(id string) (*models.KategoriProduk, error) {
	KategoriProduk, err := k.KategoriProdukR.GetKategoriProdukRepository(id)
	if err != nil {
		return nil, err
	}

	return KategoriProduk, nil
}

func (k *kategoriProdukService) CreateService(KategoriProduk models.KategoriProduk) (*models.KategoriProduk, error) {
	KategoriProdukR, err := k.KategoriProdukR.CreateRepository(KategoriProduk)
	if err != nil {
		return nil, err
	}

	return KategoriProdukR, nil
}

func (k *kategoriProdukService) UpdateService(id string, KategoriProdukBody models.KategoriProduk) (*models.KategoriProduk, error) {
	KategoriProduk, err := k.KategoriProdukR.UpdateRepository(id, KategoriProdukBody)
	if err != nil {
		return KategoriProduk, err
	}

	return KategoriProduk, nil
}

func (k *kategoriProdukService) DeleteService(id string) error {
	err := k.KategoriProdukR.DeleteRepository(id)
	if err != nil {
		return err
	}

	return nil
}
