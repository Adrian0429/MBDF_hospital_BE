package routes

import (
	"github.com/Caknoooo/golang-clean_template/controller"
	"github.com/Caknoooo/golang-clean_template/middleware"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/gin-gonic/gin"
)

func Router(route *gin.Engine, userController controller.UserController, jwtService services.JWTService, pasienController controller.PasienController, dokterController controller.DokterController, transaksiController controller.TransaksiController, ruanganController controller.RuanganController, perawatController controller.PerawatController, obatController controller.ObatController, sesidokterController controller.SesiDokterController, reservasiController controller.ReservasiController) {

	pasienRoutes := route.Group("/api/pasien")
	{
		pasienRoutes.POST("/register", pasienController.RegisterPasien)
		pasienRoutes.POST("/login", pasienController.LoginPasien)
		pasienRoutes.POST("/newTransaksi", transaksiController.NewTransaksi)
		pasienRoutes.POST("/newReservasi", reservasiController.NewBulkTransaksiReservasi)
		pasienRoutes.PUT("/", middleware.Authenticate(jwtService), pasienController.UpdatePasien)
		pasienRoutes.GET("/me", middleware.Authenticate(jwtService), pasienController.MePasien)
		pasienRoutes.GET("/jadwalDokter", middleware.Authenticate(jwtService), pasienController.Jadwal_Dokter_User)
		pasienRoutes.GET("/ListDokter", middleware.Authenticate(jwtService), dokterController.GetAllDokter)
		pasienRoutes.GET("/PembelianObat/:id", middleware.Authenticate(jwtService), pasienController.GetLatestPembelianObat)
		pasienRoutes.GET("/:nik", pasienController.GetLatestReservation)
	}

	adminRoutes := route.Group("/api/admin")

	adminRoutes.POST("/login", userController.LoginUser)
	adminRoutes.POST("", userController.RegisterUser)

	adminRoutes.Use(middleware.Authenticate(jwtService))
	{
		//get all admin
		adminRoutes.GET("", userController.GetAllUser)
		adminRoutes.DELETE("/", userController.DeleteUser)
		adminRoutes.PUT("/", userController.UpdateUser)
		adminRoutes.GET("/me", userController.MeUser)

		pasienAdminRoutes := adminRoutes.Group("/pasien")
		{
			pasienAdminRoutes.GET("", pasienController.GetAllPasien)
			pasienAdminRoutes.GET("/transaksi", pasienController.Transaksi_Pasien)
			pasienRoutes.DELETE("/", pasienController.DeletePasien)
		}

		dokterRoutes := adminRoutes.Group("/dokter")
		{
			dokterRoutes.POST("/new", dokterController.RegisterDokter)
			dokterRoutes.GET("", dokterController.GetAllDokter)
			dokterRoutes.GET("/:id", dokterController.GetDokterByID)
			dokterRoutes.PUT("/edit", dokterController.UpdateDoctor)
			dokterRoutes.GET("/jadwalDokter", dokterController.Jadwal_Dokter_Admin)
		}

		transaksiRoutes := adminRoutes.Group("/transaksi")
		{

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
