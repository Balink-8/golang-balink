package controllers

import (
	h "capstone/helpers"
	"capstone/models"
	"capstone/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PembayaranProdukController interface {
	CreateController(c echo.Context) error
	UploadBuktiPembayaranController(c echo.Context) error
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
		Message: "Create Pembayaran success",
		Status:  true,
	})
}

func (p *pembayaranProdukController) UploadBuktiPembayaranController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	file, err := c.FormFile("bukti_pembayaran")
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Error To upload Image")
	}

	PembayaranProduk, err := p.pembayaranProdukS.UploadBuktiPembayaran(file, id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    PembayaranProduk,
		Message: "Create Pembayaran success",
		Status:  true,
	})
}
