package repositories

import (
	"capstone/models"

	"gorm.io/gorm"
)

type ProdukRepository interface {
	GetProduksRepository() ([]*models.Produk, error)
	GetProdukRepository(id string) (*models.Produk, error)
	CreateRepository(Produk models.Produk) (*models.Produk, error)
	UpdateRepository(id string, ProdukBody models.Produk) (*models.Produk, error)
	DeleteRepository(id string) error
}

type produkRepository struct {
	DB *gorm.DB
}

func NewProdukRepository(DB *gorm.DB) ProdukRepository {
	return &produkRepository{
		DB: DB,
	}
}

func (p *produkRepository) GetProduksRepository() ([]*models.Produk, error) {
	var Produks []*models.Produk

	if err := p.DB.Find(&Produks).Error; err != nil {
		return nil, err
	}

	return Produks, nil
}

func (p *produkRepository) GetProdukRepository(id string) (*models.Produk, error) {
	var Produk *models.Produk

	if err := p.DB.Where("ID = ?", id).Take(&Produk).Error; err != nil {
		return nil, err
	}

	return Produk, nil
}

func (p *produkRepository) CreateRepository(Produk models.Produk) (*models.Produk, error) {
	if err := p.DB.Save(&Produk).Error; err != nil {
		return nil, err
	}

	return &Produk, nil
}

func (p *produkRepository) UpdateRepository(id string, ProdukBody models.Produk) (*models.Produk, error) {
	Produk, err := p.GetProdukRepository(id)
	if err != nil {
		return nil, err
	}

	err = p.DB.Where("ID = ?", id).Updates(models.Produk{Kategori_ID: ProdukBody.Kategori_ID , Nama: ProdukBody.Nama, Deskripsi: ProdukBody.Deskripsi, Harga: ProdukBody.Harga, Stok: ProdukBody.Stok}).Error
	if err != nil {
		return nil, err
	}

	Produk.Kategori_ID = ProdukBody.Kategori_ID
	Produk.Nama = ProdukBody.Nama
	Produk.Deskripsi = ProdukBody.Deskripsi
	Produk.Harga = ProdukBody.Harga
	Produk.Stok = ProdukBody.Stok

	return Produk, nil
}

func (p *produkRepository) DeleteRepository(id string) error {
	_, err := p.GetProdukRepository(id)
	if err != nil {
		return err
	}

	if err := p.DB.Delete(&models.Produk{}, id).Error; err != nil {
		return err
	}

	return nil
}
