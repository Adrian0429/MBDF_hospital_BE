package main

import (
	"os"

	"github.com/Caknoooo/golang-clean_template/config"
	"github.com/Caknoooo/golang-clean_template/controller"
	"github.com/Caknoooo/golang-clean_template/middleware"
	"github.com/Caknoooo/golang-clean_template/repository"
	"github.com/Caknoooo/golang-clean_template/routes"
	"github.com/Caknoooo/golang-clean_template/services"
	
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	var (
		db             *gorm.DB                  = config.SetUpDatabaseConnection()
		jwtService     services.JWTService       = services.NewJWTService()
		userRepository repository.UserRepository = repository.NewUserRepository(db)
		userService    services.UserService      = services.NewUserService(userRepository)
		userController controller.UserController = controller.NewUserController(userService, jwtService)

		pasienRepository repository.PasienRepository = repository.NewPasienRepository(db)
		pasienService    services.PasienService      = services.NewPasienService(pasienRepository)
		pasienController controller.PasienController = controller.NewPasienController(pasienService, jwtService)

		dokterRepository repository.DokterRepository = repository.NewDokterRepository(db)
		dokterService    services.DokterService      = services.NewDokterService(dokterRepository)
		dokterController controller.DokterController = controller.NewDokterController(dokterService)

		transaksiRepository repository.TransaksiRepository = repository.NewTransaksiRepository(db)
		transaksiService    services.TransaksiService      = services.NewTransaksiService(transaksiRepository)
		transaksiController controller.TransaksiController = controller.NewTransaksiController(transaksiService)

		ruanganRepository repository.RuanganRepository = repository.NewRuanganRepository(db)
		ruanganService    services.RuanganService      = services.NewRuanganService(ruanganRepository)
		ruanganController controller.RuanganController = controller.NewRuanganController(ruanganService)

		perawatRepository repository.PerawatRepository = repository.NewPerawatRepository(db)
		perawatService    services.PerawatService      = services.NewPerawatService(perawatRepository)
		perawatController controller.PerawatController = controller.NewPerawatController(perawatService)

		obatRepository repository.ObatRepository = repository.NewObatRepository(db)
		obatService    services.ObatService      = services.NewObatService(obatRepository)
		obatController controller.ObatController = controller.NewObatController(obatService)

		sesidokterRepository repository.SesiDokterRepository = repository.NewSesiDokterRepository(db)
		sesidokterService    services.SesiDokterService      = services.NewSesiDokterService(sesidokterRepository)
		sesidokterController controller.SesiDokterController = controller.NewSesiDokterController(sesidokterService)
	)

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())
	routes.Router(server, userController, jwtService, pasienController, dokterController, transaksiController, ruanganController, perawatController, obatController, sesidokterController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}
	server.Run(":" + port)
}
