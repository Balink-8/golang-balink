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

	Artikels, totalData, err := a.ArtikelS.GetArtikelsService(page, limit, order, search)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	responseData := map[string]interface{}{
		"data":       Artikels,
		"page":       page,
		"data_shown": len(Artikels),
		"total_data": totalData,
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    responseData,
		Message: "Get all Artikel success",
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

func (p *artikelController) CreateController(c echo.Context) error {
	var Artikel models.Artikel

	fmt.Println("Data :", &Artikel)

	err := c.Bind(&Artikel)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	file, err := c.FormFile("gambar") // Mengubah ctx menjadi c pada bagian ini

	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: "Gambar tidak boleh kosong", // Mengubah pesan error menjadi string statis
			Status:  false,
		})
	}

	src, err := file.Open()
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: "Gagal membuka file", // Mengubah pesan error menjadi string statis
			Status:  false,
		})
	}

	re := regexp.MustCompile(`.png|.jpeg|.jpg`)

	if !re.MatchString(file.Filename) {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: "Format file yang disediakan tidak diperbolehkan. Unggah gambar JPEG atau PNG", // Mengubah pesan error menjadi string statis
			Status:  false,
		})
	}

	uploadUrl, err := services.NewMediaUpload().FileUpload(models.File{File: src})
	if err != nil {
		return h.Response(c, http.StatusInternalServerError, h.ResponseModel{
			Data:    nil,
			Message: "Terjadi kesalahan saat mengunggah foto", // Mengubah pesan error menjadi string statis
			Status:  false,
		})
	}
	Artikel.Gambar = uploadUrl // Mengubah artikelInput menjadi Produk

	Artikel, err = p.ArtikelS.CreateService(Artikel)
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
