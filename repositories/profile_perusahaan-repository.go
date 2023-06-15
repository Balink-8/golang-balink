package repositories

import (
	"capstone/models"

	"gorm.io/gorm"
)

type ProfilePerusahaanRepository interface {
	GetProfilePerusahaanRepository() (*models.ProfilePerusahaan, error)
	UpdateRepository(ProfilePerusahaanBody models.ProfilePerusahaan) (*models.ProfilePerusahaan, error)
	LoginRepository(login models.ProfilePerusahaan) (*models.ProfilePerusahaan, error)
}

type profilePerusahaanRepository struct {
	DB *gorm.DB
}

func NewProfilePerusahaanRepository(DB *gorm.DB) ProfilePerusahaanRepository {
	return &profilePerusahaanRepository{
		DB: DB,
	}
}

func (p *profilePerusahaanRepository) GetProfilePerusahaanRepository() (*models.ProfilePerusahaan, error) {
	var ProfilePerusahaan models.ProfilePerusahaan

	if err := p.DB.First(&ProfilePerusahaan).Error; err != nil {
		return nil, err
	}

	return &ProfilePerusahaan, nil
}

func (u *profilePerusahaanRepository) UpdateRepository(profilePerusahaanBody models.ProfilePerusahaan) (*models.ProfilePerusahaan, error) {
	profilePerusahaan, err := u.GetProfilePerusahaanRepository()
	if err != nil {
		return nil, err
	}

	err = u.DB.Model(&profilePerusahaan).Updates(models.ProfilePerusahaan{
		Nama : profilePerusahaanBody.Nama,
		Deskripsi : profilePerusahaanBody.Deskripsi,
		Foto_Profile : profilePerusahaanBody.Foto_Profile,
		Email : profilePerusahaanBody.Email,
		Password : profilePerusahaanBody.Password,
		No_Telepon : profilePerusahaanBody.No_Telepon,
		WhatsApp : profilePerusahaanBody.WhatsApp,
		Instagram: profilePerusahaanBody.Instagram,
		Facebook: profilePerusahaanBody.Facebook,
		Alamat : profilePerusahaanBody.Alamat,
		Negara : profilePerusahaanBody.Negara,
		Kode_Pos : profilePerusahaanBody.Kode_Pos,
		Rekening_BCA : profilePerusahaanBody.Rekening_BCA,
		Rekening_BNI : profilePerusahaanBody.Rekening_BNI,
		Rekening_BPD_Bali : profilePerusahaanBody.Rekening_BPD_Bali,
		Rekening_BRI : profilePerusahaanBody.Rekening_BRI,
		Rekening_BTN : profilePerusahaanBody.Rekening_BTN,
		Rekening_Mandiri : profilePerusahaanBody.Rekening_Mandiri,
		Rekening_Seabank : profilePerusahaanBody.Rekening_Seabank,
	}).Error
	if err != nil {
		return nil, err
	}

	return profilePerusahaan, nil
}

func (p *profilePerusahaanRepository) LoginRepository(login models.ProfilePerusahaan) (*models.ProfilePerusahaan, error) {
	if err := p.DB.Where("email = ? AND password = ?", login.Email, login.Password).First(&login).Error; err != nil {
		return nil, err
	}

	return &login, nil
}