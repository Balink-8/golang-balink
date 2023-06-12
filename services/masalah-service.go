package services

import (
	"capstone/models"
	"capstone/repositories"
)

type MasalahService interface {
	GetMasalahsService(page int, limit int, order string) ([]*models.Masalah, int, error)
	GetMasalahService(id string) (*models.Masalah, error)
	CreateService(Masalah models.Masalah) (*models.Masalah, error)
}

type masalahService struct {
	MasalahR repositories.MasalahRepository
}

func NewMasalahService(MasalahR repositories.MasalahRepository) MasalahService {
	return &masalahService{
		MasalahR: MasalahR,
	}
}

func (m *masalahService) GetMasalahsService(page int, limit int, order string) ([]*models.Masalah, int, error) {
	Masalahs, totalData, err := m.MasalahR.GetMasalahsRepository(page, limit, order)
	if err != nil {
		return nil, 0, err
	}

	return Masalahs, totalData, nil
}

func (m *masalahService) GetMasalahService(id string) (*models.Masalah, error) {
	Masalah, err := m.MasalahR.GetMasalahRepository(id)
	if err != nil {
		return nil, err
	}

	return Masalah, nil
}

func (m *masalahService) CreateService(Masalah models.Masalah) (*models.Masalah, error) {
	MasalahR, err := m.MasalahR.CreateRepository(Masalah)
	if err != nil {
		return nil, err
	}

	return MasalahR, nil
}
