package services

import (
	"capstone/models"
	"capstone/repositories"
)

type UserService interface {
	GetUsersService() ([]*models.User, error)
	GetUserService(id string) (*models.User, error)
	CreateService(user models.User) (res models.RegisterRespon, err error)
	UpdateService(id string, userBody models.User) (*models.User, error)
	DeleteService(id string) error
	LoginService(login models.User) (*models.User, error)
	RegisterService(register models.User) (*models.User, error)
}

type userService struct {
	userR repositories.UserRepository
}

func NewUserService(userR repositories.UserRepository) UserService {
	return &userService{
		userR: userR,
	}
}

func (u *userService) GetUsersService() ([]*models.User, error) {
	users, err := u.userR.GetUsersRepository()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userService) GetUserService(id string) (*models.User, error) {
	user, err := u.userR.GetUserRepository(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userService) GetUserByEmailService(email string) (*models.User, error) {
	user, err := u.userR.GetUserByEmailRepository(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userService) CreateService(user models.User) (res models.RegisterRespon, err error) {
	userR, err := u.userR.CreateRepository(user)
	if err != nil {
		return res, err
	}

	RegisterRespon := models.RegisterRespon{
		ID:           userR.ID,
		Nama:         userR.Nama,
		Foto_Profile: userR.Foto_Profile,
		Email:        userR.Email,
		Password:     userR.Password,
		No_Telepon:   userR.No_Telepon,
		Alamat:       userR.Alamat,
	}

	return RegisterRespon, nil
}

func (u *userService) UpdateService(id string, userBody models.User) (*models.User, error) {
	user, err := u.userR.UpdateRepository(id, userBody)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userService) DeleteService(id string) error {
	err := u.userR.DeleteRepository(id)
	if err != nil {
		return err
	}

	return nil
}

func (u *userService) LoginService(login models.User) (*models.User, error) {
	loginR, err := u.userR.LoginRepository(login)
	if err != nil {
		return nil, err
	}

	return loginR, nil
}

func (u *userService) RegisterService(register models.User) (*models.User, error) {
	_, err := u.userR.GetUserByEmailRepository(register.Email)
	if err != nil {
		reg, _ := u.userR.CreateRepository(register)
		return reg, nil
	}

	return nil, err
}
