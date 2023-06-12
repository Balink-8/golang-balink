package routes

import (
	"capstone/configs"
	c "capstone/controllers"
	m "capstone/middlewares"
	r "capstone/repositories"
	s "capstone/services"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	DB = configs.InitDB()

	JWT = m.NewJWTS()

	userR = r.NewUserRepository(DB)
	userS = s.NewUserService(userR)
	userC = c.NewUserController(userS, JWT)

	produkR = r.NewProdukRepository(DB)
	produkS = s.NewProdukService(produkR)
	produkC = c.NewProdukController(produkS)

	keranjangR = r.NewKeranjangRepository(DB)
	keranjangS = s.NewKeranjangService(keranjangR)
	keranjangC = c.NewKeranjangController(keranjangS)

	artikelR = r.NewArtikelRepository(DB)
	artikelS = s.NewArtikelService(artikelR)
	artikelC = c.NewArtikelController(artikelS)

	eventR = r.NewEventRepository(DB)
	eventS = s.NewEventService(eventR)
	eventC = c.NewEventController(eventS)

	kategoriProdukR = r.NewKategoriProdukRepository(DB)
	kategoriProdukS = s.NewKategoriProdukService(kategoriProdukR)
	kategoriProdukC = c.NewKategoriProdukController(kategoriProdukS)

	promoR = r.NewPromoRepository(DB)
	promoS = s.NewPromoService(promoR)
	promoC = c.NewPromoController(promoS)

	masalahR = r.NewMasalahRepository(DB)
	masalahS = s.NewMasalahService(masalahR)
	masalahC = c.NewMasalahController(masalahS)

)

func New() *echo.Echo {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	e := echo.New()

	m.LoggerMiddleware(e)

	auth := e.Group("")
	auth.Use(middleware.JWT([]byte(os.Getenv("JWT_KEY"))))
	auth.GET("/user", userC.GetUsersController)
	auth.GET("/user/:id", userC.GetUserController)
	e.POST("/user", userC.CreateController)
	auth.DELETE("/user/:id", userC.DeleteController)
	auth.PUT("/user/:id", userC.UpdateUserController)

	auth.GET("/admin", userC.GetAdminController)
	auth.PUT("/admin", userC.UpdateAdminController)

	auth.GET("/produk", produkC.GetProduksController)
	auth.GET("/produk/:id", produkC.GetProdukController)
	auth.POST("/produk", produkC.CreateController)
	auth.DELETE("/produk/:id", produkC.DeleteController)
	auth.PUT("/produk/:id", produkC.UpdateController)

	auth.GET("/keranjang", keranjangC.GetKeranjangsController)
	auth.GET("/keranjang/:id", keranjangC.GetKeranjangController)
	auth.POST("/keranjang", keranjangC.CreateController)
	auth.DELETE("/keranjang/:id", keranjangC.DeleteController)
	auth.PUT("/keranjang/:id", keranjangC.UpdateController)

	auth.GET("/artikel", artikelC.GetArtikelsController)
	auth.GET("/artikel/:id", artikelC.GetArtikelController)
	auth.POST("/artikel", artikelC.CreateController)
	auth.DELETE("/artikel/:id", artikelC.DeleteController)
	auth.PUT("/artikel/:id", artikelC.UpdateController)

	auth.GET("/event", eventC.GetEventsController)
	auth.GET("/event/:id", eventC.GetEventController)
	auth.POST("/event", eventC.CreateController)
	auth.DELETE("/event/:id", eventC.DeleteController)
	auth.PUT("/event/:id", eventC.UpdateController)

	auth.GET("/kategori_produk", kategoriProdukC.GetKategoriProduksController)
	auth.GET("/kategori_produk/:id", kategoriProdukC.GetKategoriProdukController)
	auth.POST("/kategori_produk", kategoriProdukC.CreateController)
	auth.DELETE("/kategori_produk/:id", kategoriProdukC.DeleteController)
	auth.PUT("/kategori_produk/:id", kategoriProdukC.UpdateController)

	auth.GET("/promo", promoC.GetPromosController)
	auth.GET("/promo/:id", promoC.GetPromoController)
	auth.POST("/promo", promoC.CreateController)
	auth.DELETE("/promo/:id", promoC.DeleteController)
	auth.PUT("/promo/:id", promoC.UpdateController)

	auth.GET("/masalah", masalahC.GetMasalahsController)
	auth.GET("/masalah/:id", masalahC.GetMasalahController)
	auth.POST("/masalah", masalahC.CreateController)

	e.POST("/login", userC.LoginController)

	auth.GET("/user/:id_user/keranjang", keranjangC.GetKeranjangByUserController)

	return e
}
