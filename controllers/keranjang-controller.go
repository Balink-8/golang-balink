package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	h "capstone/helpers"
	"capstone/models"
	"capstone/services"
)

type KeranjangController interface {
	GetKeranjangsController(c echo.Context) error
	GetKeranjangController(c echo.Context) error
	CreateController(c echo.Context) error
	UpdateController(c echo.Context) error
	DeleteController(c echo.Context) error
	GetKeranjangByUserController(c echo.Context) error
}

type keranjangController struct {
	KeranjangS services.KeranjangService
}

func NewKeranjangController(KeranjangS services.KeranjangService) KeranjangController {
	return &keranjangController{
		KeranjangS: KeranjangS,
	}
}

func (k *keranjangController) GetKeranjangsController(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	order := c.QueryParam("order")
	search := c.QueryParam("search")

	Keranjangs, totalData, err := k.KeranjangS.GetKeranjangsService(page, limit, order, search)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	responseData := map[string]interface{}{
		"data":       Keranjangs,
		"page":       page,
		"data_shown": len(Keranjangs),
		"total_data": totalData,
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    responseData,
		Message: "Get all Keranjang success",
		Status:  true,
	})
}

func (k *keranjangController) GetKeranjangController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	var Keranjang *models.Keranjang

	Keranjang, err = k.KeranjangS.GetKeranjangService(id)
	if err != nil {
		return h.Response(c, http.StatusNotFound, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Keranjang,
		Message: "Get Keranjang success",
		Status:  true,
	})
}

func (k *keranjangController) CreateController(c echo.Context) error {
	var Keranjang *models.Keranjang

	err := c.Bind(&Keranjang)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	Keranjang, err = k.KeranjangS.CreateService(*Keranjang)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Keranjang,
		Message: "Create Keranjang success",
		Status:  true,
	})
}

func (k *keranjangController) UpdateController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	var Keranjang *models.Keranjang

	err = c.Bind(&Keranjang)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	Keranjang, err = k.KeranjangS.UpdateService(id, *Keranjang)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Keranjang,
		Message: "Update Keranjang success",
		Status:  true,
	})
}

func (k *keranjangController) DeleteController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	err = k.KeranjangS.DeleteService(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    nil,
		Message: "Delete Keranjang success",
		Status:  true,
	})
}

func (k *keranjangController) GetKeranjangByUserController(c echo.Context) error {
	id_user := c.Param("id_user")
	err := h.IsNumber(id_user)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}
	Keranjangs, err := k.KeranjangS.GetKeranjangByUserService(id_user)
	if err != nil {
		return h.Response(c, http.StatusNotFound, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}
	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Keranjangs,
		Message: "Get Keranjang By User success",
		Status:  true,
	})
}
