package controllers

import (
	h "capstone/helpers"
	"capstone/models"
	"capstone/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TransaksiEventControllers interface {
	GetTransaksiEventsController(c echo.Context) error
	GetTransaksiEventController(c echo.Context) error
	CreateTransaksiEventController(c echo.Context) error
	GetTransaksiEventByUserController(c echo.Context) error
}

type transaksiEventControllers struct {
	TransaksiEventServices services.TransaksiEventServices
}

func NewTransaksiEventController(TransaksiEventServices services.TransaksiEventServices) TransaksiEventControllers {
	return &transaksiEventControllers{
		TransaksiEventServices: TransaksiEventServices,
	}
}

func (t *transaksiEventControllers) GetTransaksiEventsController(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	order := c.QueryParam("order")

	TransaksiEvents, totalData, err := t.TransaksiEventServices.GetTransaksiEventsService(page, limit, order)
	if err != nil {
		return h.Response(c, 400, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	responseData := map[string]interface{}{
		"data":       TransaksiEvents,
		"total_data": totalData,
	}

	return h.Response(c, 200, h.ResponseModel{
		Data:    responseData,
		Message: "Success",
		Status:  true,
	})
}

func (t *transaksiEventControllers) GetTransaksiEventController(c echo.Context) error {
	id := c.Param("id")

	TransaksiEvent, err := t.TransaksiEventServices.GetTransaksiEventService(id)
	if err != nil {
		return h.Response(c, 400, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, 200, h.ResponseModel{
		Data:    TransaksiEvent,
		Message: "Success",
		Status:  true,
	})
}

func (t *transaksiEventControllers) CreateTransaksiEventController(c echo.Context) error {
	id_keranjang, err := strconv.Atoi(c.Param("id_keranjang"))
	if err != nil {
		return h.Response(c, 400, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	IDKeranjang := uint(id_keranjang)

	var TransaksiEventBody models.TransaksiEvent
	c.Bind(&TransaksiEventBody)

	TransaksiEvent, err := t.TransaksiEventServices.CreateTransaksiEvent(IDKeranjang, TransaksiEventBody)
	if err != nil {
		return h.Response(c, 400, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, 201, h.ResponseModel{
		Data:    TransaksiEvent,
		Message: "Success",
		Status:  true,
	})
}

func (t *transaksiEventControllers) GetTransaksiEventByUserController(c echo.Context) error {
	id := c.Param("id")

	TransaksiEvents, err := t.TransaksiEventServices.GetTransaksiEventByUserServices(id)
	if err != nil {
		return h.Response(c, 400, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}
	responseData := map[string]interface{}{
		"data": TransaksiEvents,
	}

	return h.Response(c, 200, h.ResponseModel{
		Data:    responseData,
		Message: "Success",
		Status:  true,
	})
}
