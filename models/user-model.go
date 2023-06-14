package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nama         string `json:"nama" form:"nama"`
	Foto_Profile string `json:"foto_profile" form:"foto_profile"`
	Email        string `json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
	No_Telepon   string `json:"no_telepon" form:"no_telepon"`
	Alamat       string `json:"alamat" form:"alamat"`
	Role         string `json:"role" form:"role" gorm:"type:enum('User');default:'User'; not-null"`
}

type CreateUser struct {
	User  *User  `json:"user" form:"user"`
	Token string `json:"token" form:"token"`
}

type RegisterUser struct {
	Email      string `json:"email" form:"email"`
	Password   string `json:"password" form:"password"`
	No_Telepon string `json:"no_telepon" form:"no_telepon"`
}

type RegisterRespon struct {
	ID           uint   `json:"id" form:"id"`
	Nama         string `json:"nama" form:"nama"`
	Foto_Profile string `json:"foto_profile" form:"foto_profile"`
	Email        string `json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
	No_Telepon   string `json:"no_telepon" form:"no_telepon"`
	Alamat       string `json:"alamat" form:"alamat"`
	Role         string `json:"role" form:"role" gorm:"type:enum('User');default:'User'; not-null"`
}
