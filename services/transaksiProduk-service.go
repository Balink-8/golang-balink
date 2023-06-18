package services

import (
	"capstone/models"
	"capstone/repositories"
	"fmt"
	"strconv"
)

type TransaksiProdukServices interface {
	GetTransaksiProduksService(page int, limit int, order string) ([]*models.TransaksiProduk, int, error)
	GetTransaksiProdukService(id string) (*models.TransaksiProduk, error)
	CreateTransaksiProduk(id_keranjang uint, TransaksiProduk models.TransaksiProduk) (*models.TransaksiProduk, error)
	DeleteTransaksiProduk(id string) error
	GetTransaksiProdukByUserServices(User_ID string) ([]*models.TransaksiProduk, error)
}

type transaksiProdukServices struct {
	TransaksiProdukRepo repositories.TransaksiProdukRepository
	KeranjangRepo   repositories.KeranjangRepository
	ProdukRepo      repositories.ProdukRepository
}

func NewTransaksiProdukServices(TransaksiProdukRepo repositories.TransaksiProdukRepository, KeranjangRepo repositories.KeranjangRepository, ProdukRepo repositories.ProdukRepository) *transaksiProdukServices {
	return &transaksiProdukServices{
		TransaksiProdukRepo: TransaksiProdukRepo,
		KeranjangRepo:   KeranjangRepo,
		ProdukRepo:      ProdukRepo,
	}
}

func (t *transaksiProdukServices) GetTransaksiProduksService(page int, limit int, order string) ([]*models.TransaksiProduk, int, error) {
	TransaksiProduks, totalData, err := t.TransaksiProdukRepo.GetTransaksiProduksRepository(page, limit, order)
	if err != nil {
		return nil, 0, err
	}

	return TransaksiProduks, totalData, nil
}

func (t *transaksiProdukServices) GetTransaksiProdukService(id string) (*models.TransaksiProduk, error) {
	TransaksiProduk, err := t.TransaksiProdukRepo.GetTransaksiProdukRepository(id)
	if err != nil {
		return nil, err
	}

	return TransaksiProduk, nil
}

func (t *transaksiProdukServices) CreateTransaksiProduk(id_keranjang uint, TransaksiProduk models.TransaksiProduk) (*models.TransaksiProduk, error) {
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
	DescriptionTransaksiProduk := fmt.Sprintf("Total harga barang kamu adalah %.2f", TotalHargaProduk)

	IDKeranjang := id_keranjang

	TransaksiProdukBaru := models.TransaksiProduk{
		KeranjangID: IDKeranjang,
		Description: DescriptionTransaksiProduk,
		Status:      "unpaid",
	}

	TransaksiProdukR, err := t.TransaksiProdukRepo.CreateTransaksiProduk(TransaksiProdukBaru)
	if err != nil {
		return nil, err
	}

	return TransaksiProdukR, nil
}

func (t *transaksiProdukServices) DeleteTransaksiProduk(id string) error {
	err := t.TransaksiProdukRepo.DeleteTransaksiProduk(id)
	if err != nil {
		return err
	}

	return nil
}

func (t *transaksiProdukServices) GetTransaksiProdukByUserServices(User_ID string) ([]*models.TransaksiProduk, error) {
	TransaksiProduks, err := t.TransaksiProdukRepo.GetTransaksiProdukByUserRepository(User_ID)
	if err != nil {
		return nil, err
	}

	return TransaksiProduks, nil
}
