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
	KeranjangRMock = &repositories.IKeranjangRepositoryMock{Mock: mock.Mock{}}
	KeranjangSMock = services.NewKeranjangService(KeranjangRMock)
	KeranjangCTest = NewKeranjangController(KeranjangSMock)
)

func TestGetKeranjangsController_Success(t *testing.T) {
	Keranjangs := []*models.Keranjang{
		{
			User_ID: "1",
			Produk_ID: "1",
			Jumlah: 3,
		},
		{
			User_ID: "1",
			Produk_ID: "1",
			Jumlah: 3,
		},
	}

	KeranjangRMock.Mock.On("GetKeranjangsRepository").Return(Keranjangs, nil)

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	e := echo.New()

	c := e.NewContext(req, rec)

	err := KeranjangCTest.GetKeranjangsController(c)
	assert.Nil(t, err)
}

func TestGetKeranjangsController_Failure(t *testing.T) {
	KeranjangRMock = &repositories.IKeranjangRepositoryMock{Mock: mock.Mock{}}
	KeranjangSMock = services.NewKeranjangService(KeranjangRMock)
	KeranjangCTest = NewKeranjangController(KeranjangSMock)
	KeranjangRMock.Mock.On("GetKeranjangsRepository").Return(nil, errors.New("get all Keranjangs failed"))

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	e := echo.New()

	c := e.NewContext(req, rec)

	err := KeranjangCTest.GetKeranjangsController(c)
	assert.Nil(t, err)
}

func TestGetKeranjangController_Success(t *testing.T) {
	Keranjang := models.Keranjang{
		Model: gorm.Model{
			ID: 2,
		},
		User_ID: "1",
        Produk_ID: "1",
		Jumlah: 3,
	}

	KeranjangRMock.Mock.On("GetKeranjangRepository", "2").Return(Keranjang, nil)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/Keranjangs/2", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	err := KeranjangCTest.GetKeranjangController(c)
	assert.Nil(t, err)
}

func TestGetKeranjangController_Failure1(t *testing.T) {
	KeranjangRMock.Mock.On("GetKeranjangRepository", "qwe").Return(nil, errors.New("get Keranjang failed"))

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/Keranjangs/qwe", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("qwe")

	err := KeranjangCTest.GetKeranjangController(c)
	assert.Nil(t, err)
}

func TestGetKeranjangController_Failure2(t *testing.T) {
	KeranjangRMock.Mock.On("GetKeranjangRepository", "3").Return(nil, fmt.Errorf("Keranjang not found"))

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/Keranjangs/3", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("3")

	err := KeranjangCTest.GetKeranjangController(c)
	assert.Nil(t, err)
}

func TestCreateKeranjangController_Success(t *testing.T) {
	Keranjang := models.Keranjang{
		User_ID: "1",
        Produk_ID: "1",
		Jumlah: 3,
	}

	KeranjangRMock.Mock.On("CreateRepository", Keranjang).Return(Keranjang, nil)

	rec := httptest.NewRecorder()

	KeranjangByte, err := json.Marshal(Keranjang)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(KeranjangByte))

	req := httptest.NewRequest(http.MethodPost, "/Keranjangs", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err = KeranjangCTest.CreateController(c)
	assert.Nil(t, err)
}

func TestCreateKeranjangController_Failure1(t *testing.T) {
	Keranjang := models.Keranjang{}

	KeranjangRMock.Mock.On("CreateRepository", Keranjang).Return(nil, errors.New("something wrong"))

	rec := httptest.NewRecorder()

	KeranjangByte, err := json.Marshal(Keranjang)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(KeranjangByte))

	req := httptest.NewRequest(http.MethodPost, "/Keranjangs", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err = KeranjangCTest.CreateController(c)
	assert.Nil(t, err)
}

