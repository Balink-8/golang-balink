package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	h "capstone/helpers"
	m "capstone/middlewares"
	"capstone/models"
	"capstone/services"
)

type ProfilePerusahaanController interface {
	GetProfilePerusahaanController(c echo.Context) error
	UpdateController(c echo.Context) error
	LoginController(c echo.Context) error
}

type profilePerusahaanController struct {
	ProfilePerusahaanS services.ProfilePerusahaanService
	jwt m.JWTS
}

func NewProfilePerusahaanController(ProfilePerusahaanS services.ProfilePerusahaanService, jwtS m.JWTS) ProfilePerusahaanController {
	return &profilePerusahaanController{
		ProfilePerusahaanS: ProfilePerusahaanS,
		jwt: jwtS,
	}
}

func (p *profilePerusahaanController) GetProfilePerusahaanController(c echo.Context) error {
	ProfilePerusahaan, err := p.ProfilePerusahaanS.GetProfilePerusahaanService()
	if err != nil {
		return h.Response(c, http.StatusNotFound, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    ProfilePerusahaan,
		Message: "Get Profile Perusahaan success",
		Status:  true,
	})
}

func (p *profilePerusahaanController) UpdateController(c echo.Context) error {
	var profilePerusahaanBody models.ProfilePerusahaan
	if err := c.Bind(&profilePerusahaanBody); err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	profilePerusahaan, err := p.ProfilePerusahaanS.UpdateService(profilePerusahaanBody)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    profilePerusahaan,
		Message: "Update Profile Perusahaan success",
		Status:  true,
	})
}

func (p *profilePerusahaanController) LoginController(c echo.Context) error {
	var ProfilePerusahaan models.CreateProfilePerusahaan

	err := c.Bind(&ProfilePerusahaan.ProfilePerusahaan)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	ProfilePerusahaan.ProfilePerusahaan, err = p.ProfilePerusahaanS.LoginService(*ProfilePerusahaan.ProfilePerusahaan)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	token, err := p.jwt.CreateJWTToken(ProfilePerusahaan.ProfilePerusahaan.ID, ProfilePerusahaan.ProfilePerusahaan.Nama)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	ProfilePerusahaan.Token = token
	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    ProfilePerusahaan,
		Message: "Login success",
		Status:  true,
	})
}