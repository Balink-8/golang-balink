package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	h "capstone/helpers"
	"capstone/models"
	"capstone/services"
)

type KategoriProdukController interface {
	GetKategoriProduksController(c echo.Context) error
	GetKategoriProdukController(c echo.Context) error
	CreateController(c echo.Context) error
	UpdateController(c echo.Context) error
	DeleteController(c echo.Context) error
}

type kategoriProdukController struct {
	KategoriProdukS services.KategoriProdukService
}

func NewKategoriProdukController(KategoriProdukS services.KategoriProdukService) KategoriProdukController {
	return &kategoriProdukController{
		KategoriProdukS: KategoriProdukS,
	}
}

func (k *kategoriProdukController) GetKategoriProduksController(c echo.Context) error {
	KategoriProduks, err := k.KategoriProdukS.GetKategoriProduksService()
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    KategoriProduks,
		Message: "Get all KategoriProduks success",
		Status:  true,
	})
}

func (k *kategoriProdukController) GetKategoriProdukController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	var KategoriProduk *models.KategoriProduk

	KategoriProduk, err = k.KategoriProdukS.GetKategoriProdukService(id)
	if err != nil {
		return h.Response(c, http.StatusNotFound, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    KategoriProduk,
		Message: "Get KategoriProduk success",
		Status:  true,
	})
}

func (k *kategoriProdukController) CreateController(c echo.Context) error {
	var KategoriProduk *models.KategoriProduk

	err := c.Bind(&KategoriProduk)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	KategoriProduk, err = k.KategoriProdukS.CreateService(*KategoriProduk)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    KategoriProduk,
		Message: "Create KategoriProduk success",
		Status:  true,
	})
}

func (k *kategoriProdukController) UpdateController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	var KategoriProduk *models.KategoriProduk

	err = c.Bind(&KategoriProduk)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	KategoriProduk, err = k.KategoriProdukS.UpdateService(id, *KategoriProduk)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    KategoriProduk,
		Message: "Update KategoriProduk success",
		Status:  true,
	})
}

func (k *kategoriProdukController) DeleteController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	err = k.KategoriProdukS.DeleteService(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    nil,
		Message: "Delete KategoriProduk success",
		Status:  true,
	})
}
