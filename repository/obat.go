package repository

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"gorm.io/gorm"
)

type ObatRepository interface {
	RegisterObat(ctx context.Context, obat entities.Obat) (entities.Obat, error)
	GetAllObat(ctx context.Context) ([]dto.AllObat, error)
	GetObatByID(ctx context.Context, ObatID string) (entities.Obat, error)
}

type obatRepository struct {
	connection *gorm.DB
}

func NewObatRepository(db *gorm.DB) ObatRepository {
	return &obatRepository{
		connection: db,
	}
}

func (dr *obatRepository) RegisterObat(ctx context.Context, obat entities.Obat) (entities.Obat, error) {
	if err := dr.connection.Table("obats").Create(&obat).Error; err != nil {
		return entities.Obat{}, err
	}
	return obat, nil
}

func (dr *obatRepository) GetAllObat(ctx context.Context) ([]dto.AllObat, error) {
	var obat []dto.AllObat

	if err := dr.connection.Raw("select * from view_obat").Scan(&obat).Error; err != nil {
		return nil, err
	}
	return obat, nil
}

func (dr *obatRepository) GetObatByID(ctx context.Context, ObatID string) (entities.Obat, error) {
	var obat entities.Obat
	if err := dr.connection.Table("obats").Where("ID_Obat = ?", ObatID).First(&obat).Error; err != nil {
		return entities.Obat{}, err
	}
	return obat, nil
}
