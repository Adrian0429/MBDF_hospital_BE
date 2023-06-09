package services

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/repository"
	"github.com/mashingan/smapping"
)

type SesiDokterService interface {
	RegisterSesiDokter(ctx context.Context, sesidokterDTO dto.SesiDokterCreateDTO) (entities.Sesi_Dokter, error)
	GetAllSesiDokter(ctx context.Context) ([]dto.AmbilSesiDTO, error)
	GetSesiDokterByID(ctx context.Context, SesiDokterID string) (entities.Sesi_Dokter, error)
}

type sesidokterService struct {
	sesidokterRepository repository.SesiDokterRepository
}

func NewSesiDokterService(sdr repository.SesiDokterRepository) SesiDokterService {
	return &sesidokterService{
		sesidokterRepository: sdr,
	}
}

func (sds *sesidokterService) RegisterSesiDokter(ctx context.Context, sesidokterDTO dto.SesiDokterCreateDTO) (entities.Sesi_Dokter, error) {
	sesidokter := entities.Sesi_Dokter{}
	err := smapping.FillStruct(&sesidokter, smapping.MapFields(sesidokterDTO))
	if err != nil {
		return entities.Sesi_Dokter{}, err
	}
	return sds.sesidokterRepository.RegisterSesiDokter(ctx, sesidokter)
}

func (sds *sesidokterService) GetAllSesiDokter(ctx context.Context) ([]dto.AmbilSesiDTO, error) {
	return sds.sesidokterRepository.GetAllSesiDokter(ctx)
}

func (sds *sesidokterService) GetSesiDokterByID(ctx context.Context, SesiDokterID string) (entities.Sesi_Dokter, error) {
	return sds.sesidokterRepository.GetSesiDokterByID(ctx, SesiDokterID)
}
