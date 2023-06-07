package services

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/repository"
	"github.com/mashingan/smapping"
)

type PasienService interface {
	RegisterPasien(ctx context.Context, pasienDTO dto.PasienCreateDTO) (entities.Pasien, error)
	GetAllPasien(ctx context.Context) ([]entities.Pasien, int, error)
	GetPasienByID(ctx context.Context, PasienID string) (entities.Pasien, error)
}

type pasienService struct {
	pasienRepo repository.PasienRepository
}

func NewPasienService(pasienRepo repository.PasienRepository) PasienService {
	return &pasienService{
		pasienRepo: pasienRepo,
	}
}

func (ps *pasienService) RegisterPasien(ctx context.Context, pasienDTO dto.PasienCreateDTO) (entities.Pasien, error) {
	pasien := entities.Pasien{}
	err := smapping.FillStruct(&pasien, smapping.MapFields(pasienDTO))
	if err != nil {
		return entities.Pasien{}, err
	}
	return ps.pasienRepo.RegisterPasien(ctx, pasien)
}

func (ps *pasienService) GetAllPasien(ctx context.Context) ([]entities.Pasien, int, error) {
	return ps.pasienRepo.GetAllPasien(ctx)
}

func (ps *pasienService) GetPasienByID(ctx context.Context, PasienID string) (entities.Pasien, error) {
	return ps.pasienRepo.GetPasienByID(ctx, PasienID)
}
