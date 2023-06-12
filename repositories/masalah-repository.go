package repositories

import (
	"capstone/models"

	"gorm.io/gorm"
)

type MasalahRepository interface {
	GetMasalahsRepository(page int, limit int, order string) ([]*models.Masalah, int, error)
	GetMasalahRepository(id string) (*models.Masalah, error)
	CreateRepository(Masalah models.Masalah) (*models.Masalah, error)
}

type masalahRepository struct {
	DB *gorm.DB
}

func NewMasalahRepository(DB *gorm.DB) MasalahRepository {
	return &masalahRepository{
		DB: DB,
	}
}

func (m *masalahRepository) GetMasalahsRepository(page int, limit int, order string) ([]*models.Masalah, int, error) {
	var Masalahs []*models.Masalah
	var totalData int64

	if err := m.DB.Model(&models.Masalah{}).Count(&totalData).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	query := m.DB.Offset(offset).Limit(limit)

	switch order {
	case "asc":
		query = query.Order("ID ASC")
	case "desc":
		query = query.Order("ID DESC")
	}

	if err := query.Find(&Masalahs).Error; err != nil {
		return nil, 0, err
	}

	return Masalahs, int(totalData), nil
}

func (m *masalahRepository) GetMasalahRepository(id string) (*models.Masalah, error) {
	var Masalah *models.Masalah

	if err := m.DB.Where("ID = ?", id).Take(&Masalah).Error; err != nil {
		return nil, err
	}

	return Masalah, nil
}

func (m *masalahRepository) CreateRepository(Masalah models.Masalah) (*models.Masalah, error) {
	if err := m.DB.Save(&Masalah).Error; err != nil {
		return nil, err
	}

	return &Masalah, nil
}
