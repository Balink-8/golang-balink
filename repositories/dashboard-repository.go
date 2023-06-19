package repositories

import (
	"capstone/models"

	"gorm.io/gorm"
)

type DashboardRepository interface {
	DashboardGetAll() ([]models.Artikel, []models.Produk, []models.Promo, []models.Event, error)
}

type dashboardRepository struct {
	db *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) DashboardRepository {
	return &dashboardRepository{db}
}

func (r *dashboardRepository) DashboardGetAll() ([]models.Artikel, []models.Produk, []models.Promo, []models.Event, error) {
	artikels := []models.Artikel{}
	err := r.db.Find(&artikels).Error
	if err != nil {
		return nil, nil, nil, nil, err
	}

	produks := []models.Produk{}
	err = r.db.Find(&produks).Error
	if err != nil {
		return nil, nil, nil, nil, err
	}

	promos := []models.Promo{}
	err = r.db.Find(&promos).Error
	if err != nil {
		return nil, nil, nil, nil, err
	}

	events := []models.Event{}
	err = r.db.Find(&events).Error
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return artikels, produks, promos, events, nil
}
