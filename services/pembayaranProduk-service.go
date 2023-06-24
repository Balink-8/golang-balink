package services

import (
	"capstone/models"
	"capstone/repositories"
	"context"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type PembayaranProdukService interface {
	CreateService(PembayaranProduk models.PembayaranProduk) (*models.PembayaranProduk, error)
	Credentials() *cloudinary.Cloudinary
	UploadBuktiPembayaran(fileHeader *multipart.FileHeader, id int) (pembayaranProduk *models.PembayaranProduk, err error)
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

func (pr *pembayaranProdukService) Credentials() *cloudinary.Cloudinary {
	cld, _ := cloudinary.New()
	cld.Config.URL.Secure = true
	return cld
}

func (pr *pembayaranProdukService) UploadBuktiPembayaran(fileHeader *multipart.FileHeader, id int) (pembayaranProduk *models.PembayaranProduk, err error) {

	file, _ := fileHeader.Open()

	cld := pr.Credentials()

	resp, err := cld.Upload.Upload(context.Background(), file, uploader.UploadParams{
		PublicID:       "Balink" + "/" + fileHeader.Filename,
		UniqueFilename: api.Bool(false),
		Overwrite:      api.Bool(true),
	})

	if err != nil {
		return
	}

	pembayaranProduk, err = pr.PP.GetPembayaranProdukRepository(id)
	if err != nil {
		return
	}

	pembayaranProduk.BuktiPembayaran = resp.SecureURL
	pembayaranProduk.Status = "Menunggu"
	err = pr.PP.UpdateRepository(*pembayaranProduk)
	if err != nil {
		return
	}

	return pembayaranProduk, nil
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
