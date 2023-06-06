package services

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/repository"
)

type DokterService interface {
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

func (ds *dokterService) GetAllDokter(ctx context.Context) ([]entities.Dokter, error) {
	return ds.dokterRepository.GetAllDokter(ctx)
}

func (ds *dokterService) GetDokterByID(ctx context.Context, DokterID string) (entities.Dokter, error) {
	return ds.dokterRepository.GetDokterByID(ctx, DokterID)
}
