package controllers

import (
	"fmt"
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
	var User models.User

	fmt.Println("Data :", &User)

	err := c.Bind(&User)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	file, err := c.FormFile("gambar") // Mengubah ctx menjadi c pada bagian ini

	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: "Gambar tidak boleh kosong", // Mengubah pesan error menjadi string statis
			Status:  false,
		})
	}

	src, err := file.Open()
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: "Gagal membuka file", // Mengubah pesan error menjadi string statis
			Status:  false,
		})
	}

	re := regexp.MustCompile(`.png|.jpeg|.jpg`)

	if !re.MatchString(file.Filename) {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: "Format file yang disediakan tidak diperbolehkan. Unggah gambar JPEG atau PNG", // Mengubah pesan error menjadi string statis
			Status:  false,
		})
	}

	uploadUrl, err := services.NewMediaUpload().FileUpload(models.File{File: src})
	if err != nil {
		return h.Response(c, http.StatusInternalServerError, h.ResponseModel{
			Data:    nil,
			Message: "Terjadi kesalahan saat mengunggah foto", // Mengubah pesan error menjadi string statis
			Status:  false,
		})
	}
	User.Gambar = uploadUrl // Mengubah artikelInput menjadi Produk

	User, err = u.UserS.CreateService(User)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    User,
		Message: "Create User success",
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
