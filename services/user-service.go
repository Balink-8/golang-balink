package services

import (
	"capstone/models"
	"capstone/repositories"
)

type UserService interface {
	GetUsersService(page int, limit int, order string) ([]*models.User, int, error)
	GetUserService(id string) (*models.User, error)
	GetAdminService() (*models.User, error)
	CreateService(user models.User) (*models.User, error)
	UpdateUserService(id string, userBody models.User) (*models.User, error)
	UpdateAdminService(userBody models.User) (*models.User, error)
	DeleteService(id string) error
	LoginService(login models.User) (*models.User, error)
}

type userService struct {
	UserR repositories.UserRepository
}

func NewUserService(UserR repositories.UserRepository) UserService {
	return &userService{
		UserR: UserR,
	}
}

func (u *userService) GetUsersService(page int, limit int, order string) ([]*models.User, int, error) {
	Users, totalData, err := u.UserR.GetUsersRepository(page, limit, order)
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

func (u *userService) GetAdminService() (*models.User, error) {
	user, err := u.UserR.GetAdminRepository()
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

func (u *userService) UpdateUserService(id string, userBody models.User) (*models.User, error) {
	user, err := u.UserR.UpdateUserRepository(id, userBody)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userService) UpdateAdminService(userBody models.User) (*models.User, error) {
	user, err := u.UserR.UpdateAdminRepository(userBody)
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