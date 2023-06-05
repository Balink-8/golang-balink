package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	h "capstone/helpers"
	"capstone/models"
	"capstone/services"
)

type TransaksiController interface {
	GetTransaksisController(c echo.Context) error
	GetTransaksiController(c echo.Context) error
	CreateController(c echo.Context) error
	UpdateController(c echo.Context) error
	DeleteController(c echo.Context) error
}

type transaksiController struct {
	TransaksiS services.TransaksiService
}

func NewTransaksiController(TransaksiS services.TransaksiService) TransaksiController {
	return &transaksiController{
		TransaksiS: TransaksiS,
	}
}

func (t *transaksiController) GetTransaksisController(c echo.Context) error {
	Transaksis, err := t.TransaksiS.GetTransaksisService()
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Transaksis,
		Message: "Get all Transaction success",
		Status:  true,
	})
}

func (t *transaksiController) GetTransaksiController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	var Transaksi *models.Transaksi

	Transaksi, err = t.TransaksiS.GetTransaksiService(id)
	if err != nil {
		return h.Response(c, http.StatusNotFound, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Transaksi,
		Message: "Get Transaction success",
		Status:  true,
	})
}

func (t *transaksiController) CreateController(c echo.Context) error {
	var Transaksi *models.Transaksi

	err := c.Bind(&Transaksi)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	Transaksi, err = t.TransaksiS.CreateService(*Transaksi)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Transaksi,
		Message: "Create Transaction success",
		Status:  true,
	})
}

func (t *transaksiController) UpdateController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	var Transaksi *models.Transaksi

	err = c.Bind(&Transaksi)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	Transaksi, err = t.TransaksiS.UpdateService(id, *Transaksi)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Transaksi,
		Message: "Update Transaction success",
		Status:  true,
	})
}

func (t *transaksiController) DeleteController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	err = t.TransaksiS.DeleteService(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    nil,
		Message: "Delete Transaction success",
		Status:  true,
	})
}
