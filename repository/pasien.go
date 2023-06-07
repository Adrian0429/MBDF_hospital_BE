package repository

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/entities"
	"gorm.io/gorm"
)

type PasienRepository interface {
	RegisterPasien(ctx context.Context, pasien entities.Pasien) (entities.Pasien, error)
	GetAllPasien(ctx context.Context) ([]entities.Pasien, int, error)
	GetPasienByID(ctx context.Context, PasienID string) (entities.Pasien, error)
}

type pasienRepository struct {
	connection *gorm.DB
}

func NewPasienRepository(db *gorm.DB) PasienRepository {
	return &pasienRepository{
		connection: db,
	}
}

func (pr *pasienRepository) RegisterPasien(ctx context.Context, pasien entities.Pasien) (entities.Pasien, error) {
	if err := pr.connection.Table("pasiens").Create(&pasien).Error; err != nil {
		return entities.Pasien{}, err
	}
	return pasien, nil
}

func (pr *pasienRepository) GetAllPasien(ctx context.Context) ([]entities.Pasien, int, error) {
	var pasien []entities.Pasien
	var count int64

	if err := pr.connection.Table("pasiens").Find(&pasien).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return pasien, int(count), nil
}

func (pr *pasienRepository) GetPasienByID(ctx context.Context, PasienID string) (entities.Pasien, error) {
	var pasien entities.Pasien

	if err := pr.connection.Table("pasiens").Where("id = ?", PasienID).Take(&pasien).Error; err != nil {
		return entities.Pasien{}, err
	}
	return pasien, nil
}
