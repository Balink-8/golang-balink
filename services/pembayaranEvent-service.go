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

type PembayaranEventService interface {
	CreateService(PembayaranEvent models.PembayaranEvent) (*models.PembayaranEvent, error)
	Credentials() *cloudinary.Cloudinary
	UploadBuktiPembayaran(fileHeader *multipart.FileHeader, id int) (pembayaranEvent *models.PembayaranEvent, err error)
}

type pembayaranEventService struct {
	PP      repositories.PembayaranEventRepository
	K       repositories.KeranjangRepository
	payment repositories.PaymentMethodRepository
}

func NewPembayaranEventService(PP repositories.PembayaranEventRepository, K repositories.KeranjangRepository, payment repositories.PaymentMethodRepository) PembayaranEventService {
	return &pembayaranEventService{
		PP:      PP,
		K:       K,
		payment: payment,
	}

}

func (pr *pembayaranEventService) Credentials() *cloudinary.Cloudinary {
	cld, _ := cloudinary.New()
	cld.Config.URL.Secure = true
	return cld
}

func (pr *pembayaranEventService) UploadBuktiPembayaran(fileHeader *multipart.FileHeader, id int) (pembayaranEvent *models.PembayaranEvent, err error) {

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

	pembayaranEvent, err = pr.PP.GetPembayaranEventRepository(id)
	if err != nil {
		return
	}

	pembayaranEvent.BuktiPembayaran = resp.SecureURL
	pembayaranEvent.Status = "Menunggu"
	err = pr.PP.UpdateRepository(*pembayaranEvent)
	if err != nil {
		return
	}

	return pembayaranEvent, nil
}

func (pr *pembayaranEventService) CreateService(PembayaranEvent models.PembayaranEvent) (*models.PembayaranEvent, error) {

	keranjang, err := pr.K.GetKeranjangRepository(PembayaranEvent.KeranjangID)
	if err != nil {
		return nil, err
	}

	payment, err := pr.payment.GetPaymentMethodRepository(PembayaranEvent.MetodePembayaranID)
	if err != nil {
		return nil, err
	}

	payload := models.PembayaranEvent{
		KeranjangTiket:   *keranjang,
		Status:           `unpaid`,
		PromoID:          PembayaranEvent.PromoID,
		MetodePembayaran: *payment,
	}

	PembayaranEventodukR, err := pr.PP.CreatePembayaranEvent(payload)
	if err != nil {
		return nil, err
	}

	return PembayaranEventodukR, nil
}
