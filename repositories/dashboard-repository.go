package repositories

import (
	"time"

	"gorm.io/gorm"
)

type DashboardRepository interface {
	GetCountUserSebulanRepository() (float64, error)
	// GetCountTransaksiTiketSebulanRepository() (float64, error)
	// GetSumTransaksiProdukSebulanRepository() (float64, error)
	GetCountArtikelSebulanRepository() (float64, error)
	// GetTransaksiSuksesRepository() (float64, error)
	// GetTransaksiDipesanRepository() (float64, error)
	// GetTransaksiDibatalkanRepository() (float64, error)
}

type dashboardRepository struct {
	DB *gorm.DB
}

func NewDashboardRepository(DB *gorm.DB) DashboardRepository {
	return &dashboardRepository{
		DB: DB,
	}
}

func (d *dashboardRepository) GetCountUserSebulanRepository() (float64, error) {
	var countUserSebulan float64
	lastMonth := time.Now().AddDate(0, -1, 0)

	query := `
		SELECT COUNT(*) FROM users
		WHERE created_at >= ? AND deleted_at is NULL;
	`
	if err := d.DB.Raw(query, lastMonth).Scan(&countUserSebulan).Error; err != nil {
		return 0, err
	}

	return countUserSebulan, nil
}

// func (d *dashboardRepository) GetCountTransaksiTiketSebulanRepository() (float64, error) {
// 	var countTransaksiTiketSebulan float64
// 	lastMonth := time.Now().AddDate(0, -1, 0)

// 	query := `
// 		SELECT COUNT(*) FROM pembayaran_events
// 		WHERE created_at >= ? AND deleted_at is NULL;
// 	`
// 	if err := d.DB.Raw(query, lastMonth).Scan(&countTransaksiTiketSebulan).Error; err != nil {
// 		return 0, err
// 	}

// 	return countTransaksiTiketSebulan, nil
// }

// func (d *dashboardRepository) GetSumTransaksiProdukSebulanRepository() (float64, error) {
// 	var sumTransaksiProdukSebulan float64
// 	lastMonth := time.Now().AddDate(0, -1, 0)

// 	query := `
// 		SELECT SUM(k.jumlah)
// 		FROM keranjangs k JOIN transaksi_produks t ON k.id=t.keranjang_id
// 		WHERE k.created_at >= ? AND k.deleted_at is NULL;
// 	`
// 	if err := d.DB.Raw(query, lastMonth).Scan(&sumTransaksiProdukSebulan).Error; err != nil {
// 		return 0, err
// 	}

// 	return sumTransaksiProdukSebulan, nil
// }

func (d *dashboardRepository) GetCountArtikelSebulanRepository() (float64, error) {
	var countArtikelSebulan float64
	lastMonth := time.Now().AddDate(0, -1, 0)

	query := `
		SELECT COUNT(*) FROM artikels
		WHERE created_at >= ? AND deleted_at is NULL;
	`
	if err := d.DB.Raw(query, lastMonth).Scan(&countArtikelSebulan).Error; err != nil {
		return 0, err
	}

	return countArtikelSebulan, nil
}

// func (d *dashboardRepository) GetTransaksiSuksesRepository() (float64, error) {
// 	var countTransaksiProdukSuksesSebulan float64
// 	var countTransaksiProdukSebulan float64
// 	var countTransaksiEventSuksesSebulan float64
// 	var countTransaksiEventSebulan float64
// 	lastMonth := time.Now().AddDate(0, -1, 0)

// 	queryProduk := `
// 		SELECT COUNT(IF('status' = 'Sukses', 1, NULL)) FROM pembayaran_produks
// 		WHERE created_at >= ? AND deleted_at is NULL;
// 	`
// 	if err := d.DB.Raw(queryProduk, lastMonth).Scan(&countTransaksiProdukSuksesSebulan).Error; err != nil {
// 		return 0, err
// 	}

// 	queryTotalProduk := `
// 		SELECT COUNT(*) FROM pembayaran_produks
// 		WHERE created_at >= ? AND deleted_at is NULL;
// 	`
// 	if err := d.DB.Raw(queryTotalProduk, lastMonth).Scan(&countTransaksiProdukSebulan).Error; err != nil {
// 		return 0, err
// 	}

// 	queryEvent := `
// 		SELECT COUNT(IF('status' = 'Sukses', 1, NULL)) FROM pembayaran_events
// 		WHERE created_at >= ? AND deleted_at is NULL;
// 	`
// 	if err := d.DB.Raw(queryEvent, lastMonth).Scan(&countTransaksiProdukSuksesSebulan).Error; err != nil {
// 		return 0, err
// 	}

// 	queryTotalEvent := `
// 		SELECT COUNT(*) FROM pembayaran_events
// 		WHERE created_at >= ? AND deleted_at is NULL;
// 	`
// 	if err := d.DB.Raw(queryTotalEvent, lastMonth).Scan(&countTransaksiProdukSebulan).Error; err != nil {
// 		return 0, err
// 	}

