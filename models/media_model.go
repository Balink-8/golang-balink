package models

import "mime/multipart"

type GambarProduk struct {
	GambarId string `json:"gambar_id" validate:"required"`
	ProdukId string `json:"produk_id" validate:"required"`
	Gambar multipart.File `json:"file,omitempty" validate:"required"`
}