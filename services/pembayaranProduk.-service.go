package services

import (
	"capstone/models"
	"capstone/repositories"
)

type PembayaranProdukService interface {
	CreateService(PembayaranProduk models.PembayaranProduk) (*models.PembayaranProduk, error)
}

type pembayaranProdukService struct {
	PP      repositories.PembayaranProdukRepository
	K       repositories.KeranjangRepository
	payment repositories.PaymentMethodRepository
}

func NewPembayaranProdukService(PP repositories.PembayaranProdukRepository, K repositories.KeranjangRepository, payment repositories.PaymentMethodRepository) PembayaranProdukService {
	return &pembayaranProdukService{
		PP:      PP,
		K:       K,
		payment: payment,
	}

}

func (pr *pembayaranProdukService) CreateService(PembayaranProduk models.PembayaranProduk) (*models.PembayaranProduk, error) {

	keranjang, err := pr.K.GetKeranjangRepository(PembayaranProduk.KeranjangID)
	if err != nil {
		return nil, err
	}

	payment, err := pr.payment.GetPaymentMethodRepository(PembayaranProduk.MetodePembayaranID)
	if err != nil {
		return nil, err
	}

	payload := models.PembayaranProduk{
		Keranjang:        *keranjang,
		Status:           `unpaid`,
		AlamatPengiriman: PembayaranProduk.AlamatPengiriman,
		Pesan:            PembayaranProduk.Pesan,
		PromoID:          PembayaranProduk.PromoID,
		MetodePembayaran: *payment,
	}

	PembayaranProdukodukR, err := pr.PP.CreatePembayaranProduk(payload)
	if err != nil {
		return nil, err
	}

	return PembayaranProdukodukR, nil
}
