package controllers

import (
	h "capstone/helpers"
	"capstone/models"
	"capstone/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PembayaranProdukController interface {
	CreateController(c echo.Context) error
}

type pembayaranProdukController struct {
	pembayaranProdukS services.PembayaranProdukService
}

func NewPembayaranProdukController(pembayaranProdukS services.PembayaranProdukService) PembayaranProdukController {
	return &pembayaranProdukController{
		pembayaranProdukS: pembayaranProdukS,
	}
}

func (p *pembayaranProdukController) CreateController(c echo.Context) error {
	var PembayaranProduk *models.PembayaranProduk

	err := c.Bind(&PembayaranProduk)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	PembayaranProduk, err = p.pembayaranProdukS.CreateService(*PembayaranProduk)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    PembayaranProduk,
		Message: "Create Produk success",
		Status:  true,
	})
}
