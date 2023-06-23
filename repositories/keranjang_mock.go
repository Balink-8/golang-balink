package repositories

import (
	"capstone/models"
	"fmt"

	"github.com/stretchr/testify/mock"
)

type KeranjangRepositoryMock interface {
	GetKeranjangsRepository(page int, limit int, order string, search string) ([]*models.Keranjang, int, error)
	GetKeranjangRepository(id string) (*models.Keranjang, error)
	CreateRepository(Keranjang models.Keranjang) (*models.Keranjang, error)
	UpdateRepository(id string, KeranjangBody models.Keranjang) (*models.Keranjang, error)
	DeleteRepository(id string) error
	GetKeranjangByUserRepository(User_ID string) ([]*models.Keranjang, error)
}

type IKeranjangRepositoryMock struct {
	Mock mock.Mock
}

func NewKeranjangRepositoryMock(mock mock.Mock) KeranjangRepositoryMock {
	return &IKeranjangRepositoryMock{
		Mock: mock,
	}
}

func (a *IKeranjangRepositoryMock) GetKeranjangsRepository(page int, limit int, order string, search string) ([]*models.Keranjang, int, error) {
	args := a.Mock.Called()
	if args.Get(0) == nil {
		return nil, 0, args.Get(1).(error)
	}

	Keranjangs := args.Get(0).([]*models.Keranjang)

	return Keranjangs, 0, nil
}

func (a *IKeranjangRepositoryMock) GetKeranjangRepository(id string) (*models.Keranjang, error) {
	args := a.Mock.Called(id)
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	Keranjang := args.Get(0).(models.Keranjang)

	return &Keranjang, nil
}

func (a *IKeranjangRepositoryMock) CreateRepository(KeranjangData models.Keranjang) (*models.Keranjang, error) {
	args := a.Mock.Called(KeranjangData)
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	Keranjang := args.Get(0).(models.Keranjang)

	return &Keranjang, nil
}

func (a *IKeranjangRepositoryMock) UpdateRepository(id string, KeranjangData models.Keranjang) (*models.Keranjang, error) {
	args := a.Mock.Called(id, KeranjangData)
	if args.Get(0) == nil {
		return nil,  args.Get(1).(error)
	}

	Keranjang := args.Get(0).(models.Keranjang)

	return &Keranjang, nil
}

func (a *IKeranjangRepositoryMock) DeleteRepository(id string) error {
	args := a.Mock.Called(id)
	if args.Get(0) != nil {
		return fmt.Errorf("must nil")
	}

	return nil
}

func (u *IKeranjangRepositoryMock) GetKeranjangByUserRepository(User_ID string) ([]*models.Keranjang, error) {
	args := u.Mock.Called(User_ID)
	if args.Get(0) == nil {
		return nil, args.Get(0).(error)
	}

	Keranjangs := args.Get(0).([]*models.Keranjang)

	return Keranjangs, nil
}