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
	EventRMock = &repositories.IEventRepositoryMock{Mock: mock.Mock{}}
	EventSMock = services.NewEventService(EventRMock)
	EventCTest = NewEventController(EventSMock)
)

func TestGetEventsController_Success(t *testing.T) {
	Events := []*models.Event{
		{
			Artikel_ID: "1",
			Gambar: "123",
			Nama: "Kecak 2",
			Deskripsi: "Lorem Ipsum",
			Harga_Tiket: 45000,
			Stok_Tiket: 15,
			Waktu_Mulai: "10.00",
			Waktu_Selesai: "12.00",
			Tanggal_Mulai: "12 Desember 2012",
			Tanggal_Selesai: "12 Desember 2012",
			Lokasi: "Jln. 123",
			Link_Lokasi: "123",
		},
		{
			Artikel_ID: "1",
			Gambar: "123",
			Nama: "Kecak 2",
			Deskripsi: "Lorem Ipsum",
			Harga_Tiket: 45000,
			Stok_Tiket: 15,
			Waktu_Mulai: "10.00",
			Waktu_Selesai: "12.00",
			Tanggal_Mulai: "12 Desember 2012",
			Tanggal_Selesai: "12 Desember 2012",
			Lokasi: "Jln. 123",
			Link_Lokasi: "123",
		},
	}

	EventRMock.Mock.On("GetEventsRepository").Return(Events, nil)

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	e := echo.New()

	c := e.NewContext(req, rec)

	err := EventCTest.GetEventsController(c)
	assert.Nil(t, err)
}

func TestGetEventsController_Failure(t *testing.T) {
	EventRMock = &repositories.IEventRepositoryMock{Mock: mock.Mock{}}
	EventSMock = services.NewEventService(EventRMock)
	EventCTest = NewEventController(EventSMock)
	EventRMock.Mock.On("GetEventsRepository").Return(nil, errors.New("get all Events failed"))

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	e := echo.New()

	c := e.NewContext(req, rec)

	err := EventCTest.GetEventsController(c)
	assert.Nil(t, err)
}

func TestGetEventController_Success(t *testing.T) {
	Event := models.Event{
		Model: gorm.Model{
			ID: 2,
		},
		Artikel_ID: "1",
        Gambar: "123",
        Nama: "Kecak 2",
        Deskripsi: "Lorem Ipsum",
        Harga_Tiket: 45000,
        Stok_Tiket: 15,
        Waktu_Mulai: "10.00",
        Waktu_Selesai: "12.00",
        Tanggal_Mulai: "12 Desember 2012",
        Tanggal_Selesai: "12 Desember 2012",
        Lokasi: "Jln. 123",
        Link_Lokasi: "123",
	}

	EventRMock.Mock.On("GetEventRepository", "2").Return(Event, nil)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/Events/2", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	err := EventCTest.GetEventController(c)
	assert.Nil(t, err)
}

func TestGetEventController_Failure1(t *testing.T) {
	EventRMock.Mock.On("GetEventRepository", "qwe").Return(nil, errors.New("get Event failed"))

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/Events/qwe", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("qwe")

	err := EventCTest.GetEventController(c)
	assert.Nil(t, err)
}

func TestGetEventController_Failure2(t *testing.T) {
	EventRMock.Mock.On("GetEventRepository", "3").Return(nil, fmt.Errorf("Event not found"))

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/Events/3", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("3")

	err := EventCTest.GetEventController(c)
	assert.Nil(t, err)
}

func TestCreateEventController_Success(t *testing.T) {
	Event := models.Event{
		Artikel_ID: "1",
        Gambar: "123",
        Nama: "Kecak 2",
        Deskripsi: "Lorem Ipsum",
        Harga_Tiket: 45000,
        Stok_Tiket: 15,
        Waktu_Mulai: "10.00",
        Waktu_Selesai: "12.00",
        Tanggal_Mulai: "12 Desember 2012",
        Tanggal_Selesai: "12 Desember 2012",
        Lokasi: "Jln. 123",
        Link_Lokasi: "123",
	}

	EventRMock.Mock.On("CreateRepository", Event).Return(Event, nil)

	rec := httptest.NewRecorder()

	EventByte, err := json.Marshal(Event)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(EventByte))

	req := httptest.NewRequest(http.MethodPost, "/Events", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err = EventCTest.CreateController(c)
	assert.Nil(t, err)
}

func TestCreateEventController_Failure1(t *testing.T) {
	Event := models.Event{}

	EventRMock.Mock.On("CreateRepository", Event).Return(nil, errors.New("something wrong"))

	rec := httptest.NewRecorder()

	EventByte, err := json.Marshal(Event)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(EventByte))

	req := httptest.NewRequest(http.MethodPost, "/Events", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err = EventCTest.CreateController(c)
	assert.Nil(t, err)
}

