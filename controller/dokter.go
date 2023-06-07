package controller

import (
	"net/http"

	//"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/Caknoooo/golang-clean_template/utils"
	"github.com/gin-gonic/gin"
)

type DokterController interface {
	RegisterDokter(ctx *gin.Context)
	GetAllDokter(ctx *gin.Context)
	GetDokterByID(ctx *gin.Context)
}

type dokterController struct {
	dokterService services.DokterService
}

func NewDokterController(ds services.DokterService) DokterController {
	return &dokterController{
		dokterService: ds,
	}
}

func (dc *dokterController) RegisterDokter(ctx *gin.Context) {
	var dokter dto.DokterCreateDTO
	if err := ctx.ShouldBind(&dokter); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, err := dc.dokterService.RegisterDokter(ctx.Request.Context(), dokter)

	if err != nil {
		res := utils.BuildResponseFailed("Gagal Menambahkan Dokter", "Failed", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Menambahkan Dokter", result)
	ctx.JSON(http.StatusOK, res)
}

func (dc *dokterController) GetAllDokter(ctx *gin.Context) {
	dokter, err := dc.dokterService.GetAllDokter(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dokter)
}

func (dc *dokterController) GetDokterByID(ctx *gin.Context) {
	id := ctx.Param("id")
	doctor, err := dc.dokterService.GetDokterByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found"})
		return
	}
	ctx.JSON(http.StatusOK, doctor)
}
