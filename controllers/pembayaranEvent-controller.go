<<<<<<< HEAD
package controllers
=======
<<<<<<< HEAD
package controllers

import (
	h "capstone/helpers"
	"capstone/models"
	"capstone/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PembayaranEventController interface {
	CreateController(c echo.Context) error
	UploadBuktiPembayaranController(c echo.Context) error
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
		Message: "Create Pembayaran success",
		Status:  true,
	})
}

func (p *pembayaranEventController) UploadBuktiPembayaranController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	file, err := c.FormFile("bukti_pembayaran")
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Error To upload Image")
	}

	PembayaranEvent, err := p.pembayaranEventS.UploadBuktiPembayaran(file, id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    PembayaranEvent,
		Message: "Create Pembayaran success",
		Status:  true,
	})
}
=======
package controllers
>>>>>>> e9c008b3a413c76df853ebd85f12f1dbfcc1e083
>>>>>>> dce874a50a0a2727680431f9b66e19883635b44a
