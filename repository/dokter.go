package repository

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/entities"
	"gorm.io/gorm"
)

type DokterRepository interface { 
	GetAllDokter(ctx context.Context)([]entities.Dokter, error)
	GetDokterByID(ctx context.Context, DokterID string)(entities.Dokter, error)
}

type dokterRepository struct {
    connection *gorm.DB
}

func NewDokterRepository(db *gorm.DB) DokterRepository {
    return &dokterRepository{
        connection: db,
    }
}

func (dr *dokterRepository) GetAllDokter(ctx context.Context)([]entities.Dokter, error){
	var dokter []entities.Dokter
	if err := dr.connection.Table("dokter").Find(&dokter).Error; err != nil{
		return nil, err
	}
	return dokter, nil
}

func (dr *dokterRepository) GetDokterByID(ctx context.Context, DokterID string)(entities.Dokter, error){
	var dokter entities.Dokter
	if err := dr.connection.Table("dokter").Where("ID_Dokter = ?", DokterID).First(&dokter).Error; err != nil {
		return entities.Dokter{}, err
	}
	return dokter, nil
}