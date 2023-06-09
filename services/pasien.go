package services

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/repository"
	"github.com/google/uuid"
	"github.com/mashingan/smapping"
)

type PasienService interface {
	RegisterPasien(ctx context.Context, pasienDTO dto.PasienCreateDTO) (entities.Pasien, error)
	GetAllPasien(ctx context.Context) ([]entities.Pasien, int, error)
	GetPasienByID(ctx context.Context, PasienID uuid.UUID) (entities.Pasien, error)
	GetPasienByEmail(ctx context.Context, Email string) (entities.Pasien, error)
	UpdatePasien(ctx context.Context, UpdatePasienDTO dto.UpdatePasienDTO) error
	DeletePasien(ctx context.Context, UserID uuid.UUID) error
	CheckPasien(ctx context.Context, email string) (bool, error)
	Verify(ctx context.Context, email string, password string) (bool, error)
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

func (ps *pasienService) GetPasienByID(ctx context.Context, PasienID uuid.UUID) (entities.Pasien, error) {
	return ps.pasienRepo.GetPasienByID(ctx, PasienID)
}

func (ps *pasienService) GetPasienByEmail(ctx context.Context, email string) (entities.Pasien, error) {
	return ps.pasienRepo.GetPasienByEmail(ctx, email)
}

func (ps *pasienService) DeletePasien(ctx context.Context, PasienID uuid.UUID) error {
	return ps.pasienRepo.DeletePasien(ctx, PasienID)
}

func (ps *pasienService) UpdatePasien(ctx context.Context, UpdatePasienDTO dto.UpdatePasienDTO) error {
	pasien := entities.Pasien{}
	if err := smapping.FillStruct(&pasien, smapping.MapFields(UpdatePasienDTO)); err != nil {
		return nil
	}
	return ps.pasienRepo.UpdatePasien(ctx, pasien)
}

func (ps *pasienService) CheckPasien(ctx context.Context, email string) (bool, error) {
	res, err := ps.pasienRepo.GetPasienByEmail(ctx, email)
	if err != nil {
		return false, err
	}

	if res.Email == "" {
		return false, err
	}
	return true, nil
}

func (ps *pasienService) Verify(ctx context.Context, email string, password string) (bool, error) {
	res, err := ps.pasienRepo.GetPasienByEmail(ctx, email)
	if err != nil {
		return false, err
	}

	if res.Email == email && res.Password == password {
		return true, nil
	}

	return false, nil
}
