package services

import (
	"capstone/models"
	"capstone/repositories"
)

type ProfilePerusahaanService interface {
	GetProfilePerusahaanService() (*models.ProfilePerusahaan, error)
	UpdateService(ProfilePerusahaanBody models.ProfilePerusahaan) (*models.ProfilePerusahaan, error)
	LoginService(login models.ProfilePerusahaan) (*models.ProfilePerusahaan, error)
}

type profilePerusahaanService struct {
	ProfilePerusahaanR repositories.ProfilePerusahaanRepository
}

func NewProfilePerusahaanService(ProfilePerusahaanR repositories.ProfilePerusahaanRepository) ProfilePerusahaanService {
	return &profilePerusahaanService{
		ProfilePerusahaanR: ProfilePerusahaanR,
	}
}

func (p *profilePerusahaanService) GetProfilePerusahaanService() (*models.ProfilePerusahaan, error) {
	ProfilePerusahaan, err := p.ProfilePerusahaanR.GetProfilePerusahaanRepository()
	if err != nil {
		return nil, err
	}

	return ProfilePerusahaan, nil
}

func (p *profilePerusahaanService) UpdateService(ProfilePerusahaanBody models.ProfilePerusahaan) (*models.ProfilePerusahaan, error) {
	ProfilePerusahaan, err := p.ProfilePerusahaanR.UpdateRepository(ProfilePerusahaanBody)
	if err != nil {
		return nil, err
	}

	return ProfilePerusahaan, nil
}

func (p *profilePerusahaanService) LoginService(login models.ProfilePerusahaan) (*models.ProfilePerusahaan, error) {
	loginR, err := p.ProfilePerusahaanR.LoginRepository(login)
	if err != nil {
		return nil, err
	}

	return loginR, nil
}