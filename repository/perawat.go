package repository

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"gorm.io/gorm"
)

type PerawatRepository interface {
	RegisterPerawat(ctx context.Context, perawat entities.Perawat) (entities.Perawat, error)
	GetAllPerawat(ctx context.Context) ([]entities.Perawat, error)
	GetPerawatByID(ctx context.Context, PerawatID string) (entities.Perawat, error)
	GetJadwalPerawat(ctx context.Context) ([]dto.JadwalPerawatDTO, error)
}

type perawatRepository struct {
	connection *gorm.DB
}

func NewPerawatRepository(db *gorm.DB) PerawatRepository {
	return &perawatRepository{
		connection: db,
	}
}

func (dr *perawatRepository) RegisterPerawat(ctx context.Context, perawat entities.Perawat) (entities.Perawat, error) {
	if err := dr.connection.Table("perawats").Create(&perawat).Error; err != nil {
		return entities.Perawat{}, err
	}
	return perawat, nil
}

func (dr *perawatRepository) GetAllPerawat(ctx context.Context) ([]entities.Perawat, error) {
	var perawat []entities.Perawat
	if err := dr.connection.Table("perawats").Find(&perawat).Error; err != nil {
		return nil, err
	}
	return perawat, nil
}

func (dr *perawatRepository) GetJadwalPerawat(ctx context.Context) ([]dto.JadwalPerawatDTO, error) {
	var jadwal_perawat []dto.JadwalPerawatDTO
	if err := dr.connection.Table("sesi_jaga_nginaps").Find(&jadwal_perawat).Error; err != nil {
		return nil, err
	}
	return jadwal_perawat, nil
}

func (dr *perawatRepository) GetPerawatByID(ctx context.Context, PerawatID string) (entities.Perawat, error) {
	var perawat entities.Perawat
	if err := dr.connection.Table("perawats").Where("ID_Perawat = ?", PerawatID).First(&perawat).Error; err != nil {
		return entities.Perawat{}, err
	}
	return perawat, nil
}
