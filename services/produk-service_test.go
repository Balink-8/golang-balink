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
	ProdukRMock = &repositories.IProdukRepositoryMock{Mock: mock.Mock{}}
	ProdukSMock = NewProdukService(ProdukRMock)
)

func TestGetProdukService_Success(t *testing.T) {
	Produk := models.Produk{
		Kategori_ID: "1",
		Nama:        "Perhiasan",
		Deskripsi:   "Lorem Ipsum",
		Harga:       10000,
		Stok:        10,
	}

	ProdukRMock.Mock.On("GetProdukRepository", "1").Return(Produk, nil)
	Produks, err := ProdukSMock.GetProdukService("1")

	assert.Nil(t, err)
	assert.NotNil(t, Produks)

	assert.Equal(t, Produk.Kategori_ID, Produks.Kategori_ID)
	assert.Equal(t, Produk.Nama, Produks.Nama)
	assert.Equal(t, Produk.Deskripsi, Produks.Deskripsi)
	assert.Equal(t, Produk.Harga, Produks.Harga)
	assert.Equal(t, Produk.Stok, Produks.Stok)
}

func TestGetProdukService_Failure(t *testing.T) {
	ProdukRMock.Mock.On("GetProdukRepository", "3").Return(nil, fmt.Errorf("Produk not found"))
	Produk, err := ProdukSMock.GetProdukService("3")

	assert.NotNil(t, err)
	assert.Nil(t, Produk)
}

func TestCreateProdukService_Success(t *testing.T) {
	Produk := models.Produk{
		Kategori_ID: "1",
		Nama:        "Perhiasan",
		Deskripsi:   "Lorem Ipsum",
		Harga:       10000,
		Stok:        10,
	}

	ProdukRMock.Mock.On("CreateRepository", Produk).Return(Produk, nil)
	Produks, err := ProdukSMock.CreateService(Produk)

	assert.Nil(t, err)
	assert.NotNil(t, Produks)

	assert.Equal(t, Produk.Kategori_ID, Produks.Kategori_ID)
	assert.Equal(t, Produk.Nama, Produks.Nama)
	assert.Equal(t, Produk.Deskripsi, Produks.Deskripsi)
	assert.Equal(t, Produk.Harga, Produks.Harga)
	assert.Equal(t, Produk.Stok, Produks.Stok)
}

func TestUpdateProdukService_Success(t *testing.T) {
	Produk := models.Produk{
		Kategori_ID: "1",
		Nama:        "Perhiasan",
		Deskripsi:   "Lorem Ipsum",
		Harga:       10000,
		Stok:        10,
	}

	ProdukRMock.Mock.On("UpdateRepository", "1", Produk).Return(Produk, nil)
	Produks, err := ProdukSMock.UpdateService("1", Produk)

	assert.Nil(t, err)
	assert.NotNil(t, Produks)

	assert.Equal(t, Produk.Kategori_ID, Produks.Kategori_ID)
	assert.Equal(t, Produk.Nama, Produks.Nama)
	assert.Equal(t, Produk.Deskripsi, Produks.Deskripsi)
	assert.Equal(t, Produk.Harga, Produks.Harga)
	assert.Equal(t, Produk.Stok, Produks.Stok)
}

func TestUpdateProdukService_Failure(t *testing.T) {
	Produk := models.Produk{
		Kategori_ID: "1",
		Nama:        "Perhiasan",
		Deskripsi:   "Lorem Ipsum",
		Harga:       10000,
		Stok:        10,
	}

	ProdukRMock.Mock.On("UpdateRepository", "2", Produk).Return(nil, fmt.Errorf("Produk not found"))
	Produks, err := ProdukSMock.UpdateService("2", Produk)

	assert.Nil(t, Produks)
	assert.NotNil(t, err)
}

func TestDeleteProdukService_Success(t *testing.T) {
	ProdukRMock.Mock.On("DeleteRepository", "1").Return(nil)
	err := ProdukSMock.DeleteService("1")

	assert.Nil(t, err)
}

func TestDeleteProdukService_Failure(t *testing.T) {
	ProdukRMock.Mock.On("DeleteRepository", "2").Return(fmt.Errorf("Produk not found"))
	err := ProdukSMock.DeleteService("2")

	assert.NotNil(t, err)
}
