package services

import (
    "golang-img-api/helper"
    "golang-img-api/models"
    "github.com/go-playground/validator/v10"
)

var (
    validate = validator.New()
)

type mediaUpload interface {
    FileUpload(gambar models.GambarProduk) (string, error)
}

type media struct {}

func NewMediaUpload() mediaUpload {
    return &media{}
}

func (*media) FileUpload(gambar models.GambarProduk) (string, error) {
    // Validate
    err := validate.Struct(gambar)
    if err != nil {
        return "", err
    }

    // Upload
    uploadUrl, err := helper.ImageUploadHelper(gambar.Gambar)
    if err != nil {
        return "", err
    }
    return uploadUrl, nil
}