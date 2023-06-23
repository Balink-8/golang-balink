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
	KeranjangRMock = &repositories.IKeranjangRepositoryMock{Mock: mock.Mock{}}
	KeranjangSMock = NewKeranjangService(KeranjangRMock)
)

func TestGetKeranjangService_Success(t *testing.T) {
	Keranjang := models.Keranjang{
        User_ID: "1",
        Produk_ID: "1",
		Jumlah: 3,
	}

	KeranjangRMock.Mock.On("GetKeranjangRepository", "1").Return(Keranjang, nil)
	Keranjangs, err := KeranjangSMock.GetKeranjangService("1")

	assert.Nil(t, err)
	assert.NotNil(t, Keranjangs)

	assert.Equal(t, Keranjang.User_ID, Keranjangs.User_ID)
	assert.Equal(t, Keranjang.Produk_ID, Keranjangs.Produk_ID)
	
}

func TestGetKeranjangService_Failure(t *testing.T) {
	KeranjangRMock.Mock.On("GetKeranjangRepository", "3").Return(nil, fmt.Errorf("Keranjang not found"))
	Keranjang, err := KeranjangSMock.GetKeranjangService("3")

	assert.NotNil(t, err)
	assert.Nil(t, Keranjang)
}

func TestCreateKeranjangService_Success(t *testing.T) {
	Keranjang := models.Keranjang{
		User_ID: "1",
        Produk_ID: "1",
		Jumlah: 3,
	}

	KeranjangRMock.Mock.On("CreateRepository", Keranjang).Return(Keranjang, nil)
	Keranjangs, err := KeranjangSMock.CreateService(Keranjang)

	assert.Nil(t, err)
	assert.NotNil(t, Keranjangs)

	assert.Equal(t, Keranjang.User_ID, Keranjangs.User_ID)
	assert.Equal(t, Keranjang.Produk_ID, Keranjangs.Produk_ID)
}

func TestUpdateKeranjangService_Success(t *testing.T) {
	Keranjang := models.Keranjang{
		User_ID: "1",
        Produk_ID: "1",
		Jumlah: 3,
	}

	KeranjangRMock.Mock.On("UpdateRepository", "1", Keranjang).Return(Keranjang, nil)
	Keranjangs, err := KeranjangSMock.UpdateService("1", Keranjang)

	assert.Nil(t, err)
	assert.NotNil(t, Keranjangs)

	assert.Equal(t, Keranjang.User_ID, Keranjangs.User_ID)
	assert.Equal(t, Keranjang.Produk_ID, Keranjangs.Produk_ID)
}

func TestUpdateKeranjangService_Failure(t *testing.T) {
	Keranjang := models.Keranjang{
		User_ID: "Perhiasan",
        Produk_ID: "Lorem Ipsum",
	}

	KeranjangRMock.Mock.On("UpdateRepository", "2", Keranjang).Return(nil, fmt.Errorf("Keranjang not found"))
	Keranjangs, err := KeranjangSMock.UpdateService("2", Keranjang)

	assert.Nil(t, Keranjangs)
	assert.NotNil(t, err)
}

func TestDeleteKeranjangService_Success(t *testing.T) {
	KeranjangRMock.Mock.On("DeleteRepository", "1").Return(nil)
	err := KeranjangSMock.DeleteService("1")

	assert.Nil(t, err)
}

func TestDeleteKeranjangService_Failure(t *testing.T) {
	KeranjangRMock.Mock.On("DeleteRepository", "2").Return(fmt.Errorf("Keranjang not found"))
	err := KeranjangSMock.DeleteService("2")

	assert.NotNil(t, err)
}
