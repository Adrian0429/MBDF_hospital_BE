package controller

import (

	//"github.com/Caknoooo/golang-clean_template/entities"
	"net/http"

	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/gin-gonic/gin"
)

type RuanganController interface {
	GetAllRuangan(ctx *gin.Context)
}

type ruanganController struct {
	ruanganService services.RuanganService
}

func NewRuanganController(rs services.RuanganService) RuanganController {
	return &ruanganController{
		ruanganService: rs,
	}
}

func (rc *ruanganController) GetAllRuangan(ctx *gin.Context) {
	ruangan, err := rc.ruanganService.GetAllRuangan(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, ruangan)
}
