package controller

import (
	"net/http"

	//"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/Caknoooo/golang-clean_template/utils"
	"github.com/gin-gonic/gin"
)

type ObatController interface {
	RegisterObat(ctx *gin.Context)
	GetAllObat(ctx *gin.Context)
	GetObatByID(ctx *gin.Context)
}

type obatController struct {
	obatService services.ObatService
}

func NewObatController(ds services.ObatService) ObatController {
	return &obatController{
		obatService: ds,
	}
}

func (dc *obatController) RegisterObat(ctx *gin.Context) {
	var obat dto.ObatCreateDTO
	if err := ctx.ShouldBind(&obat); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, err := dc.obatService.RegisterObat(ctx.Request.Context(), obat)

	if err != nil {
		res := utils.BuildResponseFailed("Gagal Menambahkan Obat", "Failed", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Menambahkan Obat", result)
	ctx.JSON(http.StatusOK, res)
}

func (dc *obatController) GetAllObat(ctx *gin.Context) {
	obat, err := dc.obatService.GetAllObat(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, obat)
}

func (dc *obatController) GetObatByID(ctx *gin.Context) {
	id := ctx.Param("id")
	obat, err := dc.obatService.GetObatByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Obat not found"})
		return
	}
	ctx.JSON(http.StatusOK, obat)
}
