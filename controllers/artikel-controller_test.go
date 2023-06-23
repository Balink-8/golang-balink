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
	ArtikelRMock = &repositories.IArtikelRepositoryMock{Mock: mock.Mock{}}
	ArtikelSMock = services.NewArtikelService(ArtikelRMock)
	ArtikelCTest = NewArtikelController(ArtikelSMock)
)

func TestGetArtikelsController_Success(t *testing.T) {
	Artikels := []*models.Artikel{
		{
			Gambar: "123",
			Judul: "Kecak",
			Deskripsi: "abc",
		},
		{
			Gambar: "123",
			Judul: "Kecak",
			Deskripsi: "abc",
		},
	}

	ArtikelRMock.Mock.On("GetArtikelsRepository").Return(Artikels, nil)

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	e := echo.New()

	c := e.NewContext(req, rec)

	err := ArtikelCTest.GetArtikelsController(c)
	assert.Nil(t, err)
}

func TestGetArtikelsController_Failure(t *testing.T) {
	ArtikelRMock = &repositories.IArtikelRepositoryMock{Mock: mock.Mock{}}
	ArtikelSMock = services.NewArtikelService(ArtikelRMock)
	ArtikelCTest = NewArtikelController(ArtikelSMock)
	ArtikelRMock.Mock.On("GetArtikelsRepository").Return(nil, errors.New("get all Artikels failed"))

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	e := echo.New()

	c := e.NewContext(req, rec)

	err := ArtikelCTest.GetArtikelsController(c)
	assert.Nil(t, err)
}

func TestGetArtikelController_Success(t *testing.T) {
	Artikel := models.Artikel{
		Model: gorm.Model{
			ID: 2,
		},
		Gambar: "123",
		Judul: "Kecak",
		Deskripsi: "abc",
	}

	ArtikelRMock.Mock.On("GetArtikelRepository", "2").Return(Artikel, nil)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/Artikels/2", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	err := ArtikelCTest.GetArtikelController(c)
	assert.Nil(t, err)
}

func TestGetArtikelController_Failure1(t *testing.T) {
	ArtikelRMock.Mock.On("GetArtikelRepository", "qwe").Return(nil, errors.New("get Artikel failed"))

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/Artikels/qwe", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("qwe")

	err := ArtikelCTest.GetArtikelController(c)
	assert.Nil(t, err)
}

func TestGetArtikelController_Failure2(t *testing.T) {
	ArtikelRMock.Mock.On("GetArtikelRepository", "3").Return(nil, fmt.Errorf("Artikel not found"))

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/Artikels/3", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("3")

	err := ArtikelCTest.GetArtikelController(c)
	assert.Nil(t, err)
}

func TestCreateArtikelController_Success(t *testing.T) {
	Artikel := models.Artikel{
		Gambar: "123",
		Judul: "Kecak",
		Deskripsi: "abc",
	}

	ArtikelRMock.Mock.On("CreateRepository", Artikel).Return(Artikel, nil)

	rec := httptest.NewRecorder()

	ArtikelByte, err := json.Marshal(Artikel)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(ArtikelByte))

	req := httptest.NewRequest(http.MethodPost, "/Artikels", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err = ArtikelCTest.CreateController(c)
	assert.Nil(t, err)
}

func TestCreateArtikelController_Failure1(t *testing.T) {
	Artikel := models.Artikel{}

	ArtikelRMock.Mock.On("CreateRepository", Artikel).Return(nil, errors.New("something wrong"))

	rec := httptest.NewRecorder()

	ArtikelByte, err := json.Marshal(Artikel)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(ArtikelByte))

	req := httptest.NewRequest(http.MethodPost, "/Artikels", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err = ArtikelCTest.CreateController(c)
	assert.Nil(t, err)
}

