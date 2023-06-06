package services

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/repository"
)

type TransaksiService interface {
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

func (ts *transaksiService) GetAllTransaksi(ctx context.Context) ([]entities.Transaksi, error) {
	return ts.transaksiRepo.GetAllTransaksi(ctx)
}

func (ts *transaksiService) GetTransaksiByIDPasien(ctx context.Context, PasienID string) ([]entities.Transaksi, error) {
	return ts.transaksiRepo.GetTransaksiByIDPasien(ctx, PasienID)
}
