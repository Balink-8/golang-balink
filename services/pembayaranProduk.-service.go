package services

import (
	"capstone/models"
	"capstone/repositories"
)

type PembayaranProdukService interface {
}

type pembayaranProdukService struct {
	PP repositories.PembayaranProdukRepository
	K repositories.KeranjangRepository
}

func NewPembayaranProdukService(PP repositories.PembayaranProdukRepository, K repositories.KeranjangRepository) PembayaranProdukService {
	return &pembayaranProdukService{
		PP: PP,
		K: K,
	}

}

func (pr *pembayaranProdukService) CreateService(PembayaranProduk models.PembayaranProduk) (*models.PembayaranProduk, error) {
	keranjang, err := pr.K.GetKeranjangRepository(PembayaranProduk.KeranjangID)
	if err != nil {
		return nil, err
	}
    
	payload := models.PembayaranProduk{
		KeranjangID: int64(keranjang.ID),
		Status: `unpaid`,
		AlamatPengiriman: PembayaranProduk.AlamatPengiriman,
		Pesan: PembayaranProduk.Pesan,
		PromoID: PembayaranProduk.PromoID,
	}

	PembayaranProdukodukR, err := pr.PP.CreatePembayaranProduk(payload)
	if err != nil {
		return nil, err
	}

	return PembayaranProdukodukR, nil
}
