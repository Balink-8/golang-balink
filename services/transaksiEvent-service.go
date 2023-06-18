package services

import (
	"capstone/models"
	"capstone/repositories"
	"fmt"
	"strconv"
)

type TransaksiEventServices interface {
	GetTransaksiEventsService(page int, limit int, order string) ([]*models.TransaksiEvent, int, error)
	GetTransaksiEventService(id string) (*models.TransaksiEvent, error)
	CreateTransaksiEvent(id_keranjang uint, TransaksiEvent models.TransaksiEvent) (*models.TransaksiEvent, error)
	GetTransaksiEventByUserServices(User_ID string) ([]*models.TransaksiEvent, error)
}

type transaksiEventServices struct {
	TransaksiEventRepo repositories.TransaksiEventRepository
	KeranjangRepo   repositories.KeranjangRepository
	ProdukRepo      repositories.ProdukRepository
}

func NewTransaksiEventServices(TransaksiEventRepo repositories.TransaksiEventRepository, KeranjangRepo repositories.KeranjangRepository, ProdukRepo repositories.ProdukRepository) *transaksiEventServices {
	return &transaksiEventServices{
		TransaksiEventRepo: TransaksiEventRepo,
		KeranjangRepo:   KeranjangRepo,
		ProdukRepo:      ProdukRepo,
	}
}

func (t *transaksiEventServices) GetTransaksiEventsService(page int, limit int, order string) ([]*models.TransaksiEvent, int, error) {
	TransaksiEvents, totalData, err := t.TransaksiEventRepo.GetTransaksiEventsRepository(page, limit, order)
	if err != nil {
		return nil, 0, err
	}

	return TransaksiEvents, totalData, nil
}

func (t *transaksiEventServices) GetTransaksiEventService(id string) (*models.TransaksiEvent, error) {
	TransaksiEvent, err := t.TransaksiEventRepo.GetTransaksiEventRepository(id)
	if err != nil {
		return nil, err
	}

	return TransaksiEvent, nil
}

func (t *transaksiEventServices) CreateTransaksiEvent(id_keranjang uint, TransaksiEvent models.TransaksiEvent) (*models.TransaksiEvent, error) {
	stringkeranjang := strconv.Itoa(int(id_keranjang))
	keranjang, err := t.KeranjangRepo.GetKeranjangRepository(stringkeranjang)
	if err != nil {
		return nil, err
	}

	ProdukID := keranjang.Produk_ID
	TotalBarang := keranjang.Jumlah
	Produk, err := t.ProdukRepo.GetProdukRepository(ProdukID)
	if err != nil {
		return nil, err
	}
	ProdukHarga := Produk.Harga

	TotalHargaProduk := float64(ProdukHarga) * float64(TotalBarang)
	DescriptionTransaksiEvent := fmt.Sprintf("Total harga barang kamu adalah %.2f", TotalHargaProduk)

	IDKeranjang := id_keranjang

	TransaksiEventBaru := models.TransaksiEvent{
		KeranjangID: IDKeranjang,
		Description: DescriptionTransaksiEvent,
		Status:      "unpaid",
	}

	TransaksiEventR, err := t.TransaksiEventRepo.CreateTransaksiEvent(TransaksiEventBaru)
	if err != nil {
		return nil, err
	}

	return TransaksiEventR, nil
}

func (t *transaksiEventServices) GetTransaksiEventByUserServices(User_ID string) ([]*models.TransaksiEvent, error) {
	TransaksiEvents, err := t.TransaksiEventRepo.GetTransaksiEventByUserRepository(User_ID)
	if err != nil {
		return nil, err
	}

	return TransaksiEvents, nil
}
