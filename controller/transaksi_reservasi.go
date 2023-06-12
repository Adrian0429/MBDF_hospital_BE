package controller

import (
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/gin-gonic/gin"
)

type ReservasiController interface {
	NewBulkTransaksiReservasi(ctx *gin.Context)
}

type reservasiController struct {
	reservasiService services.ReservasiService
}

func NewReservasiController(rs services.ReservasiService) ReservasiController {
	return &reservasiController{
		reservasiService: rs,
	}
}
