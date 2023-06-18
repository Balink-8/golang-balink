package repositories

import (
	"capstone/models"

	"gorm.io/gorm"
)

type TransaksiProdukRepository interface {
	GetTransaksiProduksRepository(page int, limit int, order string) ([]*models.TransaksiProduk, int, error)
	GetTransaksiProdukRepository(id string) (*models.TransaksiProduk, error)
	CreateTransaksiProduk(TransaksiProduk models.TransaksiProduk) (*models.TransaksiProduk, error)
	DeleteTransaksiProduk(id string) error
	GetTransaksiProdukByUserRepository(User_ID string) ([]*models.TransaksiProduk, error)
}

type transaksiProdukRepository struct {
	DB *gorm.DB
}

func NewTransaksiProdukRepository(db *gorm.DB) *transaksiProdukRepository {
	return &transaksiProdukRepository{
		DB: db,
	}
}

func (t *transaksiProdukRepository) GetTransaksiProduksRepository(page int, limit int, order string) ([]*models.TransaksiProduk, int, error) {
	var transaksiProduks []*models.TransaksiProduk
	var totalData int64

	if err := t.DB.Model(&models.TransaksiProduk{}).Count(&totalData).Error; err != nil {
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

	if err := query.Find(&transaksiProduks).Error; err != nil {
		return nil, 0, err
	}

	return transaksiProduks, int(totalData), nil
}

func (t *transaksiProdukRepository) GetTransaksiProdukRepository(id string) (*models.TransaksiProduk, error) {
	var transaksiProduk *models.TransaksiProduk

	if err := t.DB.Where("id = ?", id).Take(&transaksiProduk).Error; err != nil {
		return nil, err
	}

	return transaksiProduk, nil
}

func (t *transaksiProdukRepository) CreateTransaksiProduk(transaksiProduk models.TransaksiProduk) (*models.TransaksiProduk, error) {
	if err := t.DB.Create(&transaksiProduk).Error; err != nil {
		return nil, err
	}

	return &transaksiProduk, nil
}

func (t *transaksiProdukRepository) DeleteTransaksiProduk(id string) error {
	var transaksiProduk *models.TransaksiProduk

	if err := t.DB.Where("id = ?", id).Take(&transaksiProduk).Error; err != nil {
		return err
	}

	if err := t.DB.Delete(&transaksiProduk).Error; err != nil {
		return err
	}

	return nil
}

func (t *transaksiProdukRepository) GetTransaksiProdukByUserRepository(User_ID string) ([]*models.TransaksiProduk, error) {
	var transaksiProduks []*models.TransaksiProduk

	if err := t.DB.Where("user_id = ?", User_ID).Find(&transaksiProduks).Error; err != nil {
		return nil, err
	}

	return transaksiProduks, nil
}
