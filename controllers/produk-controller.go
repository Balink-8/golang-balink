package controllers

import (
	"net/http"
	"regexp"
	"strconv"

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

func (p *produkController) GetProduksController(c echo.Context) error {
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

	Produks, totalData, err := p.ProdukS.GetProduksService(page, limit, order, search)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	responseData := map[string]interface{}{
		"data":       Produks,
		"page":       page,
		"data_shown": len(Produks),
		"total_data": totalData,
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    responseData,
		Message: "Get all Produk success",
		Status:  true,
	})
}

func (p *produkController) GetProdukController(c echo.Context) error {
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

	Produk, err = p.ProdukS.GetProdukService(id)
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

func (p *produkController) CreateController(c echo.Context) error {
	var Produk *models.Produk

	err := c.Bind(&Produk)
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
	Produk.Gambar = uploadUrl // Mengubah artikelInput menjadi Produk

	Produk, err = p.ProdukS.CreateService(*Produk)
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

func (p *produkController) UpdateController(c echo.Context) error {
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

	Produk, err = p.ProdukS.UpdateService(id, *Produk)
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

func (p *produkController) DeleteController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	err = p.ProdukS.DeleteService(id)
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
