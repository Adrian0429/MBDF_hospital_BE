package controller

import (
	"context"
	"net/http"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"

	// "github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/Caknoooo/golang-clean_template/utils"
	"github.com/gin-gonic/gin"
)

type PasienController interface {
	RegisterPasien(ctx *gin.Context)
	GetAllPasien(ctx *gin.Context)
	MePasien(ctx *gin.Context)
	LoginPasien(ctx *gin.Context)
	UpdatePasien(ctx *gin.Context)
	DeletePasien(ctx *gin.Context)
	GetLatestPembelianObat(ctx *gin.Context)
}

type pasienController struct {
	jwtService    services.JWTService
	pasienService services.PasienService
}

func NewPasienController(ps services.PasienService, jwt services.JWTService) PasienController {
	return &pasienController{
		jwtService:    jwt,
		pasienService: ps,
	}
}

func (pc *pasienController) RegisterPasien(ctx *gin.Context) {
	var pasien dto.PasienCreateDTO
	if err := ctx.ShouldBind(&pasien); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := pc.pasienService.RegisterPasien(ctx.Request.Context(), pasien)

	if err != nil {
		res := utils.BuildResponseFailed("Gagal Menambahkan Pasien", "Failed", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Menambahkan Pasien", result)
	ctx.JSON(http.StatusOK, res)
}

func (pc *pasienController) GetAllPasien(ctx *gin.Context) {
	pasienList, count, err := pc.pasienService.GetAllPasien(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	dtoList := make([]dto.GetAllPasienDTO, len(pasienList))
	for i, pasien := range pasienList {
		dtoList[i] = dto.GetAllPasienDTO{
			Count:         count,
			NIK_pasien:    pasien.NIK_pasien,
			Nama_Pasien:   pasien.Nama_Pasien,
			Jenis_Kelamin: pasien.Jenis_Kelamin,
			Tanggal_Lahir: pasien.Tanggal_Lahir,
			No_Telepon:    pasien.No_Telepon,
		}
	}

	ctx.JSON(http.StatusOK, dtoList)
}

func (pc *pasienController) MePasien(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	PasienID, err := pc.jwtService.GetUserIDByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	result, err := pc.pasienService.GetPasienByID(ctx.Request.Context(), PasienID)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Pasien", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan Pasien", result)
	ctx.JSON(http.StatusOK, res)
}

func (pc *pasienController) LoginPasien(ctx *gin.Context) {
	var PasienLoginDTO dto.PasienLoginDTO
	err := ctx.ShouldBind(&PasienLoginDTO)
	res, _ := pc.pasienService.Verify(ctx.Request.Context(), PasienLoginDTO.Email, PasienLoginDTO.Password)
	if !res {
		response := utils.BuildResponseFailed("Gagal Login", "Email atau Password Salah", utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	pasien, err := pc.pasienService.GetPasienByEmail(ctx.Request.Context(), PasienLoginDTO.Email)
	if err != nil {
		response := utils.BuildResponseFailed("Gagal Login", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	token := pc.jwtService.GenerateToken(pasien.Uid, "pasien")
	userResponse := entities.Authorization{
		Token: token,
		Role:  "pasien",
	}

	response := utils.BuildResponseSuccess("Berhasil Login", userResponse)
	ctx.JSON(http.StatusOK, response)
}

func (pc *pasienController) UpdatePasien(ctx *gin.Context) {
	var UpdatePasienDTO dto.UpdatePasienDTO
	if err := ctx.ShouldBind(&UpdatePasienDTO); err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Request Dari Body", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	token := ctx.MustGet("token").(string)
	userID, err := pc.jwtService.GetUserIDByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	UpdatePasienDTO.Uid = userID
	if err = pc.pasienService.UpdatePasien(ctx.Request.Context(), UpdatePasienDTO); err != nil {
		res := utils.BuildResponseFailed("Gagal Update Pasien", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Update Pasien", UpdatePasienDTO)
	ctx.JSON(http.StatusOK, res)
}

func (pc *pasienController) DeletePasien(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := pc.jwtService.GetUserIDByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	if err = pc.pasienService.DeletePasien(ctx.Request.Context(), userID); err != nil {
		res := utils.BuildResponseFailed("Gagal Menghapus Pasien", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Menghapus Pasien", utils.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func (c *pasienController) GetLatestPembelianObat(ctx *gin.Context) {
	pasienID := ctx.Param("id")

	// Call the service to get the latest pembelian obat
	pembelianObat, err := c.pasienService.GetLatestPembelianObat(context.Background(), pasienID)
	if err != nil {
		// Handle the error appropriately
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get latest pembelian obat"})
		return
	}

	ctx.JSON(http.StatusOK, pembelianObat)
}
