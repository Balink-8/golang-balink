package services

import (
	"capstone/models"
	"capstone/repositories"
)

type ArtikelService interface {
	GetArtikelsService() ([]*models.Artikel, error)
	GetArtikelService(id string) (*models.Artikel, error)
	CreateService(Artikel models.Artikel) (*models.Artikel, error)
	UpdateService(id string, ArtikelBody models.Artikel) (*models.Artikel, error)
	DeleteService(id string) error
}

type artikelService struct {
	ArtikelR repositories.ArtikelRepository
}

func NewArtikelService(ArtikelR repositories.ArtikelRepository) ArtikelService {
	return &artikelService{
		ArtikelR: ArtikelR,
	}
}

func (a *artikelService) GetArtikelsService() ([]*models.Artikel, error) {
	Artikels, err := a.ArtikelR.GetArtikelsRepository()
	if err != nil {
		return nil, err
	}

	return Artikels, nil
}

func (a *artikelService) GetArtikelService(id string) (*models.Artikel, error) {
	Artikel, err := a.ArtikelR.GetArtikelRepository(id)
	if err != nil {
		return nil, err
	}

	return Artikel, nil
}

func (a *artikelService) CreateService(Artikel models.Artikel) (*models.Artikel, error) {
	ArtikelR, err := a.ArtikelR.CreateRepository(Artikel)
	if err != nil {
		return nil, err
	}

	return ArtikelR, nil
}

func (a *artikelService) UpdateService(id string, ArtikelBody models.Artikel) (*models.Artikel, error) {
	Artikel, err := a.ArtikelR.UpdateRepository(id, ArtikelBody)
	if err != nil {
		return Artikel, err
	}

	return Artikel, nil
}

func (a *artikelService) DeleteService(id string) error {
	err := a.ArtikelR.DeleteRepository(id)
	if err != nil {
		return err
	}

	return nil
}
