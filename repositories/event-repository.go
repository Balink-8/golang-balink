package repositories

import (
	"capstone/models"

	"gorm.io/gorm"
)

type EventRepository interface {
	GetEventsRepository(page int, limit int, order string, search string) ([]*models.Event, int, error)
	GetEventRepository(id string) (*models.Event, error)
	CreateRepository(Event models.Event) (*models.Event, error)
	UpdateRepository(id string, EventBody models.Event) (*models.Event, error)
	DeleteRepository(id string) error
}

type eventRepository struct {
	DB *gorm.DB
}

func NewEventRepository(DB *gorm.DB) EventRepository {
	return &eventRepository{
		DB: DB,
	}
}

func (e *eventRepository) GetEventsRepository(page int, limit int, order string, search string) ([]*models.Event, int, error) {
	var Events []*models.Event
	var totalData int64

	query := e.DB.Model(&models.Event{})

	if search != "" {
		query = query.Where("judul LIKE ?", "%"+search+"%")
	}

	if err := query.Count(&totalData).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	query = query.Offset(offset).Limit(limit)

	switch order {
	case "asc":
		query = query.Order("ID ASC")
	case "desc":
		query = query.Order("ID DESC")
	}

	if err := query.Find(&Events).Error; err != nil {
		return nil, 0, err
	}

	return Events, int(totalData), nil
}

func (e *eventRepository) GetEventRepository(id string) (*models.Event, error) {
	var Event *models.Event

	if err := e.DB.Where("ID = ?", id).Take(&Event).Error; err != nil {
		return nil, err
	}

	return Event, nil
}

func (e *eventRepository) CreateRepository(Event models.Event) (*models.Event, error) {
	if err := e.DB.Save(&Event).Error; err != nil {
		return nil, err
	}

	return &Event, nil
}

func (e *eventRepository) UpdateRepository(id string, EventBody models.Event) (*models.Event, error) {
	Event, err := e.GetEventRepository(id)
	if err != nil {
		return nil, err
	}

	err = e.DB.Where("ID = ?", id).Updates(models.Event{Artikel_ID: EventBody.Artikel_ID, Gambar: EventBody.Gambar, Nama: EventBody.Nama, Deskripsi: EventBody.Deskripsi, Stok_Tiket: EventBody.Stok_Tiket, Harga_Tiket: EventBody.Harga_Tiket, Waktu_Mulai: EventBody.Waktu_Mulai, Waktu_Selesai: EventBody.Waktu_Selesai, Tanggal_Mulai: EventBody.Tanggal_Mulai, Tanggal_Selesai: EventBody.Tanggal_Selesai, Lokasi: EventBody.Lokasi, Link_Lokasi: EventBody.Link_Lokasi}).Error
	if err != nil {
		return nil, err
	}

	Event.Artikel_ID = EventBody.Artikel_ID
	Event.Gambar = EventBody.Gambar
	Event.Nama = EventBody.Nama
	Event.Deskripsi = EventBody.Deskripsi
	Event.Stok_Tiket = EventBody.Stok_Tiket
	Event.Harga_Tiket = EventBody.Harga_Tiket
	Event.Waktu_Mulai = EventBody.Waktu_Mulai
	Event.Waktu_Selesai = EventBody.Waktu_Selesai
	Event.Tanggal_Mulai = EventBody.Tanggal_Mulai
	Event.Tanggal_Selesai = EventBody.Tanggal_Selesai
	Event.Lokasi = EventBody.Lokasi
	Event.Link_Lokasi = EventBody.Link_Lokasi

	return Event, nil
}

func (e *eventRepository) DeleteRepository(id string) error {
	_, err := e.GetEventRepository(id)
	if err != nil {
		return err
	}

	if err := e.DB.Delete(&models.Event{}, id).Error; err != nil {
		return err
	}

	return nil
}
