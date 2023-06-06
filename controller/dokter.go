package controller

import (
	"net/http"

	//"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/gin-gonic/gin"
)

type DokterController interface {
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
