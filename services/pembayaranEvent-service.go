package services

import (
	"capstone/models"
	"capstone/repositories"
)

type PembayaranEventService interface {
	CreateService(PembayaranEvent models.PembayaranEvent) (*models.PembayaranEvent, error)
}

type pembayaranEventService struct {
	PP      repositories.PembayaranEventRepository
	K       repositories.UserRepository
	payment repositories.PaymentMethodRepository
}

func NewPembayaranEventService(PP repositories.PembayaranEventRepository, K repositories.UserRepository, payment repositories.PaymentMethodRepository) PembayaranEventService {
	return &pembayaranEventService{
		PP: PP,
		// K:       K,
		payment: payment,
	}

}

func (pr *pembayaranEventService) CreateService(PembayaranEvent models.PembayaranEvent) (*models.PembayaranEvent, error) {

	user_id, err := pr.K.GetUserRepository(PembayaranEvent.User_ID)
	if err != nil {
		return nil, err
	}

	payment, err := pr.payment.GetPaymentMethodRepository(PembayaranEvent.MetodePembayaranID)
	if err != nil {
		return nil, err
	}

	payload := models.PembayaranEvent{
		User_ID:          *user,
		Status:           `unpaid`,
		AlamatPengiriman: PembayaranEvent.AlamatPengiriman,
		Pesan:            PembayaranEvent.Pesan,
		PromoID:          PembayaranEvent.PromoID,
		MetodePembayaran: *payment,
	}

	PembayaranEventodukR, err := pr.PP.CreatePembayaranEvent(payload)
	if err != nil {
		return nil, err
	}

	return PembayaranEventodukR, nil
}
