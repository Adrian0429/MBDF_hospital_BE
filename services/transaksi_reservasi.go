package services

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/repository"
	"github.com/mashingan/smapping"
)

type ReservasiService interface {
	NewReservasi(ctx context.Context, transaksiReservasi dto.NewTransaksiReservasiDTO) (entities.Transaksi_Reservasi, error)
}

type reservasiService struct {
	reservasiRepo repository.ReservasiRepository
}

func NewReservasiService(rs repository.ReservasiRepository) ReservasiService {
	return &reservasiService{
		reservasiRepo: rs,
	}
}

func (tr *reservasiService) NewReservasi(ctx context.Context, transaksiReservasi dto.NewTransaksiReservasiDTO) (entities.Transaksi_Reservasi, error) {
	LatestID, err := tr.reservasiRepo.GetLatestID(ctx)
	if err != nil {
		return entities.Transaksi_Reservasi{}, nil
	}

	reservasi := entities.Transaksi_Reservasi{
		Tanggal_Reservasi:      transaksiReservasi.Tanggal_Reservasi,
		Gejala:                 "null",
		Kehadiran:              false,
		Durasi_Menginap:        0,
		Transaksi_ID_Transaksi: LatestID,
		Ruangan_ID_Ruangan:     transaksiReservasi.Ruangan_ID_Ruangan,
		Sesi_Dokter_ID:         transaksiReservasi.Sesi_Dokter_ID,
	}

	err = smapping.FillStruct(&transaksiReservasi, smapping.MapFields(transaksiReservasi))

	if err != nil {
		return entities.Transaksi_Reservasi{}, err
	}

	return tr.reservasiRepo.NewReservasi(ctx, reservasi)
}
