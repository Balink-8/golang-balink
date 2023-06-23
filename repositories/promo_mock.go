package repositories

import (
	"capstone/models"
	"fmt"

	"github.com/stretchr/testify/mock"
)

type PromoRepositoryMock interface {
	GetPromosRepository(page int, limit int, order string, search string) ([]*models.Promo, int, error)
	GetPromoRepository(id string) (*models.Promo, error)
	CreateRepository(Promo models.Promo) (*models.Promo, error)
	UpdateRepository(id string, PromoBody models.Promo) (*models.Promo, error)
	DeleteRepository(id string) error
}

type IPromoRepositoryMock struct {
	Mock mock.Mock
}

func NewPromoRepositoryMock(mock mock.Mock) PromoRepositoryMock {
	return &IPromoRepositoryMock{
		Mock: mock,
	}
}

func (a *IPromoRepositoryMock) GetPromosRepository(page int, limit int, order string, search string) ([]*models.Promo, int, error) {
	args := a.Mock.Called()
	if args.Get(0) == nil {
		return nil, 0, args.Get(1).(error)
	}

	Promos := args.Get(0).([]*models.Promo)

	return Promos, 0, nil
}

func (a *IPromoRepositoryMock) GetPromoRepository(id string) (*models.Promo, error) {
	args := a.Mock.Called(id)
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	Promo := args.Get(0).(models.Promo)

	return &Promo, nil
}

func (a *IPromoRepositoryMock) CreateRepository(PromoData models.Promo) (*models.Promo, error) {
	args := a.Mock.Called(PromoData)
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	Promo := args.Get(0).(models.Promo)

	return &Promo, nil
}

func (a *IPromoRepositoryMock) UpdateRepository(id string, PromoData models.Promo) (*models.Promo, error) {
	args := a.Mock.Called(id, PromoData)
	if args.Get(0) == nil {
		return nil,  args.Get(1).(error)
	}

	Promo := args.Get(0).(models.Promo)

	return &Promo, nil
}

func (a *IPromoRepositoryMock) DeleteRepository(id string) error {
	args := a.Mock.Called(id)
	if args.Get(0) != nil {
		return fmt.Errorf("must nil")
	}

	return nil
}
