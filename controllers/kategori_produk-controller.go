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

	KategoriProduks, totalData, err := k.KategoriProdukS.GetKategoriProduksService(page, limit, order, search)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	responseData := map[string]interface{}{
		"data":       KategoriProduks,
		"page":       page,
		"data_shown": len(KategoriProduks),
		"total_data": totalData,
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    responseData,
		Message: "Get all Kategori Produk success",
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
	var kategoriProduk models.KategoriProduk

	err := c.Bind(&kategoriProduk)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	file, err := c.FormFile("image")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Image cannot be empty", err)
	}

	src, err := file.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to open file", err)
	}

	re := regexp.MustCompile(`.png|.jpeg|.jpg`)
	if !re.MatchString(file.Filename) {
		return echo.NewHTTPError(http.StatusBadRequest, "The provided file format is not allowed. Please upload a JPEG or PNG image")
	}

	uploadURL, err := services.NewMediaUpload().FileUpload(models.File{File: src})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error uploading photo", err)
	}
	kategoriProduk.Image = uploadURL

	createdKategoriProduk, err := k.KategoriProdukS.CreateService(kategoriProduk)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, h.ResponseModel{
		Data:    createdKategoriProduk,
		Message: "Create Kategori Produk success",
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
