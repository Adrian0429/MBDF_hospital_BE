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
	GetAllDokter(ctx context.Context) ([]dto.AllDokterDTO, error)
	AllDokterUser(ctx context.Context) ([]dto.AllDokterUserDTO, error)
	GetDokterByID(ctx context.Context, DokterID string) (entities.Dokter, error)
	UpdateDoctor(ctx context.Context, DokterDTO dto.UpdateDokterDTO) error
	Jadwal_Dokter_Admin(ctx context.Context) ([]dto.Jadwal_Dokter_AdminDTO, error)
	GetPolis(ctx context.Context) ([]dto.PoliDTO, error)
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

func (ds *dokterService) GetAllDokter(ctx context.Context) ([]dto.AllDokterDTO, error) {
	return ds.dokterRepository.GetAllDokter(ctx)
}

func (ds *dokterService) AllDokterUser(ctx context.Context) ([]dto.AllDokterUserDTO, error) {
	return ds.dokterRepository.AllDokterUser(ctx)
}

func (ds *dokterService) UpdateDoctor(ctx context.Context, DokterDTO dto.UpdateDokterDTO) error {
	dokter := entities.Dokter{}
	if err := smapping.FillStruct(&dokter, smapping.MapFields(DokterDTO)); err != nil {
		return nil
	}
	return ds.dokterRepository.UpdateDoctor(ctx, dokter)
}

func (ds *dokterService) GetDokterByID(ctx context.Context, DokterID string) (entities.Dokter, error) {
	return ds.dokterRepository.GetDokterByID(ctx, DokterID)
}

func (ds *dokterService) Jadwal_Dokter_Admin(ctx context.Context) ([]dto.Jadwal_Dokter_AdminDTO, error) {
	return ds.dokterRepository.Jadwal_Dokter_Admin(ctx)
}

func (ds *dokterService) GetPolis(ctx context.Context) ([]dto.PoliDTO, error) {
	return ds.dokterRepository.GetPolis(ctx)
}
