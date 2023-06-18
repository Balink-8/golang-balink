package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	h "capstone/helpers"
	"capstone/services"
)

type DashboardController interface {
	GetDashboardController(c echo.Context) error
}

type dashboardController struct {
	DashboardS services.DashboardService
}

func NewDashboardController(DashboardS services.DashboardService) DashboardController {
	return &dashboardController{
		DashboardS: DashboardS,
	}
}

func (b *dashboardController) GetDashboardController(c echo.Context) error {
	countUserSebulan, err := b.DashboardS.GetCountUserSebulanService()
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	// countTransaksiTiketSebulan, err := b.DashboardS.GetCountTransaksiTiketSebulanService()
	// if err != nil {
	// 	return h.Response(c, http.StatusBadRequest, h.ResponseModel{
	// 		Data:    nil,
	// 		Message: err.Error(),
	// 		Status:  false,
	// 	})
	// }

	// sumTransaksiProdukSebulan, err := b.DashboardS.GetSumTransaksiProdukSebulanService()
	// if err != nil {
	// 	return h.Response(c, http.StatusBadRequest, h.ResponseModel{
	// 		Data:    nil,
	// 		Message: err.Error(),
	// 		Status:  false,
	// 	})
	// }

	countArtikelSebulan, err := b.DashboardS.GetCountArtikelSebulanService()
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	// transaksiSukses, err := b.DashboardS.GetTransaksiSuksesService()
	// if err != nil {
	// 	return h.Response(c, http.StatusBadRequest, h.ResponseModel{
	// 		Data:    nil,
	// 		Message: err.Error(),
	// 		Status:  false,
	// 	})
	// }

	// transaksiDipesan, err := b.DashboardS.GetTransaksiDipesanService()
	// if err != nil {
	// 	return h.Response(c, http.StatusBadRequest, h.ResponseModel{
	// 		Data:    nil,
	// 		Message: err.Error(),
	// 		Status:  false,
	// 	})
	// }

	// transaksiDibatalkan, err := b.DashboardS.GetTransaksiDibatalkanService()
	// if err != nil {
	// 	return h.Response(c, http.StatusBadRequest, h.ResponseModel{
	// 		Data:    nil,
	// 		Message: err.Error(),
	// 		Status:  false,
	// 	})
	// }

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    map[string]interface{}{
			"countUserSebulan":             countUserSebulan,
			// "countTransaksiTiketSebulan":   countTransaksiTiketSebulan,
			// "sumTransaksiProdukSebulan":    sumTransaksiProdukSebulan,
			"countArtikelSebulan":          countArtikelSebulan,
			// "transaksiSukses":              transaksiSukses,
			// "transaksiDipesan":             transaksiDipesan,
			// "transaksiDibatalkan":          transaksiDibatalkan,
		},
		Message: "Get Dashboard Data Success",
		Status:  true,
	})
}
