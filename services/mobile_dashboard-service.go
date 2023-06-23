package services

import (
	"capstone/helpers"
	"capstone/repositories"
)

type DashboardMobileServices interface {
	DashboardGetAll() (helpers.DashboardMobileResponse, error)
}

type dashboardMobileServices struct {
	dashboardMobileRepository repositories.DashboarMobiledRepository
}

func NewDashboardMobileServices(dashboardMobileRepository repositories.DashboarMobiledRepository) *dashboardMobileServices {
	return &dashboardMobileServices{dashboardMobileRepository}
}

func (u *dashboardMobileServices) DashboardGetAll() (helpers.DashboardMobileResponse, error) {
	dashboardResponse := helpers.DashboardMobileResponse{}

	kategori, promos, events, err := u.dashboardMobileRepository.DashboardMobileGetAll()
	if err != nil {
		return dashboardResponse, err
	}

	dashboardResponse.Kategori = kategori
	dashboardResponse.Promos = promos
	dashboardResponse.Events = events

	return dashboardResponse, nil
}
