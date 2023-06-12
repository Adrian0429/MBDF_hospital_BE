package repository

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/entities"
	"gorm.io/gorm"
)

type ReservasiRepository interface {
	
}

type reservasiRepository struct {
	connection *gorm.DB
}

func NewReservasiRepository(db *gorm.DB) ReservasiRepository {
	return &reservasiRepository{
		connection: db,
	}
}

