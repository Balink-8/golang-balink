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
	KategoriProdukRMock = &repositories.IKategoriProdukRepositoryMock{Mock: mock.Mock{}}
	KategoriProdukSMock = services.NewKategoriProdukService(KategoriProdukRMock)
	KategoriProdukCTest = NewKategoriProdukController(KategoriProdukSMock)
)

func TestGetKategoriProduksController_Success(t *testing.T) {
	KategoriProduks := []*models.KategoriProduk{
		{
			Nama: "Perhiasan",
        	Deskripsi: "Lorem Ipsum",
		},
		{
			Nama: "Perhiasan",
        	Deskripsi: "Lorem Ipsum",
		},
	}

	KategoriProdukRMock.Mock.On("GetKategoriProduksRepository").Return(KategoriProduks, nil)

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	e := echo.New()

	c := e.NewContext(req, rec)

	err := KategoriProdukCTest.GetKategoriProduksController(c)
	assert.Nil(t, err)
}

func TestGetKategoriProduksController_Failure(t *testing.T) {
	KategoriProdukRMock = &repositories.IKategoriProdukRepositoryMock{Mock: mock.Mock{}}
	KategoriProdukSMock = services.NewKategoriProdukService(KategoriProdukRMock)
	KategoriProdukCTest = NewKategoriProdukController(KategoriProdukSMock)
	KategoriProdukRMock.Mock.On("GetKategoriProduksRepository").Return(nil, errors.New("get all KategoriProduks failed"))

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	e := echo.New()

	c := e.NewContext(req, rec)

	err := KategoriProdukCTest.GetKategoriProduksController(c)
	assert.Nil(t, err)
}

func TestGetKategoriProdukController_Success(t *testing.T) {
	KategoriProduk := models.KategoriProduk{
		Model: gorm.Model{
			ID: 2,
		},
		Nama: "Perhiasan",
        Deskripsi: "Lorem Ipsum",
	}

	KategoriProdukRMock.Mock.On("GetKategoriProdukRepository", "2").Return(KategoriProduk, nil)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/KategoriProduks/2", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	err := KategoriProdukCTest.GetKategoriProdukController(c)
	assert.Nil(t, err)
}

func TestGetKategoriProdukController_Failure1(t *testing.T) {
	KategoriProdukRMock.Mock.On("GetKategoriProdukRepository", "qwe").Return(nil, errors.New("get KategoriProduk failed"))

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/KategoriProduks/qwe", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("qwe")

	err := KategoriProdukCTest.GetKategoriProdukController(c)
	assert.Nil(t, err)
}

func TestGetKategoriProdukController_Failure2(t *testing.T) {
	KategoriProdukRMock.Mock.On("GetKategoriProdukRepository", "3").Return(nil, fmt.Errorf("KategoriProduk not found"))

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/KategoriProduks/3", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("3")

	err := KategoriProdukCTest.GetKategoriProdukController(c)
	assert.Nil(t, err)
}

func TestCreateKategoriProdukController_Success(t *testing.T) {
	KategoriProduk := models.KategoriProduk{
		Nama: "Perhiasan",
        Deskripsi: "Lorem Ipsum",
	}

	KategoriProdukRMock.Mock.On("CreateRepository", KategoriProduk).Return(KategoriProduk, nil)

	rec := httptest.NewRecorder()

	KategoriProdukByte, err := json.Marshal(KategoriProduk)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(KategoriProdukByte))

	req := httptest.NewRequest(http.MethodPost, "/KategoriProduks", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err = KategoriProdukCTest.CreateController(c)
	assert.Nil(t, err)
}

func TestCreateKategoriProdukController_Failure1(t *testing.T) {
	KategoriProduk := models.KategoriProduk{}

	KategoriProdukRMock.Mock.On("CreateRepository", KategoriProduk).Return(nil, errors.New("something wrong"))

	rec := httptest.NewRecorder()

	KategoriProdukByte, err := json.Marshal(KategoriProduk)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(KategoriProdukByte))

	req := httptest.NewRequest(http.MethodPost, "/KategoriProduks", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err = KategoriProdukCTest.CreateController(c)
	assert.Nil(t, err)
}

