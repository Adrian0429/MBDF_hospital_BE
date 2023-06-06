package controller

import (
	"net/http"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/Caknoooo/golang-clean_template/utils"
	"github.com/gin-gonic/gin"
)

type PasienController interface {
	RegisterPasien(ctx *gin.Context)
	GetAllPasien(ctx *gin.Context)
	GetPasienByID(ctx *gin.Context)
}

type pasienController struct {
	pasienService services.PasienService
}

func NewPasienController(ps services.PasienService) PasienController {
	return &pasienController{
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

	pasienList, err := pc.pasienService.GetAllPasien(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, pasienList)
}

func (pc *pasienController) GetPasienByID(ctx *gin.Context) {
	id := ctx.Param("id")
	pasien, err := pc.pasienService.GetPasienByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "pasien not Found"})
		return
	}
	ctx.JSON(http.StatusOK, pasien)
}
