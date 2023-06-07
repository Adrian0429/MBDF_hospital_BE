package routes

import (
	"github.com/Caknoooo/golang-clean_template/controller"
	"github.com/Caknoooo/golang-clean_template/middleware"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/gin-gonic/gin"
)

func Router(route *gin.Engine, userController controller.UserController, jwtService services.JWTService, pasienController controller.PasienController, dokterController controller.DokterController, transaksiController controller.TransaksiController, ruanganController controller.RuanganController, perawatController controller.PerawatController, obatController controller.ObatController) {
	routes := route.Group("/api/user")
	{
		routes.POST("", userController.RegisterUser)
		routes.GET("", middleware.Authenticate(jwtService), userController.GetAllUser)
		routes.POST("/login", userController.LoginUser)
		routes.DELETE("/", middleware.Authenticate(jwtService), userController.DeleteUser)
		routes.PUT("/", middleware.Authenticate(jwtService), userController.UpdateUser)
		routes.GET("/me", middleware.Authenticate(jwtService), userController.MeUser)
	}

	pasienRoutes := route.Group("/api/pasien")
	{
		pasienRoutes.POST("/register", pasienController.RegisterPasien)
		pasienRoutes.GET("", pasienController.GetAllPasien)
		// Add other pasien routes here if needed
	}

	dokterRoutes := route.Group("/api/dokter")
	{
		dokterRoutes.POST("/new", dokterController.RegisterDokter)
		dokterRoutes.GET("", dokterController.GetAllDokter)
		dokterRoutes.GET("/:id", dokterController.GetDokterByID)
	}

	transaksiRoutes := route.Group("/api/transaksi")
	{
		transaksiRoutes.POST("/new", transaksiController.NewTransaksi)
		transaksiRoutes.GET("/:id", transaksiController.GetTransaksiByIDPasien)
		transaksiRoutes.GET("/all", transaksiController.GetAllTransaksi)
	}

	ruanganRoutes := route.Group("/api/ruangan")
	{

		ruanganRoutes.GET("", ruanganController.GetAllRuangan)
	}

	perawatRoutes := route.Group("/api/perawat")
	{
		perawatRoutes.POST("/new", perawatController.RegisterPerawat)
		perawatRoutes.GET("", perawatController.GetAllPerawat)
		perawatRoutes.GET("/:id", perawatController.GetPerawatByID)
	}

	obatRoutes := route.Group("/api/obat")
	{
		obatRoutes.POST("/new", obatController.RegisterObat)
		obatRoutes.GET("", obatController.GetAllObat)
		obatRoutes.GET("/:id", obatController.GetObatByID)
	}

}
