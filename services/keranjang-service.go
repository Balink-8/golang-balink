package services

import (
	"capstone/models"
	"capstone/repositories"
)

type KeranjangService interface {
	GetKeranjangsService(page int, limit int, order string, search string) ([]*models.Keranjang, int, error)
	GetKeranjangService(id string) (*models.Keranjang, error)
	CreateService(Keranjang models.Keranjang) (*models.Keranjang, error)
	UpdateService(id string, KeranjangBody models.Keranjang) (*models.Keranjang, error)
	DeleteService(id string) error
	GetKeranjangByUserService(id string) ([]*models.Keranjang, error)
}

type keranjangService struct {
	KeranjangR repositories.KeranjangRepository
	ProdukR    repositories.ProdukRepository
}

func NewKeranjangService(KeranjangR repositories.KeranjangRepository, ProdukR repositories.ProdukRepository) KeranjangService {
	return &keranjangService{
		KeranjangR: KeranjangR,
		ProdukR:    ProdukR,
	}
}

func (k *keranjangService) GetKeranjangsService(page int, limit int, order string, search string) ([]*models.Keranjang, int, error) {
	Keranjangs, totalData, err := k.KeranjangR.GetKeranjangsRepository(page, limit, order, search)
	if err != nil {
		return nil, 0, err
	}

	return Keranjangs, totalData, nil
}

func (k *keranjangService) GetKeranjangService(id string) (*models.Keranjang, error) {
	Keranjang, err := k.KeranjangR.GetKeranjangRepository(id)
	if err != nil {
		return nil, err
	}

	return Keranjang, nil
}

func (k *keranjangService) CreateService(Keranjang models.Keranjang) (*models.Keranjang, error) {
	KeranjangR, err := k.KeranjangR.CreateRepository(Keranjang)
	if err != nil {
		return nil, err
	}

	return &KeranjangR, nil
}

func (k *keranjangService) UpdateService(id string, KeranjangBody models.Keranjang) (*models.Keranjang, error) {
	Keranjang, err := k.KeranjangR.UpdateRepository(id, KeranjangBody)
	if err != nil {
		return Keranjang, err
	}

	return Keranjang, nil
}

func (k *keranjangService) DeleteService(id string) error {
	err := k.KeranjangR.DeleteRepository(id)
	if err != nil {
		return err
	}

	return nil
}

func (k *keranjangService) GetKeranjangByUserService(id_user string) ([]*models.Keranjang, error) {
	Keranjangs, err := k.KeranjangR.GetKeranjangByUserRepository(id_user)
	if err != nil {
		return nil, err
	}
	return Keranjangs, nil
}
