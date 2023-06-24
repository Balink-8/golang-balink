package controllers

import (
	"fmt"
	"net/http"
	"regexp"
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

func (K *keranjangController) CreateController(c echo.Context) error {
	var keranjang models.Keranjang

	fmt.Println("Data :", &keranjang)

	err := c.Bind(&keranjang)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	file, err := c.FormFile("image") // Mengubah ctx menjadi c pada bagian ini

	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: "Image cannot be empty", // Mengubah pesan error menjadi string statis
			Status:  false,
		})
	}

	src, err := file.Open()
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: "Failed to open file", // Mengubah pesan error menjadi string statis
			Status:  false,
		})
	}

	re := regexp.MustCompile(`.png|.jpeg|.jpg`)

	if !re.MatchString(file.Filename) {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: "The provided file format is not allowed. Please upload a JPEG or PNG image", // Mengubah pesan error menjadi string statis
			Status:  false,
		})
	}

	uploadUrl, err := services.NewMediaUpload().FileUpload(models.File{File: src})
	if err != nil {
		return h.Response(c, http.StatusInternalServerError, h.ResponseModel{
			Data:    nil,
			Message: "Error uploading photo", // Mengubah pesan error menjadi string statis
			Status:  false,
		})
	}
	keranjang.Image = uploadUrl // Mengubah artikelInput menjadi Produk

	keranjang, err = K.KeranjangS.CreateService(keranjang)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    keranjang,
		Message: "Create Artikel success",
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
