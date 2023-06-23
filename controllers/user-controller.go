package controllers

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/labstack/echo/v4"

	h "capstone/helpers"
	m "capstone/middlewares"
	"capstone/models"
	"capstone/services"
)

type UserController interface {
	GetUsersController(c echo.Context) error
	GetUserController(c echo.Context) error
	CreateController(c echo.Context) error
	UpdateController(c echo.Context) error
	DeleteController(c echo.Context) error
	LoginController(c echo.Context) error
	ForgotPasswordController(c echo.Context) error
}

type userController struct {
	UserS services.UserService
	jwt   m.JWTS
}

func NewUserController(UserS services.UserService, jwtS m.JWTS) UserController {
	return &userController{
		UserS: UserS,
		jwt:   jwtS,
	}
}

func (u *userController) GetUsersController(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	order := c.QueryParam("order")
	search := c.QueryParam("search")

	Users, totalData, err := u.UserS.GetUsersService(page, limit, order, search)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	responseData := map[string]interface{}{
		"data":       Users,
		"page":       page,
		"data_shown": len(Users),
		"total_data": totalData,
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    responseData,
		Message: "Get all User success",
		Status:  true,
	})
}

func (u *userController) GetUserController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	var user *models.User

	user, err = u.UserS.GetUserService(id)
	if err != nil {
		return h.Response(c, http.StatusNotFound, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    user,
		Message: "Get user success",
		Status:  true,
	})
}

func (u *userController) CreateController(c echo.Context) error {
	var user models.CreateUser

	err := c.Bind(&user.User)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	file, err := c.FormFile("image")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Image cannot be empty", err)
	}

	src, err := file.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to open file", err)
	}

	re := regexp.MustCompile(`.png|.jpeg|.jpg`)
	if !re.MatchString(file.Filename) {
		return echo.NewHTTPError(http.StatusBadRequest, "The provided file format is not allowed. Please upload a JPEG or PNG image")
	}

	uploadURL, err := services.NewMediaUpload().FileUpload(models.File{File: src})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error uploading photo", err)
	}
	user.User.Image = uploadURL

	createdUser, err := u.UserS.CreateService(*user.User)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	token, err := u.jwt.CreateJWTToken(createdUser.ID, createdUser.Nama)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user.Token = token
	return c.JSON(http.StatusOK, h.ResponseModel{
		Data:    user,
		Message: "Create user success",
		Status:  true,
	})
}

func (u *userController) UpdateController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	var user *models.User

	err = c.Bind(&user)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	user, err = u.UserS.UpdateService(id, *user)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    user,
		Message: "Update user success",
		Status:  true,
	})
}

func (u *userController) DeleteController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	err = u.UserS.DeleteService(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    nil,
		Message: "Delete user success",
		Status:  true,
	})
}

func (u *userController) LoginController(c echo.Context) error {
	var user models.CreateUser

	err := c.Bind(&user.User)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	user.User, err = u.UserS.LoginService(*user.User)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	token, err := u.jwt.CreateJWTToken(user.User.ID, user.User.Nama)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	user.Token = token
	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    user,
		Message: "Login success",
		Status:  true,
	})
}

func (u *userController) ForgotPasswordController(c echo.Context) error {

	payload := models.ForgotPassword{}

	c.Bind(&payload)

	user, err := u.UserS.ForgotPasswordService(&payload)
	if err != nil {
		return h.Response(c, http.StatusNotFound, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    user,
		Message: "Update Password success",
		Status:  true,
	})
}
