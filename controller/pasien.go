package controller

import (
	"net/http"

	//"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/gin-gonic/gin"
)


type PasienController interface {
	GetAllPasien(ctx *gin.Context)
}

type pasienController struct {
	pasienService services.PasienService
}

func NewPasienController(ps services.PasienService) PasienController {
	return &pasienController{
		pasienService: ps,
	}
}

func (pc *pasienController) GetAllPasien(ctx *gin.Context) {
	
	pasienList, err := pc.pasienService.GetAllPasien(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, pasienList)
}



