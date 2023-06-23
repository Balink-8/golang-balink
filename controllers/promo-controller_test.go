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
	PromoRMock = &repositories.IPromoRepositoryMock{Mock: mock.Mock{}}
	PromoSMock = services.NewPromoService(PromoRMock)
	PromoCTest = NewPromoController(PromoSMock)
)

func TestGetPromosController_Success(t *testing.T) {
	Promos := []*models.Promo{
		{
			Nama: "Sembarang",
			Deskripsi: "Lorem Ipsum",
			Kode: "123abc",
			PotonganHarga: 10000,
		},
		{
			Nama: "Sembarang",
			Deskripsi: "Lorem Ipsum",
			Kode: "123abc",
			PotonganHarga: 10000,
		},
	}

	PromoRMock.Mock.On("GetPromosRepository").Return(Promos, nil)

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	e := echo.New()

	c := e.NewContext(req, rec)

	err := PromoCTest.GetPromosController(c)
	assert.Nil(t, err)
}

func TestGetPromosController_Failure(t *testing.T) {
	PromoRMock = &repositories.IPromoRepositoryMock{Mock: mock.Mock{}}
	PromoSMock = services.NewPromoService(PromoRMock)
	PromoCTest = NewPromoController(PromoSMock)
	PromoRMock.Mock.On("GetPromosRepository").Return(nil, errors.New("get all Promos failed"))

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	e := echo.New()

	c := e.NewContext(req, rec)

	err := PromoCTest.GetPromosController(c)
	assert.Nil(t, err)
}

func TestGetPromoController_Success(t *testing.T) {
	Promo := models.Promo{
		Model: gorm.Model{
			ID: 2,
		},
		Nama: "Sembarang",
        Deskripsi: "Lorem Ipsum",
		Kode: "123abc",
		PotonganHarga: 10000,
	}

	PromoRMock.Mock.On("GetPromoRepository", "2").Return(Promo, nil)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/Promos/2", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	err := PromoCTest.GetPromoController(c)
	assert.Nil(t, err)
}

func TestGetPromoController_Failure1(t *testing.T) {
	PromoRMock.Mock.On("GetPromoRepository", "qwe").Return(nil, errors.New("get Promo failed"))

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/Promos/qwe", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("qwe")

	err := PromoCTest.GetPromoController(c)
	assert.Nil(t, err)
}

func TestGetPromoController_Failure2(t *testing.T) {
	PromoRMock.Mock.On("GetPromoRepository", "3").Return(nil, fmt.Errorf("Promo not found"))

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/Promos/3", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("3")

	err := PromoCTest.GetPromoController(c)
	assert.Nil(t, err)
}

func TestCreatePromoController_Success(t *testing.T) {
	Promo := models.Promo{
		Nama: "Sembarang",
        Deskripsi: "Lorem Ipsum",
		Kode: "123abc",
		PotonganHarga: 10000,
	}

	PromoRMock.Mock.On("CreateRepository", Promo).Return(Promo, nil)

	rec := httptest.NewRecorder()

	PromoByte, err := json.Marshal(Promo)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(PromoByte))

	req := httptest.NewRequest(http.MethodPost, "/Promos", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err = PromoCTest.CreateController(c)
	assert.Nil(t, err)
}

func TestCreatePromoController_Failure1(t *testing.T) {
	Promo := models.Promo{}

	PromoRMock.Mock.On("CreateRepository", Promo).Return(nil, errors.New("something wrong"))

	rec := httptest.NewRecorder()

	PromoByte, err := json.Marshal(Promo)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(PromoByte))

	req := httptest.NewRequest(http.MethodPost, "/Promos", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err = PromoCTest.CreateController(c)
	assert.Nil(t, err)
}

func TestCreatePromoController_Failure2(t *testing.T) {
	Promo := models.Promo{}

	PromoRMock.Mock.On("CreateRepository", Promo).Return(nil, errors.New("something wrong"))

	rec := httptest.NewRecorder()

	reqBody := strings.NewReader(string([]byte("qwe")))

	req := httptest.NewRequest(http.MethodPost, "/Promos", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err := PromoCTest.CreateController(c)
	assert.Nil(t, err)
}

func TestUpdatePromoController_Success(t *testing.T) {
	Promo := models.Promo{
		Model: gorm.Model{
			ID: 1,
		},
		Nama: "Sembarang",
        Deskripsi: "Lorem Ipsum",
		Kode: "123abc",
		PotonganHarga: 10000,
	}

	PromoRMock.Mock.On("UpdateRepository", "1", Promo).Return(Promo, nil)

	rec := httptest.NewRecorder()

	PromoByte, err := json.Marshal(Promo)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(PromoByte))

	req := httptest.NewRequest(http.MethodPut, "/Promos/1", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err = PromoCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestUpdatePromoController_Failure1(t *testing.T) {
	Promo := models.Promo{
		Model: gorm.Model{
			ID: 1,
		},
		Nama: "Sembarang",
        Deskripsi: "Lorem Ipsum",
		Kode: "123abc",
		PotonganHarga: 10000,
	}

	PromoRMock.Mock.On("UpdateRepository", "1", Promo).Return(Promo, nil)

	rec := httptest.NewRecorder()

	PromoByte, err := json.Marshal(Promo)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(PromoByte))

	req := httptest.NewRequest(http.MethodPut, "/Promos/qwe", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("qwe")

	err = PromoCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestUpdatePromoController_Failure2(t *testing.T) {
	Promo := models.Promo{}

	PromoRMock.Mock.On("UpdateRepository", "1", Promo).Return(Promo, nil)

	rec := httptest.NewRecorder()

	_, err := json.Marshal(Promo)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string([]byte("qwe")))

	req := httptest.NewRequest(http.MethodPut, "/Promos/qwe", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err = PromoCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestUpdatePromoController_Failure3(t *testing.T) {
	PromoRMock = &repositories.IPromoRepositoryMock{Mock: mock.Mock{}}
	PromoSMock = services.NewPromoService(PromoRMock)
	PromoCTest = NewPromoController(PromoSMock)
	Promo := models.Promo{
		Model: gorm.Model{
			ID: 1,
		},
		Nama: "Sembarang",
        Deskripsi: "Lorem Ipsum",
		Kode: "123abc",
		PotonganHarga: 10000,
	}

	PromoRMock.Mock.On("UpdateRepository", "1", Promo).Return(nil, errors.New("something wrong"))

	rec := httptest.NewRecorder()

	PromoByte, err := json.Marshal(Promo)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(PromoByte))

	req := httptest.NewRequest(http.MethodPut, "/Promos/1", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err = PromoCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestDeletePromoController_Success(t *testing.T) {
	PromoRMock.Mock.On("DeleteRepository", "2").Return(nil)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/Promos/2", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	err := PromoCTest.DeleteController(c)

	assert.Nil(t, err)
}

func TestDeletePromoController_Failure1(t *testing.T) {
	PromoRMock = &repositories.IPromoRepositoryMock{Mock: mock.Mock{}}
	PromoSMock = services.NewPromoService(PromoRMock)
	PromoCTest = NewPromoController(PromoSMock)
	PromoRMock.Mock.On("DeleteRepository", "2").Return(fmt.Errorf("Promo not found"))
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/Promos/2", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	err := PromoCTest.DeleteController(c)

	assert.Nil(t, err)
}

func TestDeletePromoController_Failure2(t *testing.T) {
	PromoRMock.Mock.On("DeleteRepository", "2").Return(fmt.Errorf("Promo not found"))
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/Promos/qwe", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("qwe")

	err := PromoCTest.DeleteController(c)

	assert.Nil(t, err)
}
