package repositories

import (
	"capstone/models"

	"gorm.io/gorm"
)

type KeranjangRepository interface {
	GetKeranjangsRepository(page int, limit int, order string) ([]*models.Keranjang, int, error)
	GetKeranjangRepository(id string) (*models.Keranjang, error)
	CreateRepository(Keranjang models.Keranjang) (*models.Keranjang, error)
	UpdateRepository(id string, KeranjangBody models.Keranjang) (*models.Keranjang, error)
	DeleteRepository(id string) error
	GetKeranjangByUserRepository(User_ID string) ([]*models.Keranjang, error)
}

type keranjangRepository struct {
	DB *gorm.DB
}

func NewKeranjangRepository(DB *gorm.DB) KeranjangRepository {
	return &keranjangRepository{
		DB: DB,
	}
}

func (k *keranjangRepository) GetKeranjangsRepository(page int, limit int, order string) ([]*models.Keranjang, int, error) {
	var Keranjangs []*models.Keranjang
	var totalData int64

	if err := k.DB.Model(&models.Keranjang{}).Count(&totalData).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	query := k.DB.Offset(offset).Limit(limit)

	switch order {
	case "asc":
		query = query.Order("ID ASC")
	case "desc":
		query = query.Order("ID DESC")
	}

	if err := query.Find(&Keranjangs).Error; err != nil {
		return nil, 0, err
	}

	return Keranjangs, int(totalData), nil
}

func (k *keranjangRepository) GetKeranjangRepository(id string) (*models.Keranjang, error) {
	var Keranjang *models.Keranjang

	if err := k.DB.Where("id = ?", id).Take(&Keranjang).Error; err != nil {
		return nil, err
	}

	return Keranjang, nil
}

func (k *keranjangRepository) CreateRepository(Keranjang models.Keranjang) (*models.Keranjang, error) {
	if err := k.DB.Save(&Keranjang).Error; err != nil {
		return nil, err
	}

	return &Keranjang, nil
}

func (k *keranjangRepository) UpdateRepository(id string, KeranjangBody models.Keranjang) (*models.Keranjang, error) {
	Keranjang, err := k.GetKeranjangRepository(id)
	if err != nil {
		return nil, err
	}

	err = k.DB.Where("ID = ?", id).Updates(models.Keranjang{User_ID: KeranjangBody.User_ID, Produk_ID: KeranjangBody.Produk_ID, Total: KeranjangBody.Total}).Error
	if err != nil {
		return nil, err
	}

	Keranjang.User_ID = KeranjangBody.User_ID
	Keranjang.Produk_ID = KeranjangBody.Produk_ID
	Keranjang.Total = KeranjangBody.Total

	return Keranjang, nil
}

func (k *keranjangRepository) DeleteRepository(id string) error {
	_, err := k.GetKeranjangRepository(id)
	if err != nil {
		return err
	}

	if err := k.DB.Delete(&models.Keranjang{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (k *keranjangRepository) GetKeranjangByUserRepository(User_ID string) ([]*models.Keranjang, error) {
	var Keranjangs []*models.Keranjang

	if err := k.DB.Where("User_ID = ?", User_ID).Find(&Keranjangs).Error; err != nil {
		return nil, err
	}
	return Keranjangs, nil
}
