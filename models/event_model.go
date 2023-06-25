package models

import (
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Artikel_ID      string `json:"artikel_id" form:"artikel_id"`
	Nama            string `json:"nama" form:"nama"`
	Deskripsi       string `json:"deskripsi" form:"deskripsi"`
	Harga_Tiket     int64  `json:"harga_tiket" form:"harga_tiket"`
	Stok_Tiket      int64  `json:"stok_tiket" form:"stok_tiket"`
	Waktu_Mulai     string `json:"waktu_mulai" form:"waktu_mulai"`
	Waktu_Selesai   string `json:"waktu_selesai" form:"waktu_selesai"`
	Tanggal_Mulai   string `json:"tanggal_mulai" form:"tanggal_mulai"`
	Tanggal_Selesai string `json:"tanggal_selesai" form:"tanggal_selesai"`
	Lokasi          string `json:"lokasi" form:"lokasi"`
	Link_Lokasi     string `json:"link_lokasi" form:"link_lokasi"`
	Gambar          string `json:"gambar" from:"gambar"`
}
