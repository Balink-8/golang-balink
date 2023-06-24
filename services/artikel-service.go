package services

import (
	"capstone/models"
	"capstone/repositories"
)

type ArtikelService interface {
	GetArtikelsService(page int, limit int, order string, search string) ([]*models.Artikel, int, error)
	GetArtikelService(id string) (*models.Artikel, error)
	CreateService(Artikel models.Artikel) (models.Artikel, error)
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

func (a *artikelService) GetArtikelsService(page int, limit int, order string, search string) ([]*models.Artikel, int, error) {
	Artikels, totalData, err := a.ArtikelR.GetArtikelsRepository(page, limit, order, search)
	if err != nil {
		return nil, 0, err
	}

	return Artikels, totalData, nil
}

func (a *artikelService) GetArtikelService(id string) (*models.Artikel, error) {
	Artikel, err := a.ArtikelR.GetArtikelRepository(id)
	if err != nil {
		return nil, err
	}

	return Artikel, nil
}

func (a *artikelService) CreateService(Produk models.Artikel) (models.Artikel, error) {
	ArtielR, err := a.ArtikelR.CreateRepository(Produk)
	if err != nil {
		return models.Artikel{}, err
	}

	return ArtielR, nil
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
