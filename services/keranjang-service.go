package services

import (
	"capstone/models"
	"capstone/repositories"
)

type KeranjangService interface {
	GetKeranjangsService() ([]*models.Keranjang, error)
	GetKeranjangService(id string) (*models.Keranjang, error)
	CreateService(Keranjang models.Keranjang) (*models.Keranjang, error)
	UpdateService(id string, KeranjangBody models.Keranjang) (*models.Keranjang, error)
	DeleteService(id string) error
	GetKeranjangByUserService(id string) ([]*models.Keranjang, error)
}

type keranjangService struct {
	KeranjangR repositories.KeranjangRepository
}

func NewKeranjangService(KeranjangR repositories.KeranjangRepository) KeranjangService {
	return &keranjangService{
		KeranjangR: KeranjangR,
	}
}

func (k *keranjangService) GetKeranjangsService() ([]*models.Keranjang, error) {
	Keranjangs, err := k.KeranjangR.GetKeranjangsRepository()
	if err != nil {
		return nil, err
	}

	return Keranjangs, nil
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

	return KeranjangR, nil
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
