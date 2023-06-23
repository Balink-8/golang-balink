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

	profilePerusahaanR = r.NewProfilePerusahaanRepository(DB)
	profilePerusahaanS = s.NewProfilePerusahaanService(profilePerusahaanR)
	profilePerusahaanC = c.NewProfilePerusahaanController(profilePerusahaanS, JWT)

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

	dashboardR = r.NewDashboardRepository(DB)
	dashboardS = s.NewDashboardService(dashboardR)
	dashboardC = c.NewDashboardController(dashboardS)

	dashboardMobileR = r.NewDashboardMobileRepository(DB)
	dashboardMobileS = s.NewDashboardMobileServices(dashboardMobileR)
	dashboardMobileC = c.NewDashboardMobileController(dashboardMobileS)
)

func New() *echo.Echo {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	e := echo.New()

	m.LoggerMiddleware(e)

	e.Use(middleware.CORS())

	e.GET("/dashboard-mobile", dashboardMobileC.DashboardMobileGetAll)

	auth := e.Group("")
	auth.Use(middleware.JWT([]byte(os.Getenv("JWT_KEY"))))
	auth.GET("/user", userC.GetUsersController)
	auth.GET("/user/:id", userC.GetUserController)
	e.POST("/user_register", userC.CreateController)
	auth.DELETE("/user/:id", userC.DeleteController)
	auth.PUT("/user/:id", userC.UpdateController)
	e.POST("/user_login", userC.LoginController)
	e.GET("/logout", userC.LogoutController)

	auth.GET("/admin", profilePerusahaanC.GetProfilePerusahaanController)
	auth.PUT("/admin", profilePerusahaanC.UpdateController)
	e.POST("/admin_login", profilePerusahaanC.LoginController)

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

	auth.GET("/user/:id_user/keranjang", keranjangC.GetKeranjangByUserController)

	auth.GET("/web_dashboard", dashboardC.GetDashboardController)

	return e
}
