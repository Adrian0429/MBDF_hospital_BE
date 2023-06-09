package controller

import (
	"net/http"

	//"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/Caknoooo/golang-clean_template/utils"
	"github.com/gin-gonic/gin"
)

type SesiDokterController interface {
	RegisterSesiDokter(ctx *gin.Context)
	GetAllSesiDokter(ctx *gin.Context)
	GetSesiDokterByID(ctx *gin.Context)
}

type sesidokterController struct {
	sesidokterService services.SesiDokterService
}

func NewSesiDokterController(sds services.SesiDokterService) SesiDokterController {
	return &sesidokterController{
		sesidokterService: sds,
	}
}

func (sdc *sesidokterController) RegisterSesiDokter(ctx *gin.Context) {
	var sesidokter dto.SesiDokterCreateDTO
	if err := ctx.ShouldBind(&sesidokter); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, err := sdc.sesidokterService.RegisterSesiDokter(ctx.Request.Context(), sesidokter)

	if err != nil {
		res := utils.BuildResponseFailed("Gagal Menambahkan SesiDokter", "Failed", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Menambahkan SesiDokter", result)
	ctx.JSON(http.StatusOK, res)
}

func (sdc *sesidokterController) GetAllSesiDokter(ctx *gin.Context) {
	sesidokter, err := sdc.sesidokterService.GetAllSesiDokter(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, sesidokter)
}

func (sdc *sesidokterController) GetSesiDokterByID(ctx *gin.Context) {
	id := ctx.Param("id")
	sesidokter, err := sdc.sesidokterService.GetSesiDokterByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "SesiDokter not found"})
		return
	}
	ctx.JSON(http.StatusOK, sesidokter)
}
