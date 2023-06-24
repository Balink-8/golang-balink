package repositories

import (
	"capstone/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsersRepository(page int, limit int, order string, search string) ([]*models.User, int, error)
	GetUserRepository(id string) (*models.User, error)
	CreateRepository(user models.User) (models.User, error)
	UpdateRepository(id string, userBody models.User) (*models.User, error)
	DeleteRepository(id string) error
	LoginRepository(login models.User) (*models.User, error)
	CekEmailRepository(email string) (*models.User, error)
	UpdatesRepository(user models.User) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &userRepository{
		DB: DB,
	}
}

func (u *userRepository) GetUsersRepository(page int, limit int, order string, search string) ([]*models.User, int, error) {
	var Users []*models.User
	var totalData int64

	query := u.DB.Model(&models.User{})

	if search != "" {
		query = query.Where("nama LIKE ?", "%"+search+"%")
	}

	if err := query.Count(&totalData).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	query = query.Offset(offset).Limit(limit)

	switch order {
	case "asc":
		query = query.Order("ID ASC")
	case "desc":
		query = query.Order("ID DESC")
	}

	if err := query.Find(&Users).Error; err != nil {
		return nil, 0, err
	}

	return Users, int(totalData), nil
}

func (u *userRepository) GetUserRepository(id string) (*models.User, error) {
	var user models.User

	if err := u.DB.Where("ID = ?", id).Take(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) CreateRepository(User models.User) (models.User, error) {
	if err := u.DB.Save(&User).Error; err != nil {
		return models.User{}, err
	}

	return User, nil
}

func (u *userRepository) UpdateRepository(id string, userBody models.User) (*models.User, error) {
	user, err := u.GetUserRepository(id)
	if err != nil {
		return nil, err
	}

	err = u.DB.Where("ID = ?", id).Updates(models.User{Nama: userBody.Nama, Foto_Profile: userBody.Foto_Profile, Email: userBody.Email, Password: userBody.Password, No_Telepon: userBody.No_Telepon, Alamat: userBody.Alamat}).Error
	if err != nil {
		return nil, err
	}

	user.Nama = userBody.Nama
	user.Foto_Profile = userBody.Foto_Profile
	user.Email = userBody.Email
	user.Password = userBody.Password
	user.No_Telepon = userBody.No_Telepon
	user.Alamat = userBody.Alamat

	return user, nil
}

func (u *userRepository) DeleteRepository(id string) error {
	_, err := u.GetUserRepository(id)
	if err != nil {
		return err
	}

	if err := u.DB.Delete(&models.User{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (u *userRepository) LoginRepository(login models.User) (*models.User, error) {
	if err := u.DB.Where("email = ? AND password = ?", login.Email, login.Password).First(&login).Error; err != nil {
		return nil, err
	}

	return &login, nil
}

func (u *userRepository) CekEmailRepository(email string) (*models.User, error) {
	var user = models.User{}
	if err := u.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) UpdatesRepository(user models.User) error {

	if err := u.DB.Updates(&user).Error; err != nil {
		return err
	}

	return nil
}