func TestCreateKeranjangController_Failure2(t *testing.T) {
	Keranjang := models.Keranjang{}

	KeranjangRMock.Mock.On("CreateRepository", Keranjang).Return(nil, errors.New("something wrong"))

	rec := httptest.NewRecorder()

	reqBody := strings.NewReader(string([]byte("qwe")))

	req := httptest.NewRequest(http.MethodPost, "/Keranjangs", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err := KeranjangCTest.CreateController(c)
	assert.Nil(t, err)
}

func TestUpdateKeranjangController_Success(t *testing.T) {
	Keranjang := models.Keranjang{
		Model: gorm.Model{
			ID: 1,
		},
		User_ID: "1",
        Produk_ID: "1",
		Jumlah: 3,
	}

	KeranjangRMock.Mock.On("UpdateRepository", "1", Keranjang).Return(Keranjang, nil)

	rec := httptest.NewRecorder()

	KeranjangByte, err := json.Marshal(Keranjang)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(KeranjangByte))

	req := httptest.NewRequest(http.MethodPut, "/Keranjangs/1", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err = KeranjangCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestUpdateKeranjangController_Failure1(t *testing.T) {
	Keranjang := models.Keranjang{
		Model: gorm.Model{
			ID: 1,
		},
		User_ID: "1",
        Produk_ID: "1",
		Jumlah: 3,
	}

	KeranjangRMock.Mock.On("UpdateRepository", "1", Keranjang).Return(Keranjang, nil)

	rec := httptest.NewRecorder()

	KeranjangByte, err := json.Marshal(Keranjang)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(KeranjangByte))

	req := httptest.NewRequest(http.MethodPut, "/Keranjangs/qwe", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("qwe")

	err = KeranjangCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestUpdateKeranjangController_Failure2(t *testing.T) {
	Keranjang := models.Keranjang{}

	KeranjangRMock.Mock.On("UpdateRepository", "1", Keranjang).Return(Keranjang, nil)

	rec := httptest.NewRecorder()

	_, err := json.Marshal(Keranjang)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string([]byte("qwe")))

	req := httptest.NewRequest(http.MethodPut, "/Keranjangs/qwe", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err = KeranjangCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestUpdateKeranjangController_Failure3(t *testing.T) {
	KeranjangRMock = &repositories.IKeranjangRepositoryMock{Mock: mock.Mock{}}
	KeranjangSMock = services.NewKeranjangService(KeranjangRMock)
	KeranjangCTest = NewKeranjangController(KeranjangSMock)
	Keranjang := models.Keranjang{
		Model: gorm.Model{
			ID: 1,
		},
		User_ID: "1",
        Produk_ID: "1",
		Jumlah: 3,
	}

	KeranjangRMock.Mock.On("UpdateRepository", "1", Keranjang).Return(nil, errors.New("something wrong"))

	rec := httptest.NewRecorder()

	KeranjangByte, err := json.Marshal(Keranjang)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(KeranjangByte))

	req := httptest.NewRequest(http.MethodPut, "/Keranjangs/1", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err = KeranjangCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestDeleteKeranjangController_Success(t *testing.T) {
	KeranjangRMock.Mock.On("DeleteRepository", "2").Return(nil)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/Keranjangs/2", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	err := KeranjangCTest.DeleteController(c)

	assert.Nil(t, err)
}

func TestDeleteKeranjangController_Failure1(t *testing.T) {
	KeranjangRMock = &repositories.IKeranjangRepositoryMock{Mock: mock.Mock{}}
	KeranjangSMock = services.NewKeranjangService(KeranjangRMock)
	KeranjangCTest = NewKeranjangController(KeranjangSMock)
	KeranjangRMock.Mock.On("DeleteRepository", "2").Return(fmt.Errorf("Keranjang not found"))
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/Keranjangs/2", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	err := KeranjangCTest.DeleteController(c)

	assert.Nil(t, err)
}

func TestDeleteKeranjangController_Failure2(t *testing.T) {
	KeranjangRMock.Mock.On("DeleteRepository", "2").Return(fmt.Errorf("Keranjang not found"))
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/Keranjangs/qwe", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("qwe")

	err := KeranjangCTest.DeleteController(c)

	assert.Nil(t, err)
}
