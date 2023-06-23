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
	KategoriProdukRMock = &repositories.IKategoriProdukRepositoryMock{Mock: mock.Mock{}}
	KategoriProdukSMock = NewKategoriProdukService(KategoriProdukRMock)
)

func TestGetKategoriProdukService_Success(t *testing.T) {
	KategoriProduk := models.KategoriProduk{
        Nama: "Perhiasan",
        Deskripsi: "Lorem Ipsum",
	}

	KategoriProdukRMock.Mock.On("GetKategoriProdukRepository", "1").Return(KategoriProduk, nil)
	KategoriProduks, err := KategoriProdukSMock.GetKategoriProdukService("1")

	assert.Nil(t, err)
	assert.NotNil(t, KategoriProduks)

	assert.Equal(t, KategoriProduk.Nama, KategoriProduks.Nama)
	assert.Equal(t, KategoriProduk.Deskripsi, KategoriProduks.Deskripsi)
}

func TestGetKategoriProdukService_Failure(t *testing.T) {
	KategoriProdukRMock.Mock.On("GetKategoriProdukRepository", "3").Return(nil, fmt.Errorf("KategoriProduk not found"))
	KategoriProduk, err := KategoriProdukSMock.GetKategoriProdukService("3")

	assert.NotNil(t, err)
	assert.Nil(t, KategoriProduk)
}

func TestCreateKategoriProdukService_Success(t *testing.T) {
	KategoriProduk := models.KategoriProduk{
		Nama: "Perhiasan",
        Deskripsi: "Lorem Ipsum",
	}

	KategoriProdukRMock.Mock.On("CreateRepository", KategoriProduk).Return(KategoriProduk, nil)
	KategoriProduks, err := KategoriProdukSMock.CreateService(KategoriProduk)

	assert.Nil(t, err)
	assert.NotNil(t, KategoriProduks)

	assert.Equal(t, KategoriProduk.Nama, KategoriProduks.Nama)
	assert.Equal(t, KategoriProduk.Deskripsi, KategoriProduks.Deskripsi)
}

func TestUpdateKategoriProdukService_Success(t *testing.T) {
	KategoriProduk := models.KategoriProduk{
		Nama: "Perhiasan",
        Deskripsi: "Lorem Ipsum",
	}

	KategoriProdukRMock.Mock.On("UpdateRepository", "1", KategoriProduk).Return(KategoriProduk, nil)
	KategoriProduks, err := KategoriProdukSMock.UpdateService("1", KategoriProduk)

	assert.Nil(t, err)
	assert.NotNil(t, KategoriProduks)

	assert.Equal(t, KategoriProduk.Nama, KategoriProduks.Nama)
	assert.Equal(t, KategoriProduk.Deskripsi, KategoriProduks.Deskripsi)
}

func TestUpdateKategoriProdukService_Failure(t *testing.T) {
	KategoriProduk := models.KategoriProduk{
		Nama: "Perhiasan",
        Deskripsi: "Lorem Ipsum",
	}

	KategoriProdukRMock.Mock.On("UpdateRepository", "2", KategoriProduk).Return(nil, fmt.Errorf("KategoriProduk not found"))
	KategoriProduks, err := KategoriProdukSMock.UpdateService("2", KategoriProduk)

	assert.Nil(t, KategoriProduks)
	assert.NotNil(t, err)
}

func TestDeleteKategoriProdukService_Success(t *testing.T) {
	KategoriProdukRMock.Mock.On("DeleteRepository", "1").Return(nil)
	err := KategoriProdukSMock.DeleteService("1")

	assert.Nil(t, err)
}

func TestDeleteKategoriProdukService_Failure(t *testing.T) {
	KategoriProdukRMock.Mock.On("DeleteRepository", "2").Return(fmt.Errorf("KategoriProduk not found"))
	err := KategoriProdukSMock.DeleteService("2")

	assert.NotNil(t, err)
}
