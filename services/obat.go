package services

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/repository"
	"github.com/mashingan/smapping"
)

type ObatService interface {
	RegisterObat(ctx context.Context, obatDTO dto.ObatCreateDTO) (entities.Obat, error)
	GetAllObat(ctx context.Context) ([]entities.Obat, error)
	GetObatByID(ctx context.Context, ObatID string) (entities.Obat, error)
}

type obatService struct {
	obatRepository repository.ObatRepository
}

func NewObatService(dr repository.ObatRepository) ObatService {
	return &obatService{
		obatRepository: dr,
	}
}

func (ds *obatService) RegisterObat(ctx context.Context, obatDTO dto.ObatCreateDTO) (entities.Obat, error) {
	obat := entities.Obat{}
	err := smapping.FillStruct(&obat, smapping.MapFields(obatDTO))
	if err != nil {
		return entities.Obat{}, err
	}
	return ds.obatRepository.RegisterObat(ctx, obat)
}
func (ds *obatService) GetAllObat(ctx context.Context) ([]entities.Obat, error) {
	return ds.obatRepository.GetAllObat(ctx)
}

func (ds *obatService) GetObatByID(ctx context.Context, ObatID string) (entities.Obat, error) {
	return ds.obatRepository.GetObatByID(ctx, ObatID)
}
