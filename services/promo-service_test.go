package services

import (
	"capstone/models"
	"capstone/repositories"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	PromoRMock = &repositories.IPromoRepositoryMock{Mock: mock.Mock{}}
	PromoSMock = NewPromoService(PromoRMock)
)

func TestGetPromoService_Success(t *testing.T) {
	Promo := models.Promo{
        Nama: "Sembarang",
        Deskripsi: "Lorem Ipsum",
		Kode: "123abc",
		PotonganHarga: 10000,
	}

	PromoRMock.Mock.On("GetPromoRepository", "1").Return(Promo, nil)
	Promos, err := PromoSMock.GetPromoService("1")

	assert.Nil(t, err)
	assert.NotNil(t, Promos)

	assert.Equal(t, Promo.Nama, Promos.Nama)
	assert.Equal(t, Promo.Deskripsi, Promos.Deskripsi)
	assert.Equal(t, Promo.Kode, Promos.Kode)
	assert.Equal(t, Promo.PotonganHarga, Promos.PotonganHarga)
}

func TestGetPromoService_Failure(t *testing.T) {
	PromoRMock.Mock.On("GetPromoRepository", "3").Return(nil, fmt.Errorf("Promo not found"))
	Promo, err := PromoSMock.GetPromoService("3")

	assert.NotNil(t, err)
	assert.Nil(t, Promo)
}

func TestCreatePromoService_Success(t *testing.T) {
	Promo := models.Promo{
		Nama: "Sembarang",
        Deskripsi: "Lorem Ipsum",
		Kode: "123abc",
		PotonganHarga: 10000,
	}

	PromoRMock.Mock.On("CreateRepository", Promo).Return(Promo, nil)
	Promos, err := PromoSMock.CreateService(Promo)

	assert.Nil(t, err)
	assert.NotNil(t, Promos)

	assert.Equal(t, Promo.Nama, Promos.Nama)
	assert.Equal(t, Promo.Deskripsi, Promos.Deskripsi)
	assert.Equal(t, Promo.Kode, Promos.Kode)
	assert.Equal(t, Promo.PotonganHarga, Promos.PotonganHarga)
}

func TestUpdatePromoService_Success(t *testing.T) {
	Promo := models.Promo{
		Nama: "Sembarang",
        Deskripsi: "Lorem Ipsum",
		Kode: "123abc",
		PotonganHarga: 10000,
	}

	PromoRMock.Mock.On("UpdateRepository", "1", Promo).Return(Promo, nil)
	Promos, err := PromoSMock.UpdateService("1", Promo)

	assert.Nil(t, err)
	assert.NotNil(t, Promos)

	assert.Equal(t, Promo.Nama, Promos.Nama)
	assert.Equal(t, Promo.Deskripsi, Promos.Deskripsi)
	assert.Equal(t, Promo.Kode, Promos.Kode)
	assert.Equal(t, Promo.PotonganHarga, Promos.PotonganHarga)
}

func TestUpdatePromoService_Failure(t *testing.T) {
	Promo := models.Promo{
		Nama: "Sembarang",
        Deskripsi: "Lorem Ipsum",
		Kode: "123abc",
		PotonganHarga: 10000,
	}

	PromoRMock.Mock.On("UpdateRepository", "2", Promo).Return(nil, fmt.Errorf("Promo not found"))
	Promos, err := PromoSMock.UpdateService("2", Promo)

	assert.Nil(t, Promos)
	assert.NotNil(t, err)
}

func TestDeletePromoService_Success(t *testing.T) {
	PromoRMock.Mock.On("DeleteRepository", "1").Return(nil)
	err := PromoSMock.DeleteService("1")

	assert.Nil(t, err)
}

func TestDeletePromoService_Failure(t *testing.T) {
	PromoRMock.Mock.On("DeleteRepository", "2").Return(fmt.Errorf("Promo not found"))
	err := PromoSMock.DeleteService("2")

	assert.NotNil(t, err)
}
