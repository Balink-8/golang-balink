package controllers

import (
	h "capstone/helpers"
	"capstone/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DashboardMobileController interface {
	DashboardMobileGetAll(c echo.Context) error
}

type dashboardMobileController struct {
	dashboardMobileServices services.DashboardMobileServices
}

func NewDashboardMobileController(dashboardMobileServices services.DashboardMobileServices) *dashboardMobileController {
	return &dashboardMobileController{dashboardMobileServices}
}

func (c *dashboardMobileController) DashboardMobileGetAll(ctx echo.Context) error {
	dashboard, err := c.dashboardMobileServices.DashboardGetAll()
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
