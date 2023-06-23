package services

import (
	"capstone/models"
	"capstone/repositories"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	userRMock = &repositories.IuserRepositoryMock{Mock: mock.Mock{}}
	userSMock = NewUserService(userRMock)
)

func TestGetUserService_Success(t *testing.T) {
	user := models.User{
		Nama:     "Sean 1",
		Email:    "sean1@gmail.com",
		Password: "123456",
		Alamat: "Jln. 123",
	}

	userRMock.Mock.On("GetUserRepository", "1").Return(user, nil)
	users, err := userSMock.GetUserService("1")

	assert.Nil(t, err)
	assert.NotNil(t, users)

	assert.Equal(t, user.Nama, users.Nama)
	assert.Equal(t, user.Password, users.Password)
	assert.Equal(t, user.Email, users.Email)
	assert.Equal(t, user.Alamat, users.Alamat)
}

func TestGetUserService_Failure(t *testing.T) {
	userRMock.Mock.On("GetUserRepository", "3").Return(nil, fmt.Errorf("user not found"))
	user, err := userSMock.GetUserService("3")

	assert.NotNil(t, err)
	assert.Nil(t, user)
}

func TestCreateUserService_Success(t *testing.T) {
	user := models.User{
		Nama:     "Sean 1",
		Email:    "sean1@gmail.com",
		Password: "123456",
		Alamat: "Jln. 123",
	}

	userRMock.Mock.On("CreateRepository", user).Return(user, nil)
	users, err := userSMock.CreateService(user)

	assert.Nil(t, err)
	assert.NotNil(t, users)

	assert.Equal(t, user.Nama, users.Nama)
	assert.Equal(t, user.Password, users.Password)
	assert.Equal(t, user.Email, users.Email)
	assert.Equal(t, user.Alamat, users.Alamat)
}

func TestUpdateUserService_Success(t *testing.T) {
	user := models.User{
		Model: gorm.Model{
			ID: 1,
		},
		Nama:     "Sean 1",
		Email:    "sean1@gmail.com",
		Password: "123456",
		Alamat: "Jln. 123",
	}

	userRMock.Mock.On("UpdateRepository", "1", user).Return(user, nil)
	users, err := userSMock.UpdateService("1", user)

	assert.Nil(t, err)
	assert.NotNil(t, users)

	assert.Equal(t, uint(1), users.ID)
	assert.Equal(t, user.Nama, users.Nama)
	assert.Equal(t, user.Password, users.Password)
	assert.Equal(t, user.Email, users.Email)
	assert.Equal(t, user.Alamat, users.Alamat)
}

func TestUpdateUserService_Failure(t *testing.T) {
	user := models.User{
		Model: gorm.Model{
			ID: 2,
		},
		Nama:     "Sean 1",
		Email:    "sean1@gmail.com",
		Password: "123456",
		Alamat: "Jln. 123",
	}

	userRMock.Mock.On("UpdateRepository", "2", user).Return(nil, fmt.Errorf("user not found"))
	users, err := userSMock.UpdateService("2", user)

	assert.Nil(t, users)
	assert.NotNil(t, err)
}

func TestDeleteUserService_Success(t *testing.T) {
	userRMock.Mock.On("DeleteRepository", "1").Return(nil)
	err := userSMock.DeleteService("1")

	assert.Nil(t, err)
}

func TestDeleteUserService_Failure(t *testing.T) {
	userRMock.Mock.On("DeleteRepository", "2").Return(fmt.Errorf("user not found"))
	err := userSMock.DeleteService("2")

	assert.NotNil(t, err)
}
