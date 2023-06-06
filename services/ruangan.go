package services

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/repository"
)

type RuanganService interface {
	GetAllRuangan(ctx context.Context) ([]entities.Ruangan, error)
}

type ruanganService struct {
	ruanganRepo repository.RuanganRepository
}

func NewRuanganService(ruanganRepo repository.RuanganRepository) RuanganService {
	return &ruanganService{
		ruanganRepo: ruanganRepo,
	}
}

func (rs *ruanganService) GetAllRuangan(ctx context.Context) ([]entities.Ruangan, error) {
	ruangan, err := rs.ruanganRepo.GetAllRuangan(ctx)
	if err != nil {
		return nil, err
	}
	return ruangan, nil
}
