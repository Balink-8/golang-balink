package controllers

import (
	h "capstone/helpers"
	"capstone/models"
	"capstone/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PembayaranEventController interface {
	CreateController(c echo.Context) error
}

type pembayaranEventController struct {
	pembayaranEventS services.PembayaranEventService
}

func NewPembayaranEventController(pembayaranEventS services.PembayaranEventService) PembayaranEventController {
	return &pembayaranEventController{
		pembayaranEventS: pembayaranEventS,
	}
}

func (p *pembayaranEventController) CreateController(c echo.Context) error {
	var PembayaranEvent *models.PembayaranEvent

	err := c.Bind(&PembayaranEvent)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	PembayaranEvent, err = p.pembayaranEventS.CreateService(*PembayaranEvent)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    PembayaranEvent,
		Message: "Pembayaran Event sukses",
		Status:  true,
	})
}
