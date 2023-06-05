package services

import (
	"capstone/models"
	"capstone/repositories"
)

type TransaksiService interface {
	GetTransaksisService() ([]*models.Transaksi, error)
	GetTransaksiService(id string) (*models.Transaksi, error)
	CreateService(Transaksi models.Transaksi) (*models.Transaksi, error)
	UpdateService(id string, TransaksiBody models.Transaksi) (*models.Transaksi, error)
	DeleteService(id string) error
}

type transaksiService struct {
	TransaksiR repositories.TransaksiRepository
}

func NewTransaksiService(TransaksiR repositories.TransaksiRepository) TransaksiService {
	return &transaksiService{
		TransaksiR: TransaksiR,
	}
}

func (t *transaksiService) GetTransaksisService() ([]*models.Transaksi, error) {
	Transaksis, err := t.TransaksiR.GetTransaksisRepository()
	if err != nil {
		return nil, err
	}

	return Transaksis, nil
}

func (t *transaksiService) GetTransaksiService(id string) (*models.Transaksi, error) {
	Transaksi, err := t.TransaksiR.GetTransaksiRepository(id)
	if err != nil {
		return nil, err
	}

	return Transaksi, nil
}

func (t *transaksiService) CreateService(Transaksi models.Transaksi) (*models.Transaksi, error) {
	TransaksiR, err := t.TransaksiR.CreateRepository(Transaksi)
	if err != nil {
		return nil, err
	}

	return TransaksiR, nil
}

func (t *transaksiService) UpdateService(id string, TransaksiBody models.Transaksi) (*models.Transaksi, error) {
	Transaksi, err := t.TransaksiR.UpdateRepository(id, TransaksiBody)
	if err != nil {
		return Transaksi, err
	}

	return Transaksi, nil
}

func (t *transaksiService) DeleteService(id string) error {
	err := t.TransaksiR.DeleteRepository(id)
	if err != nil {
		return err
	}

	return nil
}
