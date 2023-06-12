package services

import (
	"capstone/models"
	"capstone/repositories"
)

type TransactionService interface {
	SaveTransactionProduct(transactionProduct models.TransactionProduct) error
	DeleteTransactionProduct(transactionProduct models.TransactionProduct) error
	GetTransactionProducts(productID uint64) ([]*models.TransactionProduct, error)
	SaveTransationTicket(ticket models.Ticket) error
	GetTransactionTicket(customerID uint64) ([]*models.Ticket, error)
}

type Service struct {
	Repository repositories.TransactionRepository
}

func (s *Service) SaveTransactionProduct(transactionProduct models.TransactionProduct) error {
	return s.Repository.SaveTransactionProduct(transactionProduct)
}

func (s *Service) DeleteTransactionProduct(transactionProduct models.TransactionProduct) error {
	return s.Repository.DeleteTransactionProduct(transactionProduct)
}

func (s *Service) GetTransactionProducts(productID uint64) ([]*models.TransactionProduct, error) {
	return s.Repository.GetTransactionProducts(productID)
}

func (s *Service) SaveTransactionTicket(ticket models.Ticket) error {
	return s.Repository.SaveTransactionTicket(ticket)
}

func (s *Service) GetTransactionTicket(ticketID uint64) ([]*models.Ticket, error) {
	return s.Repository.GetTransactionTicket(ticketID)
}

// type transaksiService struct {
// 	TransaksiR repositories.TransaksiRepository
// }

// func NewTransaksiService(TransaksiR repositories.TransaksiRepository) TransaksiService {
// 	return &transaksiService{
// 		TransaksiR: TransaksiR,
// 	}
// }

// func (t *transaksiService) GetTransaksisService() ([]*models.Transaksi, error) {
// 	Transaksis, err := t.TransaksiR.GetTransaksisRepository()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return Transaksis, nil
// }

// func (t *transaksiService) GetTransaksiService(id string) (*models.Transaksi, error) {
// 	Transaksi, err := t.TransaksiR.GetTransaksiRepository(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return Transaksi, nil
// }

// func (t *transaksiService) CreateService(Transaksi models.Transaksi) (*models.Transaksi, error) {
// 	TransaksiR, err := t.TransaksiR.CreateRepository(Transaksi)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return TransaksiR, nil
// }

// func (t *transaksiService) UpdateService(id string, TransaksiBody models.Transaksi) (*models.Transaksi, error) {
// 	Transaksi, err := t.TransaksiR.UpdateRepository(id, TransaksiBody)
// 	if err != nil {
// 		return Transaksi, err
// 	}

// 	return Transaksi, nil
// }

// func (t *transaksiService) DeleteService(id string) error {
// 	err := t.TransaksiR.DeleteRepository(id)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (s *Service) SaveTransactionProduct(transactionProduct *TransactionProduct) error {
// 	return s.Repository.SaveTransactionProduct(transactionProduct)
// }
