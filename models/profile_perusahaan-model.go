package models

import "gorm.io/gorm"

type ProfilePerusahaan struct {
	gorm.Model
	Nama              string `json:"nama" form:"nama"`
	Deskripsi         string `json:"deskripsi" form:"deskripsi"`
	Foto_Profile      string `json:"foto_profile" form:"foto_profile"`
	Email             string `json:"email" form:"email"`
	Password          string `json:"password" form:"password"`
	No_Telepon        string `json:"no_telepon" form:"no_telepon"`
	WhatsApp          string `json:"whatsapp" form:"whatsapp"`
	Instagram         string `json:"instagram" form:"instagram"`
	Facebook          string `json:"facebook" form:"facebook"`
	Alamat            string `json:"alamat" form:"alamat"`
	Negara            string `json:"negara" form:"negara"`
	Kode_Pos          string `json:"kode_pos" form:"kode_pos"`
	Rekening_BCA      string `json:"rekening_bca" form:"rekening_bca"`
	Rekening_Mandiri  string `json:"rekening_mandiri" form:"rekening_mandiri"`
	Rekening_BRI      string `json:"rekening_bri" form:"rekening_bri"`
	Rekening_BNI      string `json:"rekening_bni" form:"rekening_bni"`
	Rekening_BTN      string `json:"rekening_btn" form:"rekening_btn"`
	Rekening_Seabank  string `json:"rekening_seabank" form:"rekening_seabank"`
	Rekening_BPD_Bali string `json:"rekening_bpd_bali" form:"rekening_bpd_bali"`
	Image             string `json:"image" from:"image"`
}

type CreateProfilePerusahaan struct {
	ProfilePerusahaan *ProfilePerusahaan `json:"profile_perusahaan" form:"profile_perusahaan"`
	Token             string             `json:"token" form:"token"`
}