func TestCreateEventController_Failure2(t *testing.T) {
	Event := models.Event{}

	EventRMock.Mock.On("CreateRepository", Event).Return(nil, errors.New("something wrong"))

	rec := httptest.NewRecorder()

	reqBody := strings.NewReader(string([]byte("qwe")))

	req := httptest.NewRequest(http.MethodPost, "/Events", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err := EventCTest.CreateController(c)
	assert.Nil(t, err)
}

func TestUpdateEventController_Success(t *testing.T) {
	Event := models.Event{
		Model: gorm.Model{
			ID: 1,
		},
		Artikel_ID: "1",
        Gambar: "123",
        Nama: "Kecak 2",
        Deskripsi: "Lorem Ipsum",
        Harga_Tiket: 45000,
        Stok_Tiket: 15,
        Waktu_Mulai: "10.00",
        Waktu_Selesai: "12.00",
        Tanggal_Mulai: "12 Desember 2012",
        Tanggal_Selesai: "12 Desember 2012",
        Lokasi: "Jln. 123",
        Link_Lokasi: "123",
	}

	EventRMock.Mock.On("UpdateRepository", "1", Event).Return(Event, nil)

	rec := httptest.NewRecorder()

	EventByte, err := json.Marshal(Event)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(EventByte))

	req := httptest.NewRequest(http.MethodPut, "/Events/1", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err = EventCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestUpdateEventController_Failure1(t *testing.T) {
	Event := models.Event{
		Model: gorm.Model{
			ID: 1,
		},
		Artikel_ID: "1",
        Gambar: "123",
        Nama: "Kecak 2",
        Deskripsi: "Lorem Ipsum",
        Harga_Tiket: 45000,
        Stok_Tiket: 15,
        Waktu_Mulai: "10.00",
        Waktu_Selesai: "12.00",
        Tanggal_Mulai: "12 Desember 2012",
        Tanggal_Selesai: "12 Desember 2012",
        Lokasi: "Jln. 123",
        Link_Lokasi: "123",
	}

	EventRMock.Mock.On("UpdateRepository", "1", Event).Return(Event, nil)

	rec := httptest.NewRecorder()

	EventByte, err := json.Marshal(Event)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(EventByte))

	req := httptest.NewRequest(http.MethodPut, "/Events/qwe", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("qwe")

	err = EventCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestUpdateEventController_Failure2(t *testing.T) {
	Event := models.Event{}

	EventRMock.Mock.On("UpdateRepository", "1", Event).Return(Event, nil)

	rec := httptest.NewRecorder()

	_, err := json.Marshal(Event)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string([]byte("qwe")))

	req := httptest.NewRequest(http.MethodPut, "/Events/qwe", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err = EventCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestUpdateEventController_Failure3(t *testing.T) {
	EventRMock = &repositories.IEventRepositoryMock{Mock: mock.Mock{}}
	EventSMock = services.NewEventService(EventRMock)
	EventCTest = NewEventController(EventSMock)
	Event := models.Event{
		Model: gorm.Model{
			ID: 1,
		},
		Artikel_ID: "1",
        Gambar: "123",
        Nama: "Kecak 2",
        Deskripsi: "Lorem Ipsum",
        Harga_Tiket: 45000,
        Stok_Tiket: 15,
        Waktu_Mulai: "10.00",
        Waktu_Selesai: "12.00",
        Tanggal_Mulai: "12 Desember 2012",
        Tanggal_Selesai: "12 Desember 2012",
        Lokasi: "Jln. 123",
        Link_Lokasi: "123",
	}

	EventRMock.Mock.On("UpdateRepository", "1", Event).Return(nil, errors.New("something wrong"))

	rec := httptest.NewRecorder()

	EventByte, err := json.Marshal(Event)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(EventByte))

	req := httptest.NewRequest(http.MethodPut, "/Events/1", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err = EventCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestDeleteEventController_Success(t *testing.T) {
	EventRMock.Mock.On("DeleteRepository", "2").Return(nil)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/Events/2", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	err := EventCTest.DeleteController(c)

	assert.Nil(t, err)
}

func TestDeleteEventController_Failure1(t *testing.T) {
	EventRMock = &repositories.IEventRepositoryMock{Mock: mock.Mock{}}
	EventSMock = services.NewEventService(EventRMock)
	EventCTest = NewEventController(EventSMock)
	EventRMock.Mock.On("DeleteRepository", "2").Return(fmt.Errorf("Event not found"))
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/Events/2", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	err := EventCTest.DeleteController(c)

	assert.Nil(t, err)
}

func TestDeleteEventController_Failure2(t *testing.T) {
	EventRMock.Mock.On("DeleteRepository", "2").Return(fmt.Errorf("Event not found"))
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/Events/qwe", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("qwe")

	err := EventCTest.DeleteController(c)

	assert.Nil(t, err)
}
