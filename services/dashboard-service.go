package services

import (
	"capstone/helpers"
	"capstone/repositories"
)

type DashboardServices interface {
	DashboardGetAll() (helpers.DashboardResponse, error)
}

type dashboardServices struct {
	dashboardRepository repositories.DashboardRepository
}

func NewDashboardServices(dashboardRepository repositories.DashboardRepository) DashboardServices {
	return &dashboardServices{dashboardRepository}
}

func (u *dashboardServices) DashboardGetAll() (helpers.DashboardResponse, error) {
	dashboardResponse := helpers.DashboardResponse{}

	artikels, produks, promos, events, err := u.dashboardRepository.DashboardGetAll()
	if err != nil {
		return dashboardResponse, err
	}

	dashboardResponse.Artikels = artikels
	dashboardResponse.Produks = produks
	dashboardResponse.Promos = promos
	dashboardResponse.Events = events

	return dashboardResponse, nil
}
