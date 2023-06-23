package services

import (
	"capstone/models"
	"capstone/repositories"
	"errors"
)

type UserService interface {
	GetUsersService(page int, limit int, order string, search string) ([]*models.User, int, error)
	GetUserService(id string) (*models.User, error)
	CreateService(user models.User) (*models.User, error)
	UpdateService(id string, userBody models.User) (*models.User, error)
	DeleteService(id string) error
	LoginService(login models.User) (*models.User, error)
	ForgotPasswordService(req *models.ForgotPassword) (*models.User, error)
}

type userService struct {
	UserR repositories.UserRepository
}

func NewUserService(UserR repositories.UserRepository) UserService {
	return &userService{
		UserR: UserR,
	}
}

func (u *userService) GetUsersService(page int, limit int, order string, search string) ([]*models.User, int, error) {
	Users, totalData, err := u.UserR.GetUsersRepository(page, limit, order, search)
	if err != nil {
		return nil, 0, err
	}

	return Users, totalData, nil
}

func (u *userService) GetUserService(id string) (*models.User, error) {
	user, err := u.UserR.GetUserRepository(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userService) CreateService(user models.User) (*models.User, error) {
	UserR, err := u.UserR.CreateRepository(user)
	if err != nil {
		return nil, err
	}

	return UserR, nil
}

func (u *userService) UpdateService(id string, userBody models.User) (*models.User, error) {
	user, err := u.UserR.UpdateRepository(id, userBody)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userService) DeleteService(id string) error {
	err := u.UserR.DeleteRepository(id)
	if err != nil {
		return err
	}

	return nil
}

func (u *userService) LoginService(login models.User) (*models.User, error) {
	loginR, err := u.UserR.LoginRepository(login)
	if err != nil {
		return nil, err
	}

	return loginR, nil
}

func (u *userService) ForgotPasswordService(req *models.ForgotPassword) (*models.User, error) {
	user, err := u.UserR.CekEmailRepository(req.Email)
	if err != nil {
		return nil, err
	}

	if req.Password != req.Confirm_Password {
		return nil, errors.New("Password Tidak Cocok")
	}
	user.Password = req.Password
	err = u.UserR.UpdatesRepository(*user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