func TestCreateKategoriProdukController_Failure2(t *testing.T) {
	KategoriProduk := models.KategoriProduk{}

	KategoriProdukRMock.Mock.On("CreateRepository", KategoriProduk).Return(nil, errors.New("something wrong"))

	rec := httptest.NewRecorder()

	reqBody := strings.NewReader(string([]byte("qwe")))

	req := httptest.NewRequest(http.MethodPost, "/KategoriProduks", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err := KategoriProdukCTest.CreateController(c)
	assert.Nil(t, err)
}

func TestUpdateKategoriProdukController_Success(t *testing.T) {
	KategoriProduk := models.KategoriProduk{
		Model: gorm.Model{
			ID: 1,
		},
		Nama: "Perhiasan",
        Deskripsi: "Lorem Ipsum",
	}

	KategoriProdukRMock.Mock.On("UpdateRepository", "1", KategoriProduk).Return(KategoriProduk, nil)

	rec := httptest.NewRecorder()

	KategoriProdukByte, err := json.Marshal(KategoriProduk)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(KategoriProdukByte))

	req := httptest.NewRequest(http.MethodPut, "/KategoriProduks/1", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err = KategoriProdukCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestUpdateKategoriProdukController_Failure1(t *testing.T) {
	KategoriProduk := models.KategoriProduk{
		Model: gorm.Model{
			ID: 1,
		},
		Nama: "Perhiasan",
        Deskripsi: "Lorem Ipsum",
	}

	KategoriProdukRMock.Mock.On("UpdateRepository", "1", KategoriProduk).Return(KategoriProduk, nil)

	rec := httptest.NewRecorder()

	KategoriProdukByte, err := json.Marshal(KategoriProduk)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(KategoriProdukByte))

	req := httptest.NewRequest(http.MethodPut, "/KategoriProduks/qwe", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("qwe")

	err = KategoriProdukCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestUpdateKategoriProdukController_Failure2(t *testing.T) {
	KategoriProduk := models.KategoriProduk{}

	KategoriProdukRMock.Mock.On("UpdateRepository", "1", KategoriProduk).Return(KategoriProduk, nil)

	rec := httptest.NewRecorder()

	_, err := json.Marshal(KategoriProduk)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string([]byte("qwe")))

	req := httptest.NewRequest(http.MethodPut, "/KategoriProduks/qwe", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err = KategoriProdukCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestUpdateKategoriProdukController_Failure3(t *testing.T) {
	KategoriProdukRMock = &repositories.IKategoriProdukRepositoryMock{Mock: mock.Mock{}}
	KategoriProdukSMock = services.NewKategoriProdukService(KategoriProdukRMock)
	KategoriProdukCTest = NewKategoriProdukController(KategoriProdukSMock)
	KategoriProduk := models.KategoriProduk{
		Model: gorm.Model{
			ID: 1,
		},
		Nama: "Perhiasan",
        Deskripsi: "Lorem Ipsum",
	}

	KategoriProdukRMock.Mock.On("UpdateRepository", "1", KategoriProduk).Return(nil, errors.New("something wrong"))

	rec := httptest.NewRecorder()

	KategoriProdukByte, err := json.Marshal(KategoriProduk)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(KategoriProdukByte))

	req := httptest.NewRequest(http.MethodPut, "/KategoriProduks/1", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err = KategoriProdukCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestDeleteKategoriProdukController_Success(t *testing.T) {
	KategoriProdukRMock.Mock.On("DeleteRepository", "2").Return(nil)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/KategoriProduks/2", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	err := KategoriProdukCTest.DeleteController(c)

	assert.Nil(t, err)
}

func TestDeleteKategoriProdukController_Failure1(t *testing.T) {
	KategoriProdukRMock = &repositories.IKategoriProdukRepositoryMock{Mock: mock.Mock{}}
	KategoriProdukSMock = services.NewKategoriProdukService(KategoriProdukRMock)
	KategoriProdukCTest = NewKategoriProdukController(KategoriProdukSMock)
	KategoriProdukRMock.Mock.On("DeleteRepository", "2").Return(fmt.Errorf("KategoriProduk not found"))
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/KategoriProduks/2", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	err := KategoriProdukCTest.DeleteController(c)

	assert.Nil(t, err)
}

func TestDeleteKategoriProdukController_Failure2(t *testing.T) {
	KategoriProdukRMock.Mock.On("DeleteRepository", "2").Return(fmt.Errorf("KategoriProduk not found"))
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/KategoriProduks/qwe", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("qwe")

	err := KategoriProdukCTest.DeleteController(c)

	assert.Nil(t, err)
}
