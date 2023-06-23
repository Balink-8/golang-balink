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
	EventRMock = &repositories.IEventRepositoryMock{Mock: mock.Mock{}}
	EventSMock = NewEventService(EventRMock)
)

func TestGetEventService_Success(t *testing.T) {
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

	EventRMock.Mock.On("GetEventRepository", "1").Return(Event, nil)
	Events, err := EventSMock.GetEventService("1")

	assert.Nil(t, err)
	assert.NotNil(t, Events)

	assert.Equal(t, Event.Artikel_ID, Events.Artikel_ID)
	assert.Equal(t, Event.Gambar, Events.Gambar)
	assert.Equal(t, Event.Nama, Events.Nama)
	assert.Equal(t, Event.Deskripsi, Events.Deskripsi)
	assert.Equal(t, Event.Harga_Tiket, Events.Harga_Tiket)
	assert.Equal(t, Event.Stok_Tiket, Events.Stok_Tiket)
	assert.Equal(t, Event.Waktu_Mulai, Events.Waktu_Mulai)
	assert.Equal(t, Event.Waktu_Selesai, Events.Waktu_Selesai)
	assert.Equal(t, Event.Lokasi, Events.Lokasi)
	assert.Equal(t, Event.Link_Lokasi, Events.Link_Lokasi)
}

func TestGetEventService_Failure(t *testing.T) {
	EventRMock.Mock.On("GetEventRepository", "3").Return(nil, fmt.Errorf("Event not found"))
	Event, err := EventSMock.GetEventService("3")

	assert.NotNil(t, err)
	assert.Nil(t, Event)
}

func TestCreateEventService_Success(t *testing.T) {
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
	Events, err := EventSMock.CreateService(Event)

	assert.Nil(t, err)
	assert.NotNil(t, Events)

	assert.Equal(t, Event.Artikel_ID, Events.Artikel_ID)
	assert.Equal(t, Event.Gambar, Events.Gambar)
	assert.Equal(t, Event.Nama, Events.Nama)
	assert.Equal(t, Event.Deskripsi, Events.Deskripsi)
	assert.Equal(t, Event.Harga_Tiket, Events.Harga_Tiket)
	assert.Equal(t, Event.Stok_Tiket, Events.Stok_Tiket)
	assert.Equal(t, Event.Waktu_Mulai, Events.Waktu_Mulai)
	assert.Equal(t, Event.Waktu_Selesai, Events.Waktu_Selesai)
	assert.Equal(t, Event.Lokasi, Events.Lokasi)
	assert.Equal(t, Event.Link_Lokasi, Events.Link_Lokasi)
}

func TestUpdateEventService_Success(t *testing.T) {
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

	EventRMock.Mock.On("UpdateRepository", "1", Event).Return(Event, nil)
	Events, err := EventSMock.UpdateService("1", Event)

	assert.Nil(t, err)
	assert.NotNil(t, Events)

	assert.Equal(t, Event.Artikel_ID, Events.Artikel_ID)
	assert.Equal(t, Event.Gambar, Events.Gambar)
	assert.Equal(t, Event.Nama, Events.Nama)
	assert.Equal(t, Event.Deskripsi, Events.Deskripsi)
	assert.Equal(t, Event.Harga_Tiket, Events.Harga_Tiket)
	assert.Equal(t, Event.Stok_Tiket, Events.Stok_Tiket)
	assert.Equal(t, Event.Waktu_Mulai, Events.Waktu_Mulai)
	assert.Equal(t, Event.Waktu_Selesai, Events.Waktu_Selesai)
	assert.Equal(t, Event.Lokasi, Events.Lokasi)
	assert.Equal(t, Event.Link_Lokasi, Events.Link_Lokasi)
}

func TestUpdateEventService_Failure(t *testing.T) {
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

	EventRMock.Mock.On("UpdateRepository", "2", Event).Return(nil, fmt.Errorf("Event not found"))
	Events, err := EventSMock.UpdateService("2", Event)

	assert.Nil(t, Events)
	assert.NotNil(t, err)
}

func TestDeleteEventService_Success(t *testing.T) {
	EventRMock.Mock.On("DeleteRepository", "1").Return(nil)
	err := EventSMock.DeleteService("1")

	assert.Nil(t, err)
}

func TestDeleteEventService_Failure(t *testing.T) {
	EventRMock.Mock.On("DeleteRepository", "2").Return(fmt.Errorf("Event not found"))
	err := EventSMock.DeleteService("2")

	assert.NotNil(t, err)
}
