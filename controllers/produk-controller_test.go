package controllers

import (
	"capstone/models"
	"capstone/repositories"
	"capstone/services"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	ProdukRMock = &repositories.IProdukRepositoryMock{Mock: mock.Mock{}}
	ProdukSMock = services.NewProdukService(ProdukRMock)
	ProdukCTest = NewProdukController(ProdukSMock)
)

func TestGetProduksController_Success(t *testing.T) {
	Produks := []*models.Produk{
		{
			Kategori_ID: "1",
			Nama: "Perhiasan",
			Deskripsi: "Lorem Ipsum",
			Harga: 10000,
			Stok: 10,
		},
		{
			Kategori_ID: "1",
			Nama: "Perhiasan",
			Deskripsi: "Lorem Ipsum",
			Harga: 10000,
			Stok: 10,
		},
	}

	ProdukRMock.Mock.On("GetProduksRepository").Return(Produks, nil)

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	e := echo.New()

	c := e.NewContext(req, rec)

	err := ProdukCTest.GetProduksController(c)
	assert.Nil(t, err)
}

func TestGetProduksController_Failure(t *testing.T) {
	ProdukRMock = &repositories.IProdukRepositoryMock{Mock: mock.Mock{}}
	ProdukSMock = services.NewProdukService(ProdukRMock)
	ProdukCTest = NewProdukController(ProdukSMock)
	ProdukRMock.Mock.On("GetProduksRepository").Return(nil, errors.New("get all Produks failed"))

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	e := echo.New()

	c := e.NewContext(req, rec)

	err := ProdukCTest.GetProduksController(c)
	assert.Nil(t, err)
}

func TestGetProdukController_Success(t *testing.T) {
	Produk := models.Produk{
		Model: gorm.Model{
			ID: 2,
		},
		Kategori_ID: "1",
        Nama: "Perhiasan",
        Deskripsi: "Lorem Ipsum",
		Harga: 10000,
		Stok: 10,
	}

	ProdukRMock.Mock.On("GetProdukRepository", "2").Return(Produk, nil)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/Produks/2", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	err := ProdukCTest.GetProdukController(c)
	assert.Nil(t, err)
}

func TestGetProdukController_Failure1(t *testing.T) {
	ProdukRMock.Mock.On("GetProdukRepository", "qwe").Return(nil, errors.New("get Produk failed"))

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/Produks/qwe", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("qwe")

	err := ProdukCTest.GetProdukController(c)
	assert.Nil(t, err)
}

func TestGetProdukController_Failure2(t *testing.T) {
	ProdukRMock.Mock.On("GetProdukRepository", "3").Return(nil, fmt.Errorf("Produk not found"))

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/Produks/3", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("3")

	err := ProdukCTest.GetProdukController(c)
	assert.Nil(t, err)
}

func TestCreateProdukController_Success(t *testing.T) {
	Produk := models.Produk{
		Kategori_ID: "1",
        Nama: "Perhiasan",
        Deskripsi: "Lorem Ipsum",
		Harga: 10000,
		Stok: 10,
	}

	ProdukRMock.Mock.On("CreateRepository", Produk).Return(Produk, nil)

	rec := httptest.NewRecorder()

	ProdukByte, err := json.Marshal(Produk)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(ProdukByte))

	req := httptest.NewRequest(http.MethodPost, "/Produks", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err = ProdukCTest.CreateController(c)
	assert.Nil(t, err)
}

func TestCreateProdukController_Failure1(t *testing.T) {
	Produk := models.Produk{}

	ProdukRMock.Mock.On("CreateRepository", Produk).Return(nil, errors.New("something wrong"))

	rec := httptest.NewRecorder()

	ProdukByte, err := json.Marshal(Produk)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(ProdukByte))

	req := httptest.NewRequest(http.MethodPost, "/Produks", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err = ProdukCTest.CreateController(c)
	assert.Nil(t, err)
}

