package repositories

import (
	"capstone/models"
	"fmt"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock interface {
	GetUsersRepository(page int, limit int, order string, search string) ([]*models.User, int, error)
	GetUserRepository(id string) (*models.User, error)
	CreateRepository(userData models.User) (*models.User, error)
	UpdateRepository(id string, userBody models.User) (*models.User, error)
	DeleteRepository(id string) error
	LoginRepository(login models.User) (*models.User, error)
}

type IuserRepositoryMock struct {
	Mock mock.Mock
}

func NewUserRepositoryMock(mock mock.Mock) UserRepositoryMock {
	return &IuserRepositoryMock{
		Mock: mock,
	}
}

func (u *IuserRepositoryMock) GetUsersRepository(page int, limit int, order string, search string) ([]*models.User, int, error) {
	args := u.Mock.Called()
	if args.Get(0) == nil {
		return nil, 0, args.Get(1).(error)
	}

	users := args.Get(0).([]*models.User)

	return users, 0, nil
}

func (u *IuserRepositoryMock) GetUserRepository(id string) (*models.User, error) {
	args := u.Mock.Called(id)
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	user := args.Get(0).(models.User)
	return &user, nil
}

func (u *IuserRepositoryMock) CreateRepository(userData models.User) (*models.User, error) {
	args := u.Mock.Called(userData)
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	user := args.Get(0).(models.User)

	return &user, nil
}

func (u *IuserRepositoryMock) UpdateRepository(id string, userData models.User) (*models.User, error) {
	args := u.Mock.Called(id, userData)
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	user := args.Get(0).(models.User)

	return &user, nil
}

func (u *IuserRepositoryMock) DeleteRepository(id string) error {
	args := u.Mock.Called(id)
	if args.Get(0) != nil {
		return fmt.Errorf("must nil")
	}

	return nil
}

func (u *IuserRepositoryMock) LoginRepository(login models.User) (*models.User, error) {
	args := u.Mock.Called(login)
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	user := args.Get(0).(models.User)

	return &user, nil
}