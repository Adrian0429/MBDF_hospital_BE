package controller

import (
	"net/http"

	//"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/Caknoooo/golang-clean_template/utils"
	"github.com/gin-gonic/gin"
)

type PerawatController interface {
	RegisterPerawat(ctx *gin.Context)
	GetAllPerawat(ctx *gin.Context)
	GetPerawatByID(ctx *gin.Context)
}

type perawatController struct {
	perawatService services.PerawatService
}

func NewPerawatController(ds services.PerawatService) PerawatController {
	return &perawatController{
		perawatService: ds,
	}
}

func (dc *perawatController) RegisterPerawat(ctx *gin.Context) {
	var perawat dto.PerawatCreateDTO
	if err := ctx.ShouldBind(&perawat); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, err := dc.perawatService.RegisterPerawat(ctx.Request.Context(), perawat)

	if err != nil {
		res := utils.BuildResponseFailed("Gagal Menambahkan Perawat", "Failed", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Menambahkan Perawat", result)
	ctx.JSON(http.StatusOK, res)
}

func (dc *perawatController) GetAllPerawat(ctx *gin.Context) {
	perawat, err := dc.perawatService.GetAllPerawat(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, perawat)
}

func (dc *perawatController) GetPerawatByID(ctx *gin.Context) {
	id := ctx.Param("id")
	perawat, err := dc.perawatService.GetPerawatByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Perawat not found"})
		return
	}
	ctx.JSON(http.StatusOK, perawat)
}
