package repositories

import (
	"capstone/models"
	"fmt"

	"github.com/stretchr/testify/mock"
)

type KategoriProdukRepositoryMock interface {
	GetKategoriProduksRepository(page int, limit int, order string, search string) ([]*models.KategoriProduk, int, error)
	GetKategoriProdukRepository(id string) (*models.KategoriProduk, error)
	CreateRepository(KategoriProduk models.KategoriProduk) (*models.KategoriProduk, error)
	UpdateRepository(id string, KategoriProdukBody models.KategoriProduk) (*models.KategoriProduk, error)
	DeleteRepository(id string) error
}

type IKategoriProdukRepositoryMock struct {
	Mock mock.Mock
}

func NewKategoriProdukRepositoryMock(mock mock.Mock) KategoriProdukRepositoryMock {
	return &IKategoriProdukRepositoryMock{
		Mock: mock,
	}
}

func (a *IKategoriProdukRepositoryMock) GetKategoriProduksRepository(page int, limit int, order string, search string) ([]*models.KategoriProduk, int, error) {
	args := a.Mock.Called()
	if args.Get(0) == nil {
		return nil, 0, args.Get(1).(error)
	}

	KategoriProduks := args.Get(0).([]*models.KategoriProduk)

	return KategoriProduks, 0, nil
}

func (a *IKategoriProdukRepositoryMock) GetKategoriProdukRepository(id string) (*models.KategoriProduk, error) {
	args := a.Mock.Called(id)
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	KategoriProduk := args.Get(0).(models.KategoriProduk)

	return &KategoriProduk, nil
}

func (a *IKategoriProdukRepositoryMock) CreateRepository(KategoriProdukData models.KategoriProduk) (*models.KategoriProduk, error) {
	args := a.Mock.Called(KategoriProdukData)
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	KategoriProduk := args.Get(0).(models.KategoriProduk)

	return &KategoriProduk, nil
}

func (a *IKategoriProdukRepositoryMock) UpdateRepository(id string, KategoriProdukData models.KategoriProduk) (*models.KategoriProduk, error) {
	args := a.Mock.Called(id, KategoriProdukData)
	if args.Get(0) == nil {
		return nil,  args.Get(1).(error)
	}

	KategoriProduk := args.Get(0).(models.KategoriProduk)

	return &KategoriProduk, nil
}

func (a *IKategoriProdukRepositoryMock) DeleteRepository(id string) error {
	args := a.Mock.Called(id)
	if args.Get(0) != nil {
		return fmt.Errorf("must nil")
	}

	return nil
}
