package repositories

import (
	"capstone/models"
	"fmt"

	"github.com/stretchr/testify/mock"
)

type ArtikelRepositoryMock interface {
	GetArtikelsRepository(page int, limit int, order string, search string) ([]*models.Artikel, int, error)
	GetArtikelRepository(id string) (*models.Artikel, error)
	CreateRepository(Artikel models.Artikel) (*models.Artikel, error)
	UpdateRepository(id string, ArtikelBody models.Artikel) (*models.Artikel, error)
	DeleteRepository(id string) error
}

type IArtikelRepositoryMock struct {
	Mock mock.Mock
}

func NewArtikelRepositoryMock(mock mock.Mock) ArtikelRepositoryMock {
	return &IArtikelRepositoryMock{
		Mock: mock,
	}
}

func (a *IArtikelRepositoryMock) GetArtikelsRepository(page int, limit int, order string, search string) ([]*models.Artikel, int, error) {
	args := a.Mock.Called()
	if args.Get(0) == nil {
		return nil, 0, args.Get(1).(error)
	}

	Artikels := args.Get(0).([]*models.Artikel)

	return Artikels, 0, nil
}

func (a *IArtikelRepositoryMock) GetArtikelRepository(id string) (*models.Artikel, error) {
	args := a.Mock.Called(id)
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	Artikel := args.Get(0).(models.Artikel)

	return &Artikel, nil
}

func (a *IArtikelRepositoryMock) CreateRepository(ArtikelData models.Artikel) (*models.Artikel, error) {
	args := a.Mock.Called(ArtikelData)
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	Artikel := args.Get(0).(models.Artikel)

	return &Artikel, nil
}

func (a *IArtikelRepositoryMock) UpdateRepository(id string, ArtikelData models.Artikel) (*models.Artikel, error) {
	args := a.Mock.Called(id, ArtikelData)
	if args.Get(0) == nil {
		return nil,  args.Get(1).(error)
	}

	Artikel := args.Get(0).(models.Artikel)

	return &Artikel, nil
}

func (a *IArtikelRepositoryMock) DeleteRepository(id string) error {
	args := a.Mock.Called(id)
	if args.Get(0) != nil {
		return fmt.Errorf("must nil")
	}

	return nil
}
