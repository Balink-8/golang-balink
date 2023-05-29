package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	h "capstone/helpers"
	"capstone/models"
	"capstone/services"
)

type ProdukController interface {
	GetProduksController(c echo.Context) error
	GetProdukController(c echo.Context) error
	CreateController(c echo.Context) error
	UpdateController(c echo.Context) error
	DeleteController(c echo.Context) error
}

type produkController struct {
	ProdukS services.ProdukService
}

func NewProdukController(ProdukS services.ProdukService) ProdukController {
	return &produkController{
		ProdukS: ProdukS,
	}
}

func (pr *produkController) GetProduksController(c echo.Context) error {
	Produks, err := pr.ProdukS.GetProduksService()
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Produks,
		Message: "Get all Produks success",
		Status:  true,
	})
}

func (pr *produkController) GetProdukController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	var Produk *models.Produk

	Produk, err = pr.ProdukS.GetProdukService(id)
	if err != nil {
		return h.Response(c, http.StatusNotFound, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Produk,
		Message: "Get Produk success",
		Status:  true,
	})
}

func (pr *produkController) CreateController(c echo.Context) error {
	var Produk *models.Produk

	err := c.Bind(&Produk)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	Produk, err = pr.ProdukS.CreateService(*Produk)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Produk,
		Message: "Create Produk success",
		Status:  true,
	})
}

func (pr *produkController) UpdateController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	var Produk *models.Produk

	err = c.Bind(&Produk)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	Produk, err = pr.ProdukS.UpdateService(id, *Produk)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Produk,
		Message: "Update Produk success",
		Status:  true,
	})
}

func (pr *produkController) DeleteController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	err = pr.ProdukS.DeleteService(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    nil,
		Message: "Delete Produk success",
		Status:  true,
	})
}
