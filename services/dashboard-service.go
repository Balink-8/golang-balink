package services

import (
	"capstone/repositories"
)

type DashboardService interface {
	GetCountUserSebulanService() (float64, error)
	// GetCountTransaksiTiketSebulanService() (float64, error)
	// GetSumTransaksiProdukSebulanService() (float64, error)
	GetCountArtikelSebulanService() (float64, error)
	// GetTransaksiSuksesService() (float64, error)
	// GetTransaksiDipesanService() (float64, error)
	// GetTransaksiDibatalkanService() (float64, error)
}

type dashboardService struct {
	DashboardR repositories.DashboardRepository
}

func NewDashboardService(DashboardR repositories.DashboardRepository) DashboardService {
	return &dashboardService{
		DashboardR: DashboardR,
	}
}

func (d *dashboardService) GetCountUserSebulanService() (float64, error) {
	countUserSebulan, err := d.DashboardR.GetCountUserSebulanRepository()
	if err != nil {
		return 0, err
	}

	return countUserSebulan, nil
}

// func (d *dashboardService) GetCountTransaksiTiketSebulanService() (float64, error) {
// 	countTransaksiTiketSebulan, err := d.DashboardR.GetCountTransaksiTiketSebulanRepository()
// 	if err != nil {
// 		return 0, err
// 	}

// 	return countTransaksiTiketSebulan, nil
// }

// func (d *dashboardService) GetSumTransaksiProdukSebulanService() (float64, error) {
// 	sumTransaksiProdukSebulan, err := d.DashboardR.GetSumTransaksiProdukSebulanRepository()
// 	if err != nil {
// 		return 0, err
// 	}

// 	return sumTransaksiProdukSebulan, nil
// }

func (d *dashboardService) GetCountArtikelSebulanService() (float64, error) {
	countArtikelSebulan, err := d.DashboardR.GetCountArtikelSebulanRepository()
	if err != nil {
		return 0, err
	}

	return countArtikelSebulan, nil
}

// func (d *dashboardService) GetTransaksiSuksesService() (float64, error) {
// 	transaksiSukses, err := d.DashboardR.GetTransaksiSuksesRepository()
// 	if err != nil {
// 		return 0, err
// 	}

// 	return transaksiSukses, nil
// }

// func (d *dashboardService) GetTransaksiDipesanService() (float64, error) {
// 	transaksiDipesan, err := d.DashboardR.GetTransaksiDipesanRepository()
// 	if err != nil {
// 		return 0, err
// 	}

// 	return transaksiDipesan, nil
// }

// func (d *dashboardService) GetTransaksiDibatalkanService() (float64, error) {
// 	transaksiDibatalkan, err := d.DashboardR.GetTransaksiDibatalkanRepository()
// 	if err != nil {
// 		return 0, err
// 	}

// 	return transaksiDibatalkan, nil
// }