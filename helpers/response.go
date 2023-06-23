package helpers

import (
	"capstone/models"

	"github.com/labstack/echo/v4"
)

type ResponseModel struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  bool        `json:"status"`
}

type DashboardMobileResponse struct {
	Kategori []models.KategoriProduk `json:"kategori"`
	Promos   []models.Promo          `json:"promos"`
	Events   []models.Event          `json:"events"`
}

func Response(c echo.Context, code int, model ResponseModel) error {
	return c.JSON(code, model)
}
