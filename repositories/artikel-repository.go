package repositories

import (
	"capstone/models"

	"gorm.io/gorm"
)

type ArtikelRepository interface {
	GetArtikelsRepository() ([]*models.Artikel, error)
	GetArtikelRepository(id string) (*models.Artikel, error)
	CreateRepository(Artikel models.Artikel) (*models.Artikel, error)
	UpdateRepository(id string, ArtikelBody models.Artikel) (*models.Artikel, error)
	DeleteRepository(id string) error
}

type artikelRepository struct {
	DB *gorm.DB
}

func NewArtikelRepository(DB *gorm.DB) ArtikelRepository {
	return &artikelRepository{
		DB: DB,
	}
}

func (a *artikelRepository) GetArtikelsRepository() ([]*models.Artikel, error) {
	var Artikels []*models.Artikel

	if err := a.DB.Find(&Artikels).Error; err != nil {
		return nil, err
	}

	return Artikels, nil
}

func (a *artikelRepository) GetArtikelRepository(id string) (*models.Artikel, error) {
	var Artikel *models.Artikel

	if err := a.DB.Where("ID = ?", id).Take(&Artikel).Error; err != nil {
		return nil, err
	}

	return Artikel, nil
}

func (a *artikelRepository) CreateRepository(Artikel models.Artikel) (*models.Artikel, error) {
	if err := a.DB.Save(&Artikel).Error; err != nil {
		return nil, err
	}

	return &Artikel, nil
}

func (a *artikelRepository) UpdateRepository(id string, ArtikelBody models.Artikel) (*models.Artikel, error) {
	Artikel, err := a.GetArtikelRepository(id)
	if err != nil {
		return nil, err
	}

	err = a.DB.Where("ID = ?", id).Updates(models.Artikel{Gambar: ArtikelBody.Gambar, Judul: ArtikelBody.Judul, Isi: ArtikelBody.Isi}).Error
	if err != nil {
		return nil, err
	}

	Artikel.Gambar = ArtikelBody.Gambar
	Artikel.Judul = ArtikelBody.Judul
	Artikel.Isi = ArtikelBody.Isi

	return Artikel, nil
}

func (a *artikelRepository) DeleteRepository(id string) error {
	_, err := a.GetArtikelRepository(id)
	if err != nil {
		return err
	}

	if err := a.DB.Delete(&models.Artikel{}, id).Error; err != nil {
		return err
	}

	return nil
}
