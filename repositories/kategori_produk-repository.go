package repositories

import (
	"capstone/models"

	"gorm.io/gorm"
)

type KategoriProdukRepository interface {
	GetKategoriProduksRepository() ([]*models.KategoriProduk, error)
	GetKategoriProdukRepository(id string) (*models.KategoriProduk, error)
	CreateRepository(KategoriProduk models.KategoriProduk) (*models.KategoriProduk, error)
	UpdateRepository(id string, KategoriProdukBody models.KategoriProduk) (*models.KategoriProduk, error)
	DeleteRepository(id string) error
}

type kategoriProdukRepository struct {
	DB *gorm.DB
}

func NewKategoriProdukRepository(DB *gorm.DB) KategoriProdukRepository {
	return &kategoriProdukRepository{
		DB: DB,
	}
}

func (k *kategoriProdukRepository) GetKategoriProduksRepository() ([]*models.KategoriProduk, error) {
	var KategoriProduks []*models.KategoriProduk

	if err := k.DB.Find(&KategoriProduks).Error; err != nil {
		return nil, err
	}

	return KategoriProduks, nil
}

func (k *kategoriProdukRepository) GetKategoriProdukRepository(id string) (*models.KategoriProduk, error) {
	var KategoriProduk *models.KategoriProduk

	if err := k.DB.Where("ID = ?", id).Take(&KategoriProduk).Error; err != nil {
		return nil, err
	}

	return KategoriProduk, nil
}

func (k *kategoriProdukRepository) CreateRepository(KategoriProduk models.KategoriProduk) (*models.KategoriProduk, error) {
	if err := k.DB.Save(&KategoriProduk).Error; err != nil {
		return nil, err
	}

	return &KategoriProduk, nil
}

func (k *kategoriProdukRepository) UpdateRepository(id string, KategoriProdukBody models.KategoriProduk) (*models.KategoriProduk, error) {
	KategoriProduk, err := k.GetKategoriProdukRepository(id)
	if err != nil {
		return nil, err
	}

	err = k.DB.Where("ID = ?", id).Updates(models.KategoriProduk{Nama: KategoriProdukBody.Nama}).Error
	if err != nil {
		return nil, err
	}

	KategoriProduk.Nama = KategoriProdukBody.Nama

	return KategoriProduk, nil
}

func (k *kategoriProdukRepository) DeleteRepository(id string) error {
	_, err := k.GetKategoriProdukRepository(id)
	if err != nil {
		return err
	}

	if err := k.DB.Delete(&models.KategoriProduk{}, id).Error; err != nil {
		return err
	}

	return nil
}
