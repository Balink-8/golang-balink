package repositories

import (
	"capstone/models"

	"gorm.io/gorm"
)

type TransaksiEventRepository interface {
	GetTransaksiEventsRepository(page int, limit int, order string) ([]*models.TransaksiEvent, int, error)
	GetTransaksiEventRepository(id string) (*models.TransaksiEvent, error)
	CreateTransaksiEvent(TransaksiEvent models.TransaksiEvent) (*models.TransaksiEvent, error)
	GetTransaksiEventByUserRepository(User_ID string) ([]*models.TransaksiEvent, error)
}

type transaksiEventRepository struct {
	DB *gorm.DB
}

func NewTransaksiEventRepository(db *gorm.DB) *transaksiEventRepository {
	return &transaksiEventRepository{
		DB: db,
	}
}

func (t *transaksiEventRepository) GetTransaksiEventsRepository(page int, limit int, order string) ([]*models.TransaksiEvent, int, error) {
	var transaksiEvents []*models.TransaksiEvent
	var totalData int64

	if err := t.DB.Model(&models.TransaksiEvent{}).Count(&totalData).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	query := t.DB.Offset(offset).Limit(limit)

	switch order {
	case "asc":
		query = query.Order("ID ASC")
	case "desc":
		query = query.Order("ID DESC")
	}

	if err := query.Find(&transaksiEvents).Error; err != nil {
		return nil, 0, err
	}

	return transaksiEvents, int(totalData), nil
}

func (t *transaksiEventRepository) GetTransaksiEventRepository(id string) (*models.TransaksiEvent, error) {
	var transaksiEvent *models.TransaksiEvent

	if err := t.DB.Where("id = ?", id).Take(&transaksiEvent).Error; err != nil {
		return nil, err
	}

	return transaksiEvent, nil
}

func (t *transaksiEventRepository) CreateTransaksiEvent(transaksiEvent models.TransaksiEvent) (*models.TransaksiEvent, error) {
	if err := t.DB.Create(&transaksiEvent).Error; err != nil {
		return nil, err
	}

	return &transaksiEvent, nil
}

func (t *transaksiEventRepository) GetTransaksiEventByUserRepository(User_ID string) ([]*models.TransaksiEvent, error) {
	var transaksiEvents []*models.TransaksiEvent

	if err := t.DB.Where("user_id = ?", User_ID).Find(&transaksiEvents).Error; err != nil {
		return nil, err
	}

	return transaksiEvents, nil
}
