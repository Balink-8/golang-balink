package controllers

import (
	h "capstone/helpers"
	"capstone/models"
	"capstone/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TransaksiProdukControllers interface {
	GetTransaksiProduksController(c echo.Context) error
	GetTransaksiProdukController(c echo.Context) error
	CreateTransaksiProdukController(c echo.Context) error
	DeleteTransaksiProdukController(c echo.Context) error
	GetTransaksiProdukByUserController(c echo.Context) error
}

type transaksiProdukControllers struct {
	TransaksiProdukServices services.TransaksiProdukServices
}

func NewTransaksiProdukController(TransaksiProdukServices services.TransaksiProdukServices) TransaksiProdukControllers {
	return &transaksiProdukControllers{
		TransaksiProdukServices: TransaksiProdukServices,
	}
}

func (t *transaksiProdukControllers) GetTransaksiProduksController(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	order := c.QueryParam("order")

	TransaksiProduks, totalData, err := t.TransaksiProdukServices.GetTransaksiProduksService(page, limit, order)
	if err != nil {
		return h.Response(c, 400, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	responseData := map[string]interface{}{
		"data":       TransaksiProduks,
		"total_data": totalData,
	}

	return h.Response(c, 200, h.ResponseModel{
		Data:    responseData,
		Message: "Success",
		Status:  true,
	})
}

func (t *transaksiProdukControllers) GetTransaksiProdukController(c echo.Context) error {
	id := c.Param("id")

	TransaksiProduk, err := t.TransaksiProdukServices.GetTransaksiProdukService(id)
	if err != nil {
		return h.Response(c, 400, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, 200, h.ResponseModel{
		Data:    TransaksiProduk,
		Message: "Success",
		Status:  true,
	})
}

func (t *transaksiProdukControllers) CreateTransaksiProdukController(c echo.Context) error {
	id_keranjang, err := strconv.Atoi(c.Param("id_keranjang"))
	if err != nil {
		return h.Response(c, 400, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	IDKeranjang := uint(id_keranjang)

	var TransaksiProdukBody models.TransaksiProduk
	c.Bind(&TransaksiProdukBody)

	TransaksiProduk, err := t.TransaksiProdukServices.CreateTransaksiProduk(IDKeranjang, TransaksiProdukBody)
	if err != nil {
		return h.Response(c, 400, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, 201, h.ResponseModel{
		Data:    TransaksiProduk,
		Message: "Success",
		Status:  true,
	})
}

func (t *transaksiProdukControllers) DeleteTransaksiProdukController(c echo.Context) error {
	id := c.Param("id")

	err := t.TransaksiProdukServices.DeleteTransaksiProduk(id)
	if err != nil {
		return h.Response(c, 400, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, 200, h.ResponseModel{
		Data:    nil,
		Message: "Success",
		Status:  true,
	})
}

func (t *transaksiProdukControllers) GetTransaksiProdukByUserController(c echo.Context) error {
	id := c.Param("id")

	TransaksiProduks, err := t.TransaksiProdukServices.GetTransaksiProdukByUserServices(id)
	if err != nil {
		return h.Response(c, 400, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}
	responseData := map[string]interface{}{
		"data": TransaksiProduks,
	}

	return h.Response(c, 200, h.ResponseModel{
		Data:    responseData,
		Message: "Success",
		Status:  true,
	})
}
