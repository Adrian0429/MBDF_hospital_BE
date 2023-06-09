package controller

import (
	"net/http"
	"strings"

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
	UpdateDoctor(ctx *gin.Context)
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

	dtoList := make([]dto.AllDokterDTO, 0, len(dokter))
	for _, dokterItem := range dokter {
		dtoItem := dto.AllDokterDTO{
			Nama_Dokter:      dokterItem.Nama_Dokter,
			Jenis_Kelamin:    dokterItem.Jenis_Kelamin,
			Tanggal_Lahir:    dokterItem.Tanggal_Lahir,
			No_Telepon:       dokterItem.No_Telepon,
			Harga_Konsultasi: dokterItem.Harga_Konsultasi,
			Poli_Kode_Poli:   mapKodePoli(dokterItem.Poli_Kode_Poli),
		}
		dtoList = append(dtoList, dtoItem)
	}

	ctx.JSON(http.StatusOK, dtoList)
}

func mapKodePoli(kodePoli string) string {

	trimkode := strings.TrimSpace(kodePoli)

	switch trimkode {
	case "PD":
		return "Penyakit Dalam"
	case "O":
		return "Obstetri dan Ginekologi"
	case "IA":
		return "Ibu dan Anak"
	case "M":
		return "Mata"
	case "S":
		return "Saraf"
	default:
		return "unknown"
	}
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

func (dc *dokterController) UpdateDoctor(ctx *gin.Context) {
	var DokterDTO dto.UpdateDokterDTO
	if err := ctx.ShouldBind(&DokterDTO); err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Request Dari Body", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Get the user ID from the map key
	dokterID := ctx.PostForm("ID_Dokter") // Assuming the user ID is sent as a form value

	// Set the user ID in the DTO
	DokterDTO.ID_Dokter = dokterID

	if err := dc.dokterService.UpdateDoctor(ctx.Request.Context(), DokterDTO); err != nil {
		res := utils.BuildResponseFailed("Gagal Update dokter", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Update dokter", DokterDTO)
	ctx.JSON(http.StatusOK, res)
}
