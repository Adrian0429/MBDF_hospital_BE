package controller

import (
	//"github.com/Caknoooo/golang-clean_template/entities"
	"net/http"

	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/Caknoooo/golang-clean_template/utils"
	"github.com/gin-gonic/gin"
)

type TransaksiController interface {
	GetTransaksiByIDPasien(ctx *gin.Context)
	GetAllTransaksi(ctx *gin.Context)
}

type transaksiController struct {
	transaksiService services.TransaksiService
}

func NewTransaksiController(ts services.TransaksiService) TransaksiController {
	return &transaksiController{
		transaksiService: ts,
	}
}

func (tc *transaksiController) GetAllTransaksi(ctx *gin.Context) {
	result, err := tc.transaksiService.GetAllTransaksi(ctx.Request.Context())
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan List Transaksi", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess("Berhasil Mendapatkan List Transaksi", result)
	ctx.JSON(http.StatusOK, res)
}

func (tc *transaksiController) GetTransaksiByIDPasien(ctx *gin.Context) {
	id := ctx.Param("id")
	transaksi, err := tc.transaksiService.GetTransaksiByIDPasien(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "transaksi not found"})
		return
	}

	ctx.JSON(http.StatusOK, transaksi)
}