func TestCreateProdukController_Failure2(t *testing.T) {
	Produk := models.Produk{}

	ProdukRMock.Mock.On("CreateRepository", Produk).Return(nil, errors.New("something wrong"))

	rec := httptest.NewRecorder()

	reqBody := strings.NewReader(string([]byte("qwe")))

	req := httptest.NewRequest(http.MethodPost, "/Produks", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err := ProdukCTest.CreateController(c)
	assert.Nil(t, err)
}

func TestUpdateProdukController_Success(t *testing.T) {
	Produk := models.Produk{
		Model: gorm.Model{
			ID: 1,
		},
		Kategori_ID: "1",
        Nama: "Perhiasan",
        Deskripsi: "Lorem Ipsum",
		Harga: 10000,
		Stok: 10,
	}

	ProdukRMock.Mock.On("UpdateRepository", "1", Produk).Return(Produk, nil)

	rec := httptest.NewRecorder()

	ProdukByte, err := json.Marshal(Produk)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(ProdukByte))

	req := httptest.NewRequest(http.MethodPut, "/Produks/1", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err = ProdukCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestUpdateProdukController_Failure1(t *testing.T) {
	Produk := models.Produk{
		Model: gorm.Model{
			ID: 1,
		},
		Kategori_ID: "1",
        Nama: "Perhiasan",
        Deskripsi: "Lorem Ipsum",
		Harga: 10000,
		Stok: 10,
	}

	ProdukRMock.Mock.On("UpdateRepository", "1", Produk).Return(Produk, nil)

	rec := httptest.NewRecorder()

	ProdukByte, err := json.Marshal(Produk)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(ProdukByte))

	req := httptest.NewRequest(http.MethodPut, "/Produks/qwe", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("qwe")

	err = ProdukCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestUpdateProdukController_Failure2(t *testing.T) {
	Produk := models.Produk{}

	ProdukRMock.Mock.On("UpdateRepository", "1", Produk).Return(Produk, nil)

	rec := httptest.NewRecorder()

	_, err := json.Marshal(Produk)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string([]byte("qwe")))

	req := httptest.NewRequest(http.MethodPut, "/Produks/qwe", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err = ProdukCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestUpdateProdukController_Failure3(t *testing.T) {
	ProdukRMock = &repositories.IProdukRepositoryMock{Mock: mock.Mock{}}
	ProdukSMock = services.NewProdukService(ProdukRMock)
	ProdukCTest = NewProdukController(ProdukSMock)
	Produk := models.Produk{
		Model: gorm.Model{
			ID: 1,
		},
		Kategori_ID: "1",
        Nama: "Perhiasan",
        Deskripsi: "Lorem Ipsum",
		Harga: 10000,
		Stok: 10,
	}

	ProdukRMock.Mock.On("UpdateRepository", "1", Produk).Return(nil, errors.New("something wrong"))

	rec := httptest.NewRecorder()

	ProdukByte, err := json.Marshal(Produk)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(ProdukByte))

	req := httptest.NewRequest(http.MethodPut, "/Produks/1", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err = ProdukCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestDeleteProdukController_Success(t *testing.T) {
	ProdukRMock.Mock.On("DeleteRepository", "2").Return(nil)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/Produks/2", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	err := ProdukCTest.DeleteController(c)

	assert.Nil(t, err)
}

func TestDeleteProdukController_Failure1(t *testing.T) {
	ProdukRMock = &repositories.IProdukRepositoryMock{Mock: mock.Mock{}}
	ProdukSMock = services.NewProdukService(ProdukRMock)
	ProdukCTest = NewProdukController(ProdukSMock)
	ProdukRMock.Mock.On("DeleteRepository", "2").Return(fmt.Errorf("Produk not found"))
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/Produks/2", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	err := ProdukCTest.DeleteController(c)

	assert.Nil(t, err)
}

func TestDeleteProdukController_Failure2(t *testing.T) {
	ProdukRMock.Mock.On("DeleteRepository", "2").Return(fmt.Errorf("Produk not found"))
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/Produks/qwe", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("qwe")

	err := ProdukCTest.DeleteController(c)

	assert.Nil(t, err)
}
