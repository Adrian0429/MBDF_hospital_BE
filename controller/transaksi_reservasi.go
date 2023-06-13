package controller

import (
	"net/http"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/Caknoooo/golang-clean_template/utils"
	"github.com/gin-gonic/gin"
)

type ReservasiController interface {
	NewReservasi(ctx *gin.Context)
}

type reservasiController struct {
	reservasiService services.ReservasiService
}

func NewReservasiController(rs services.ReservasiService) ReservasiController {
	return &reservasiController{
		reservasiService: rs,
	}
}

func (tr *reservasiController) NewReservasi(ctx *gin.Context) {
	var reservasi dto.NewTransaksiReservasiDTO
	if err := ctx.ShouldBind(&reservasi); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, err := tr.reservasiService.NewReservasi(ctx.Request.Context(), reservasi)

	if err != nil {
		res := utils.BuildResponseFailed("Gagal Menambahkan reservasi", "Failed", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("berhasil menambahkan reservasi", result)
	ctx.JSON(http.StatusOK, res)
}
