package repositories

import (
	"capstone/models"

	"gorm.io/gorm"
)

type DashboarMobiledRepository interface {
	DashboardMobileGetAll() ([]models.KategoriProduk, []models.Promo, []models.Event, error)
}

type dashboardMobileRepository struct {
	db *gorm.DB
}

func NewDashboardMobileRepository(db *gorm.DB) *dashboardMobileRepository {
	return &dashboardMobileRepository{db}
}

func (r *dashboardMobileRepository) DashboardMobileGetAll() ([]models.KategoriProduk, []models.Promo, []models.Event, error) {
	produks := []models.KategoriProduk{}
	err := r.db.Find(&produks).Error
	if err != nil {
		return nil, nil, nil, err
	}

	promos := []models.Promo{}
	err = r.db.Find(&promos).Error
	if err != nil {
		return nil, nil, nil, err
	}

	events := []models.Event{}
	err = r.db.Find(&events).Error
	if err != nil {
		return nil, nil, nil, err
	}

	return produks, promos, events, nil
}
