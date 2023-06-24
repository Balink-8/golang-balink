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
	ArtikelRMock = &repositories.IArtikelRepositoryMock{Mock: mock.Mock{}}
	ArtikelSMock = NewArtikelService(ArtikelRMock)
)

func TestGetArtikelService_Success(t *testing.T) {
	Artikel := models.Artikel{
		Judul: "Buku 1",
	}

	ArtikelRMock.Mock.On("GetArtikelRepository", "1").Return(Artikel, nil)
	Artikels, err := ArtikelSMock.GetArtikelService("1")

	assert.Nil(t, err)
	assert.NotNil(t, Artikels)

	assert.Equal(t, Artikel.Judul, Artikels.Judul)
}

func TestGetArtikelService_Failure(t *testing.T) {
	ArtikelRMock.Mock.On("GetArtikelRepository", "3").Return(nil, fmt.Errorf("Artikel not found"))
	Artikel, err := ArtikelSMock.GetArtikelService("3")

	assert.NotNil(t, err)
	assert.Nil(t, Artikel)
}

func TestCreateArtikelService_Success(t *testing.T) {
	Artikel := models.Artikel{
		Judul: "Buku 1",
	}

	ArtikelRMock.Mock.On("CreateRepository", Artikel).Return(Artikel, nil)
	Artikels, err := ArtikelSMock.CreateService(Artikel)

	assert.Nil(t, err)
	assert.NotNil(t, Artikels)

	assert.Equal(t, Artikel.Judul, Artikels.Judul)
}

func TestUpdateArtikelService_Success(t *testing.T) {
	Artikel := models.Artikel{
		Judul: "Buku 1",
	}

	ArtikelRMock.Mock.On("UpdateRepository", "1", Artikel).Return(Artikel, nil)
	Artikels, err := ArtikelSMock.UpdateService("1", Artikel)

	assert.Nil(t, err)
	assert.NotNil(t, Artikels)

	assert.Equal(t, Artikel.ID, Artikels.ID)
	assert.Equal(t, Artikel.Judul, Artikels.Judul)
}

func TestUpdateArtikelService_Failure(t *testing.T) {
	Artikel := models.Artikel{
		Judul: "Buku 1",
	}

	ArtikelRMock.Mock.On("UpdateRepository", "2", Artikel).Return(nil, fmt.Errorf("Artikel not found"))
	Artikels, err := ArtikelSMock.UpdateService("2", Artikel)

	assert.Nil(t, Artikels)
	assert.NotNil(t, err)
}

func TestDeleteArtikelService_Success(t *testing.T) {
	ArtikelRMock.Mock.On("DeleteRepository", "1").Return(nil)
	err := ArtikelSMock.DeleteService("1")

	assert.Nil(t, err)
}

func TestDeleteArtikelService_Failure(t *testing.T) {
	ArtikelRMock.Mock.On("DeleteRepository", "2").Return(fmt.Errorf("Artikel not found"))
	err := ArtikelSMock.DeleteService("2")

	assert.NotNil(t, err)
}
