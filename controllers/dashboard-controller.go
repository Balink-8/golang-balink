package controllers

import (
	h "capstone/helpers"
	"capstone/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DashboardController interface {
	DashboardGetAll(c echo.Context) error
}

type dashboardController struct {
	dashboardServices services.DashboardServices
}

func NewDashboardController(dashboardServices services.DashboardServices) DashboardController {
	return &dashboardController{dashboardServices}
}

func (c *dashboardController) DashboardGetAll(ctx echo.Context) error {
	dashboard, err := c.dashboardServices.DashboardGetAll()
	if err != nil {
		return h.Response(ctx, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(ctx, http.StatusBadRequest, h.ResponseModel{
		Data:    dashboard,
		Message: "Successfully get all dashboards",
		Status:  true,
	})
}