// 	var hasilTransaksiSukses float64 = ((countTransaksiProdukSuksesSebulan/countTransaksiProdukSebulan) + (countTransaksiEventSuksesSebulan/countTransaksiEventSebulan))/2 * 100

// 	return hasilTransaksiSukses, nil
// }

// func (d *dashboardRepository) GetTransaksiDipesanRepository() (float64, error) {
// 	var countTransaksiProdukDipesanSebulan float64
// 	var countTransaksiProdukSebulan float64
// 	var countTransaksiEventDipesanSebulan float64
// 	var countTransaksiEventSebulan float64
// 	lastMonth := time.Now().AddDate(0, -1, 0)

// 	queryProduk := `
// 		SELECT COUNT(IF('status' = 'Dipesan', 1, NULL)) FROM pembayaran_produks
// 		WHERE created_at >= ? AND deleted_at is NULL;
// 	`
// 	if err := d.DB.Raw(queryProduk, lastMonth).Scan(&countTransaksiProdukDipesanSebulan).Error; err != nil {
// 		return 0, err
// 	}

// 	queryTotalProduk := `
// 		SELECT COUNT(*) FROM pembayaran_produks
// 		WHERE created_at >= ? AND deleted_at is NULL;
// 	`
// 	if err := d.DB.Raw(queryTotalProduk, lastMonth).Scan(&countTransaksiProdukSebulan).Error; err != nil {
// 		return 0, err
// 	}

// 	queryEvent := `
// 		SELECT COUNT(IF('status' = 'Dipesan', 1, NULL)) FROM pembayaran_events
// 		WHERE created_at >= ? AND deleted_at is NULL;
// 	`
// 	if err := d.DB.Raw(queryEvent, lastMonth).Scan(&countTransaksiProdukDipesanSebulan).Error; err != nil {
// 		return 0, err
// 	}

// 	queryTotalEvent := `
// 		SELECT COUNT(*) FROM pembayaran_events
// 		WHERE created_at >= ? AND deleted_at is NULL;
// 	`
// 	if err := d.DB.Raw(queryTotalEvent, lastMonth).Scan(&countTransaksiProdukSebulan).Error; err != nil {
// 		return 0, err
// 	}

// 	var hasilTransaksiDipesan float64 = ((countTransaksiProdukDipesanSebulan/countTransaksiProdukSebulan) + (countTransaksiEventDipesanSebulan/countTransaksiEventSebulan))/2 * 100

// 	return hasilTransaksiDipesan, nil
// }

// func (d *dashboardRepository) GetTransaksiDibatalkanRepository() (float64, error) {
// 	var countTransaksiProdukDibatalkanSebulan float64
// 	var countTransaksiProdukSebulan float64
// 	var countTransaksiEventDibatalkanSebulan float64
// 	var countTransaksiEventSebulan float64
// 	lastMonth := time.Now().AddDate(0, -1, 0)

// 	queryProduk := `
// 		SELECT COUNT(IF('status' = 'Dibatalkan', 1, NULL)) FROM pembayaran_produks
// 		WHERE created_at >= ? AND deleted_at is NULL;
// 	`
// 	if err := d.DB.Raw(queryProduk, lastMonth).Scan(&countTransaksiProdukDibatalkanSebulan).Error; err != nil {
// 		return 0, err
// 	}

// 	queryTotalProduk := `
// 		SELECT COUNT(*) FROM pembayaran_produks
// 		WHERE created_at >= ? AND deleted_at is NULL;
// 	`
// 	if err := d.DB.Raw(queryTotalProduk, lastMonth).Scan(&countTransaksiProdukSebulan).Error; err != nil {
// 		return 0, err
// 	}

// 	queryEvent := `
// 		SELECT COUNT(IF('status' = 'Dibatalkan', 1, NULL)) FROM pembayaran_events
// 		WHERE created_at >= ? AND deleted_at is NULL;
// 	`
// 	if err := d.DB.Raw(queryEvent, lastMonth).Scan(&countTransaksiProdukDibatalkanSebulan).Error; err != nil {
// 		return 0, err
// 	}

// 	queryTotalEvent := `
// 		SELECT COUNT(*) FROM pembayaran_events
// 		WHERE created_at >= ? AND deleted_at is NULL;
// 	`
// 	if err := d.DB.Raw(queryTotalEvent, lastMonth).Scan(&countTransaksiProdukSebulan).Error; err != nil {
// 		return 0, err
// 	}

// 	var hasilTransaksiDibatalkan float64 = ((countTransaksiProdukDibatalkanSebulan/countTransaksiProdukSebulan) + (countTransaksiEventDibatalkanSebulan/countTransaksiEventSebulan))/2 * 100

// 	return hasilTransaksiDibatalkan, nil
// }
