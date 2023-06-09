package routes

import (
	"github.com/Caknoooo/golang-clean_template/controller"
	"github.com/Caknoooo/golang-clean_template/middleware"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/gin-gonic/gin"
)

func Router(route *gin.Engine, userController controller.UserController, jwtService services.JWTService, pasienController controller.PasienController, dokterController controller.DokterController, transaksiController controller.TransaksiController, ruanganController controller.RuanganController, perawatController controller.PerawatController, obatController controller.ObatController, sesidokterController controller.SesiDokterController) {
	// routes := route.Group("/api/user")
	// {
	// 	routes.POST("", userController.RegisterUser)
	// 	routes.GET("", middleware.Authenticate(jwtService), userController.GetAllUser)
	// 	routes.POST("/login", userController.LoginUser)
	// 	routes.DELETE("/", middleware.Authenticate(jwtService), userController.DeleteUser)
	// 	routes.PUT("/", middleware.Authenticate(jwtService), userController.UpdateUser)
	// 	routes.GET("/me", middleware.Authenticate(jwtService), userController.MeUser)
	// }

	pasienRoutes := route.Group("/api/pasien")
	{
		pasienRoutes.POST("/register", pasienController.RegisterPasien)
		pasienRoutes.POST("/login", pasienController.LoginPasien)
		pasienRoutes.DELETE("/", pasienController.DeletePasien)
		pasienRoutes.PUT("/", middleware.Authenticate(jwtService), pasienController.UpdatePasien)
		pasienRoutes.GET("/me", middleware.Authenticate(jwtService), pasienController.MePasien)
		pasienRoutes.GET("/PembelianObat/:id", middleware.Authenticate(jwtService), pasienController.GetLatestPembelianObat)
		pasienRoutes.GET("/:nik", pasienController.GetLatestReservation)
		// Add other pasien routes here if needed
	}

	adminRoutes := route.Group("/api/admin")
	adminRoutes.POST("/login", userController.LoginUser)
	adminRoutes.POST("", userController.RegisterUser)
	adminRoutes.Use(middleware.Authenticate(jwtService))
	
	{
		adminRoutes.GET("", middleware.Authenticate(jwtService), userController.GetAllUser)
		adminRoutes.DELETE("/", middleware.Authenticate(jwtService), userController.DeleteUser)
		adminRoutes.PUT("/", middleware.Authenticate(jwtService), userController.UpdateUser)
		adminRoutes.GET("/me", middleware.Authenticate(jwtService), userController.MeUser)

		pasienAdminRoutes := adminRoutes.Group("/pasien")
		{
			pasienAdminRoutes.GET("", middleware.Authenticate(jwtService), pasienController.GetAllPasien)
		}

		dokterRoutes := adminRoutes.Group("/dokter")
		{
			dokterRoutes.POST("/new", dokterController.RegisterDokter)
			dokterRoutes.GET("", dokterController.GetAllDokter)
			dokterRoutes.GET("/:id", dokterController.GetDokterByID)
			dokterRoutes.PUT("/edit", dokterController.UpdateDoctor)
		}

		transaksiRoutes := adminRoutes.Group("/transaksi")
		{
			transaksiRoutes.POST("/new", transaksiController.NewTransaksi)
			transaksiRoutes.GET("/:id", transaksiController.GetTransaksiByIDPasien)
			transaksiRoutes.GET("/all", transaksiController.GetAllTransaksi)
		}

		ruanganRoutes := adminRoutes.Group("/ruangan")
		{

			ruanganRoutes.GET("", ruanganController.GetAllRuangan)
		}

		perawatRoutes := adminRoutes.Group("/perawat")
		{
			perawatRoutes.POST("/new", perawatController.RegisterPerawat)
			perawatRoutes.GET("", perawatController.GetAllPerawat)
			perawatRoutes.GET("/:id", perawatController.GetPerawatByID)
			perawatRoutes.GET("/jadwal", perawatController.GetJadwalPerawat)
		}

		obatRoutes := adminRoutes.Group("/obat")
		{
			obatRoutes.POST("/new", obatController.RegisterObat)
			obatRoutes.GET("", obatController.GetAllObat)
			obatRoutes.GET("/:id", obatController.GetObatByID)
		}

		sesiDokterRoutes := adminRoutes.Group("/sesi_dokter")
		{
			sesiDokterRoutes.POST("/new", sesidokterController.RegisterSesiDokter)
			sesiDokterRoutes.GET("", sesidokterController.GetAllSesiDokter)
			sesiDokterRoutes.GET("/:id", sesidokterController.GetSesiDokterByID)
		}
	}

}
