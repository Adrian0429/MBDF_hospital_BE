package services

import (
    "context"

    "github.com/Caknoooo/golang-clean_template/entities"
    "github.com/Caknoooo/golang-clean_template/repository"
)

type PasienService interface {
    GetAllPasien(ctx context.Context) ([]entities.Pasien, error)
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

func (ps *pasienService) GetAllPasien(ctx context.Context) ([]entities.Pasien, error) {
    pasien, err := ps.pasienRepo.GetAllPasien(ctx)
    if err != nil {
        return nil, err
    }
    return pasien, nil
}


func (ps *pasienService) GetPasienByID(ctx context.Context, PasienID string) (entities.Pasien, error) {
    return ps.pasienRepo.GetPasienByID(ctx, PasienID)
}
