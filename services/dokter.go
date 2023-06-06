package services

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/repository"
	"github.com/mashingan/smapping"
)

type DokterService interface {
	RegisterDokter(ctx context.Context, dokterDTO dto.DokterCreateDTO) (entities.Dokter, error)
	GetAllDokter(ctx context.Context) ([]entities.Dokter, error)
	GetDokterByID(ctx context.Context, DokterID string) (entities.Dokter, error)
}

type dokterService struct {
	dokterRepository repository.DokterRepository
}

func NewDokterService(dr repository.DokterRepository) DokterService {
	return &dokterService{
		dokterRepository: dr,
	}
}

func (ds *dokterService) RegisterDokter(ctx context.Context, dokterDTO dto.DokterCreateDTO) (entities.Dokter, error) {
	dokter := entities.Dokter{}
	err := smapping.FillStruct(&dokter, smapping.MapFields(dokterDTO))
	if err != nil {
		return entities.Dokter{}, err
	}
	return ds.dokterRepository.RegisterDokter(ctx, dokter)
}
func (ds *dokterService) GetAllDokter(ctx context.Context) ([]entities.Dokter, error) {
	return ds.dokterRepository.GetAllDokter(ctx)
}

func (ds *dokterService) GetDokterByID(ctx context.Context, DokterID string) (entities.Dokter, error) {
	return ds.dokterRepository.GetDokterByID(ctx, DokterID)
}
