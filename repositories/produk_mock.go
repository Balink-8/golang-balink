package repositories

import (
	"capstone/models"
	"fmt"

	"github.com/stretchr/testify/mock"
)

type ProdukRepositoryMock interface {
	GetProduksRepository(page int, limit int, order string, search string) ([]*models.Produk, int, error)
	GetProdukRepository(id string) (*models.Produk, error)
	CreateRepository(Produk models.Produk) (*models.Produk, error)
	UpdateRepository(id string, ProdukBody models.Produk) (*models.Produk, error)
	DeleteRepository(id string) error
	GetProdukByKategoriRepository(id_Kategori string) ([]*models.Produk, error)
}

type IProdukRepositoryMock struct {
	Mock mock.Mock
}

func NewProdukRepositoryMock(mock mock.Mock) ProdukRepositoryMock {
	return &IProdukRepositoryMock{
		Mock: mock,
	}
}

func (a *IProdukRepositoryMock) GetProduksRepository(page int, limit int, order string, search string) ([]*models.Produk, int, error) {
	args := a.Mock.Called()
	if args.Get(0) == nil {
		return nil, 0, args.Get(1).(error)
	}

	Produks := args.Get(0).([]*models.Produk)

	return Produks, 0, nil
}

func (a *IProdukRepositoryMock) GetProdukRepository(id string) (*models.Produk, error) {
	args := a.Mock.Called(id)
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	Produk := args.Get(0).(models.Produk)

	return &Produk, nil
}

func (a *IProdukRepositoryMock) CreateRepository(ProdukData models.Produk) (*models.Produk, error) {
	args := a.Mock.Called(ProdukData)
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	Produk := args.Get(0).(models.Produk)

	return &Produk, nil
}

func (a *IProdukRepositoryMock) UpdateRepository(id string, ProdukData models.Produk) (*models.Produk, error) {
	args := a.Mock.Called(id, ProdukData)
	if args.Get(0) == nil {
		return nil,  args.Get(1).(error)
	}

	Produk := args.Get(0).(models.Produk)

	return &Produk, nil
}

func (a *IProdukRepositoryMock) DeleteRepository(id string) error {
	args := a.Mock.Called(id)
	if args.Get(0) != nil {
		return fmt.Errorf("must nil")
	}

	return nil
}

func (u *IProdukRepositoryMock) GetProdukByKategoriRepository(Kategori_ID string) ([]*models.Produk, error) {
	args := u.Mock.Called(Kategori_ID)
	if args.Get(0) == nil {
		return nil, args.Get(0).(error)
	}

	Produks := args.Get(0).([]*models.Produk)

	return Produks, nil
}