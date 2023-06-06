package repository

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/entities"
	"gorm.io/gorm"
)

type RuanganRepository interface {
	GetAllRuangan(ctx context.Context) ([]entities.Ruangan, error)
}

type ruanganRepository struct {
	connection *gorm.DB
}

func NewRuanganRepository(db *gorm.DB) RuanganRepository {
	return &ruanganRepository{
		connection: db,
	}
}

func (rr *ruanganRepository) GetAllRuangan(ctx context.Context) ([]entities.Ruangan, error) {
	var ruangan []entities.Ruangan
	if err := rr.connection.Table("ruangan").Find(&ruangan).Error; err != nil {
		return nil, err
	}
	return ruangan, nil
}


