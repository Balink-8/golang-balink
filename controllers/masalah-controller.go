package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	h "capstone/helpers"
	"capstone/models"
	"capstone/services"
)

type MasalahController interface {
	GetMasalahsController(c echo.Context) error
	GetMasalahController(c echo.Context) error
	CreateController(c echo.Context) error
}

type masalahController struct {
	MasalahS services.MasalahService
}

func NewMasalahController(MasalahS services.MasalahService) MasalahController {
	return &masalahController{
		MasalahS: MasalahS,
	}
}

func (m *masalahController) GetMasalahsController(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	order := c.QueryParam("order")

	Masalahs, totalData, err := m.MasalahS.GetMasalahsService(page, limit, order)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	responseData := map[string]interface{}{
		"data":       Masalahs,
		"page":       page,
		"data_shown": len(Masalahs),
		"total_data": totalData,
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    responseData,
		Message: "Get all Masalah success",
		Status:  true,
	})
}

func (m *masalahController) GetMasalahController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	var Masalah *models.Masalah

	Masalah, err = m.MasalahS.GetMasalahService(id)
	if err != nil {
		return h.Response(c, http.StatusNotFound, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Masalah,
		Message: "Get Masalah success",
		Status:  true,
	})
}

func (m *masalahController) CreateController(c echo.Context) error {
	var Masalah *models.Masalah

	err := c.Bind(&Masalah)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	Masalah, err = m.MasalahS.CreateService(*Masalah)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Masalah,
		Message: "Create Masalah success",
		Status:  true,
	})
}
