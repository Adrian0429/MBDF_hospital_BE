package repository

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/entities"
	"gorm.io/gorm"
)

type DokterRepository interface {
	RegisterDokter(ctx context.Context, dokter entities.Dokter) (entities.Dokter, error)
	GetAllDokter(ctx context.Context) ([]entities.Dokter, error)
	GetDokterByID(ctx context.Context, DokterID string) (entities.Dokter, error)
}

type dokterRepository struct {
	connection *gorm.DB
}

func NewDokterRepository(db *gorm.DB) DokterRepository {
	return &dokterRepository{
		connection: db,
	}
}

func (dr *dokterRepository) RegisterDokter(ctx context.Context, dokter entities.Dokter) (entities.Dokter, error) {
	if err := dr.connection.Table("dokters").Create(&dokter).Error; err != nil {
		return entities.Dokter{}, err
	}
	return dokter, nil
}
func (dr *dokterRepository) GetAllDokter(ctx context.Context) ([]entities.Dokter, error) {
	var dokter []entities.Dokter
	if err := dr.connection.Table("dokters").Find(&dokter).Error; err != nil {
		return nil, err
	}
	return dokter, nil
}

func (dr *dokterRepository) GetDokterByID(ctx context.Context, DokterID string) (entities.Dokter, error) {
	var dokter entities.Dokter
	if err := dr.connection.Table("dokters").Where("ID_Dokter = ?", DokterID).First(&dokter).Error; err != nil {
		return entities.Dokter{}, err
	}
	return dokter, nil
}
