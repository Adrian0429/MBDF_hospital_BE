package services

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/repository"
	"github.com/mashingan/smapping"
)

type TransaksiService interface {
	NewTransaksi(ctx context.Context, transaksiDTO dto.NewTransaksi) (entities.Transaksi, error)
	GetAllTransaksi(ctx context.Context) ([]entities.Transaksi, error)
	GetTransaksiByIDPasien(ctx context.Context, PasienID string) ([]entities.Transaksi, error)
}

type transaksiService struct {
	transaksiRepo repository.TransaksiRepository
}

func NewTransaksiService(transaksiRepo repository.TransaksiRepository) TransaksiService {
	return &transaksiService{
		transaksiRepo: transaksiRepo,
	}
}

func (ts *transaksiService) NewTransaksi(ctx context.Context, transaksiDTO dto.NewTransaksi) (entities.Transaksi, error) {
	transaksi := entities.Transaksi{}
	err := smapping.FillStruct(&transaksi, smapping.MapFields(transaksiDTO))
	if err != nil {
		return entities.Transaksi{}, err
	}
	return ts.transaksiRepo.NewTransaksi(ctx, transaksi)
}
func (ts *transaksiService) GetAllTransaksi(ctx context.Context) ([]entities.Transaksi, error) {
	return ts.transaksiRepo.GetAllTransaksi(ctx)
}

func (ts *transaksiService) GetTransaksiByIDPasien(ctx context.Context, PasienID string) ([]entities.Transaksi, error) {
	return ts.transaksiRepo.GetTransaksiByIDPasien(ctx, PasienID)
}

