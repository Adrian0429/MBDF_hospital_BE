package repository

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"gorm.io/gorm"
)

type SesiDokterRepository interface {
	RegisterSesiDokter(ctx context.Context, sesiDokter entities.Sesi_Dokter) (entities.Sesi_Dokter, error)
	GetAllSesiDokter(ctx context.Context) ([]dto.AmbilSesiDTO, error)
	GetSesiDokterByID(ctx context.Context, SesiDokterID string) (entities.Sesi_Dokter, error)
}

type sesiDokterRepository struct {
	connection *gorm.DB
}

func NewSesiDokterRepository(db *gorm.DB) SesiDokterRepository {
	return &sesiDokterRepository{
		connection: db,
	}
}

func (sdr *sesiDokterRepository) RegisterSesiDokter(ctx context.Context, sesiDokter entities.Sesi_Dokter) (entities.Sesi_Dokter, error) {
	if err := sdr.connection.Table("sesi_dokters").Create(&sesiDokter).Error; err != nil {
		return entities.Sesi_Dokter{}, err
	}
	return sesiDokter, nil
}

func (sdr *sesiDokterRepository) GetAllSesiDokter(ctx context.Context) ([]dto.AmbilSesiDTO, error) {
	var sesiDokter []dto.AmbilSesiDTO
	query := "SELECT ID_Dokter, Nama_Dokter, Hari, Jam_Masuk, Jam_Keluar, Jenis_Kelamin, No_Telepon, Harga_Konsultasi FROM Dokters JOIN Sesi_Dokters ON ID_Dokter = Dokter_ID_Dokter;"
	// query := "select * from Jadwal_Dokter"
	if err := sdr.connection.Raw(query).Find(&sesiDokter).Error; err != nil {
		return nil, err
	}
	return sesiDokter, nil
}


func (sdr *sesiDokterRepository) GetSesiDokterByID(ctx context.Context, SesiDokterID string) (entities.Sesi_Dokter, error) {
	var sesiDokter entities.Sesi_Dokter
	if err := sdr.connection.Table("sesi_dokters").Where("ID_SesiDokter = ?", SesiDokterID).First(&sesiDokter).Error; err != nil {
		return entities.Sesi_Dokter{}, err
	}
	return sesiDokter, nil
}
