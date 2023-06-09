package services

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/repository"
	"github.com/mashingan/smapping"
)

type PerawatService interface {
	RegisterPerawat(ctx context.Context, perawatDTO dto.PerawatCreateDTO) (entities.Perawat, error)
	GetAllPerawat(ctx context.Context) ([]entities.Perawat, error)
	GetPerawatByID(ctx context.Context, PerawatID string) (entities.Perawat, error)
	GetJadwalPerawat(ctx context.Context) ([]dto.JadwalPerawatDTO, error)
}

type perawatService struct {
	perawatRepository repository.PerawatRepository
}

func NewPerawatService(dr repository.PerawatRepository) PerawatService {
	return &perawatService{
		perawatRepository: dr,
	}
}

func (ds *perawatService) RegisterPerawat(ctx context.Context, perawatDTO dto.PerawatCreateDTO) (entities.Perawat, error) {
	perawat := entities.Perawat{}
	err := smapping.FillStruct(&perawat, smapping.MapFields(perawatDTO))
	if err != nil {
		return entities.Perawat{}, err
	}
	return ds.perawatRepository.RegisterPerawat(ctx, perawat)
}
func (ds *perawatService) GetAllPerawat(ctx context.Context) ([]entities.Perawat, error) {
	return ds.perawatRepository.GetAllPerawat(ctx)
}

func (ds *perawatService) GetJadwalPerawat(ctx context.Context) ([]dto.JadwalPerawatDTO, error) {
	return ds.perawatRepository.GetJadwalPerawat(ctx)
}

func (ds *perawatService) GetPerawatByID(ctx context.Context, PerawatID string) (entities.Perawat, error) {
	return ds.perawatRepository.GetPerawatByID(ctx, PerawatID)
}
