package repositories

import (
	"capstone/models"

	"gorm.io/gorm"
)

type TransaksiRepository interface {
	GetTransaksisRepository() ([]*models.Transaksi, error)
	GetTransaksiRepository(id string) (*models.Transaksi, error)
	CreateRepository(Transaksi models.Transaksi) (*models.Transaksi, error)
	UpdateRepository(id string, TransaksiBody models.Transaksi) (*models.Transaksi, error)
	DeleteRepository(id string) error
}

type transaksiRepository struct {
	DB *gorm.DB
}

func NewTransaksiRepository(DB *gorm.DB) TransaksiRepository {
	return &transaksiRepository{
		DB: DB,
	}
}

func (t *transaksiRepository) GetTransaksisRepository() ([]*models.Transaksi, error) {
	var Transaksis []*models.Transaksi

	if err := t.DB.Find(&Transaksis).Error; err != nil {
		return nil, err
	}

	return Transaksis, nil
}

func (t *transaksiRepository) GetTransaksiRepository(id string) (*models.Transaksi, error) {
	var Transaksi *models.Transaksi

	if err := t.DB.Where("ID = ?", id).Take(&Transaksi).Error; err != nil {
		return nil, err
	}

	return Transaksi, nil
}

func (t *transaksiRepository) CreateRepository(Transaksi models.Transaksi) (*models.Transaksi, error) {
	if err := t.DB.Save(&Transaksi).Error; err != nil {
		return nil, err
	}

	return &Transaksi, nil
}

func (t *transaksiRepository) UpdateRepository(id string, TransaksiBody models.Transaksi) (*models.Transaksi, error) {
	Transaksi, err := t.GetTransaksiRepository(id)
	if err != nil {
		return nil, err
	}

	err = t.DB.Where("ID = ?", id).Updates(models.Transaksi{Keranjang_ID: TransaksiBody.Keranjang_ID, Deskripsi: TransaksiBody.Deskripsi, Status_Transaksi: TransaksiBody.Status_Transaksi}).Error
	if err != nil {
		return nil, err
	}

	Transaksi.Keranjang_ID = TransaksiBody.Keranjang_ID
	Transaksi.Deskripsi = TransaksiBody.Deskripsi
	Transaksi.Status_Transaksi = TransaksiBody.Status_Transaksi

	return Transaksi, nil
}

func (t *transaksiRepository) DeleteRepository(id string) error {
	_, err := t.GetTransaksiRepository(id)
	if err != nil {
		return err
	}

	if err := t.DB.Delete(&models.Transaksi{}, id).Error; err != nil {
		return err
	}

	return nil
}
