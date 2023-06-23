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

	Promos, totalData, err := p.PromoS.GetPromosService(page, limit, order, search)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	responseData := map[string]interface{}{
		"data":       Promos,
		"page":       page,
		"data_shown": len(Promos),
		"total_data": totalData,
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    responseData,
		Message: "Get all Promo success",
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
	var promo models.Promo

	err := c.Bind(&promo)
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
	promo.Image = uploadURL

	createdPromo, err := p.PromoS.CreateService(promo)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, h.ResponseModel{
		Data:    createdPromo,
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
