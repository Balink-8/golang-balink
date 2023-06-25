package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"capstone/models"
)

var db *gorm.DB

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	config := Config{
		DB_Username: os.Getenv("DB_USERNAME"),
		DB_Password: os.Getenv("DB_PASSWORD"),
		DB_Port:     os.Getenv("DB_PORT"),
		DB_Host:     os.Getenv("DB_HOST"),
		DB_Name:     os.Getenv("DB_NAME"),
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	env := "dev"
	switch env {
	case "prod":
		db, err = gorm.Open(mysql.Open(os.Getenv("DSN")))
		if err != nil {
			panic(err)
		}
	case "dev":
		db, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
		if err != nil {
			panic(err)
		}
	}

	InitialMigration(db)

	seeder()

	return db
}

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Produk{})
	db.AutoMigrate(&models.Keranjang{})
	db.AutoMigrate(&models.Artikel{})
	db.AutoMigrate(&models.Event{})
	db.AutoMigrate(&models.KategoriProduk{})
	db.AutoMigrate(&models.Promo{})
	db.AutoMigrate(&models.ProfilePerusahaan{})
	db.AutoMigrate(&models.MetodePembayaran{})
	db.AutoMigrate(&models.PembayaranProduk{})
	db.AutoMigrate(&models.PembayaranEvent{})

}

func seeder() {
	MetodePembayaran := []models.MetodePembayaran{
		{Nama: "BCA", VA: 135353546},
	}
	for _, v := range MetodePembayaran {
		var exist models.MetodePembayaran
		err := db.Where("nama = ?", v.Nama).First(&exist).Error
		if err != nil {
			db.Create(&v)
		}
	}
}
