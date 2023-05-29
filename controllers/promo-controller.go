package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	h "capstone/helpers"
	"capstone/models"
	"capstone/services"
)

type PromoController interface {
	GetPromosController(c echo.Context) error
	GetPromoController(c echo.Context) error
	CreateController(c echo.Context) error
	UpdateController(c echo.Context) error
	DeleteController(c echo.Context) error
}

type promoController struct {
	PromoS services.PromoService
}

func NewPromoController(PromoS services.PromoService) PromoController {
	return &promoController{
		PromoS: PromoS,
	}
}

func (p *promoController) GetPromosController(c echo.Context) error {
	Promos, err := p.PromoS.GetPromosService()
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Promos,
		Message: "Get all Promos success",
		Status:  true,
	})
}

func (p *promoController) GetPromoController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	var Promo *models.Promo

	Promo, err = p.PromoS.GetPromoService(id)
	if err != nil {
		return h.Response(c, http.StatusNotFound, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Promo,
		Message: "Get Promo success",
		Status:  true,
	})
}

func (p *promoController) CreateController(c echo.Context) error {
	var Promo *models.Promo

	err := c.Bind(&Promo)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	Promo, err = p.PromoS.CreateService(*Promo)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Promo,
		Message: "Create Promo success",
		Status:  true,
	})
}

func (p *promoController) UpdateController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	var Promo *models.Promo

	err = c.Bind(&Promo)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	Promo, err = p.PromoS.UpdateService(id, *Promo)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Promo,
		Message: "Update Promo success",
		Status:  true,
	})
}

func (p *promoController) DeleteController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	err = p.PromoS.DeleteService(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    nil,
		Message: "Delete Promo success",
		Status:  true,
	})
}
