package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	h "capstone/helpers"
	"capstone/models"
	"capstone/services"
)

type ArtikelController interface {
	GetArtikelsController(c echo.Context) error
	GetArtikelController(c echo.Context) error
	CreateController(c echo.Context) error
	UpdateController(c echo.Context) error
	DeleteController(c echo.Context) error
}

type artikelController struct {
	ArtikelS services.ArtikelService
}

func NewArtikelController(ArtikelS services.ArtikelService) ArtikelController {
	return &artikelController{
		ArtikelS: ArtikelS,
	}
}

func (a *artikelController) GetArtikelsController(c echo.Context) error {
	Artikels, err := a.ArtikelS.GetArtikelsService()
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Artikels,
		Message: "Get all Artikels success",
		Status:  true,
	})
}

func (a *artikelController) GetArtikelController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	var Artikel *models.Artikel

	Artikel, err = a.ArtikelS.GetArtikelService(id)
	if err != nil {
		return h.Response(c, http.StatusNotFound, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Artikel,
		Message: "Get Artikel success",
		Status:  true,
	})
}

func (a *artikelController) CreateController(c echo.Context) error {
	var Artikel *models.Artikel

	err := c.Bind(&Artikel)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	Artikel, err = a.ArtikelS.CreateService(*Artikel)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Artikel,
		Message: "Create Artikel success",
		Status:  true,
	})
}

func (a *artikelController) UpdateController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	var Artikel *models.Artikel

	err = c.Bind(&Artikel)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	Artikel, err = a.ArtikelS.UpdateService(id, *Artikel)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Artikel,
		Message: "Update Artikel success",
		Status:  true,
	})
}

func (a *artikelController) DeleteController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	err = a.ArtikelS.DeleteService(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    nil,
		Message: "Delete Artikel success",
		Status:  true,
	})
}