func TestCreateArtikelController_Failure2(t *testing.T) {
	Artikel := models.Artikel{}

	ArtikelRMock.Mock.On("CreateRepository", Artikel).Return(nil, errors.New("something wrong"))

	rec := httptest.NewRecorder()

	reqBody := strings.NewReader(string([]byte("qwe")))

	req := httptest.NewRequest(http.MethodPost, "/Artikels", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err := ArtikelCTest.CreateController(c)
	assert.Nil(t, err)
}

func TestUpdateArtikelController_Success(t *testing.T) {
	Artikel := models.Artikel{
		Model: gorm.Model{
			ID: 1,
		},
		Gambar: "123",
		Judul: "Kecak",
		Deskripsi: "abc",
	}

	ArtikelRMock.Mock.On("UpdateRepository", "1", Artikel).Return(Artikel, nil)

	rec := httptest.NewRecorder()

	ArtikelByte, err := json.Marshal(Artikel)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(ArtikelByte))

	req := httptest.NewRequest(http.MethodPut, "/Artikels/1", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err = ArtikelCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestUpdateArtikelController_Failure1(t *testing.T) {
	Artikel := models.Artikel{
		Model: gorm.Model{
			ID: 1,
		},
		Gambar: "123",
		Judul: "Kecak",
		Deskripsi: "abc",
	}

	ArtikelRMock.Mock.On("UpdateRepository", "1", Artikel).Return(Artikel, nil)

	rec := httptest.NewRecorder()

	ArtikelByte, err := json.Marshal(Artikel)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(ArtikelByte))

	req := httptest.NewRequest(http.MethodPut, "/Artikels/qwe", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("qwe")

	err = ArtikelCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestUpdateArtikelController_Failure2(t *testing.T) {
	Artikel := models.Artikel{}

	ArtikelRMock.Mock.On("UpdateRepository", "1", Artikel).Return(Artikel, nil)

	rec := httptest.NewRecorder()

	_, err := json.Marshal(Artikel)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string([]byte("qwe")))

	req := httptest.NewRequest(http.MethodPut, "/Artikels/qwe", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err = ArtikelCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestUpdateArtikelController_Failure3(t *testing.T) {
	ArtikelRMock = &repositories.IArtikelRepositoryMock{Mock: mock.Mock{}}
	ArtikelSMock = services.NewArtikelService(ArtikelRMock)
	ArtikelCTest = NewArtikelController(ArtikelSMock)
	Artikel := models.Artikel{
		Model: gorm.Model{
			ID: 1,
		},
		Gambar: "123",
		Judul: "Kecak",
		Deskripsi: "abc",
	}

	ArtikelRMock.Mock.On("UpdateRepository", "1", Artikel).Return(nil, errors.New("something wrong"))

	rec := httptest.NewRecorder()

	ArtikelByte, err := json.Marshal(Artikel)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(ArtikelByte))

	req := httptest.NewRequest(http.MethodPut, "/Artikels/1", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err = ArtikelCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestDeleteArtikelController_Success(t *testing.T) {
	ArtikelRMock.Mock.On("DeleteRepository", "2").Return(nil)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/Artikels/2", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	err := ArtikelCTest.DeleteController(c)

	assert.Nil(t, err)
}

func TestDeleteArtikelController_Failure1(t *testing.T) {
	ArtikelRMock = &repositories.IArtikelRepositoryMock{Mock: mock.Mock{}}
	ArtikelSMock = services.NewArtikelService(ArtikelRMock)
	ArtikelCTest = NewArtikelController(ArtikelSMock)
	ArtikelRMock.Mock.On("DeleteRepository", "2").Return(fmt.Errorf("Artikel not found"))
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/Artikels/2", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	err := ArtikelCTest.DeleteController(c)

	assert.Nil(t, err)
}

func TestDeleteArtikelController_Failure2(t *testing.T) {
	ArtikelRMock.Mock.On("DeleteRepository", "2").Return(fmt.Errorf("Artikel not found"))
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/Artikels/qwe", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("qwe")

	err := ArtikelCTest.DeleteController(c)

	assert.Nil(t, err)
}
