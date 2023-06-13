package repository

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/entities"
	"gorm.io/gorm"
)

type ReservasiRepository interface {
	NewReservasi(ctx context.Context, transaksiReservasi entities.Transaksi_Reservasi) (entities.Transaksi_Reservasi, error)
	GetLatestID(ctx context.Context) (int, error)
}

type reservasiRepository struct {
	connection *gorm.DB
}

func NewReservasiRepository(db *gorm.DB) ReservasiRepository {
	return &reservasiRepository{
		connection: db,
	}
}

func (tr *reservasiRepository) NewReservasi(ctx context.Context, transaksiReservasi entities.Transaksi_Reservasi) (entities.Transaksi_Reservasi, error) {
	if err := tr.connection.Table("transaksi_reservasis").Create(&transaksiReservasi).Error; err != nil {
		return entities.Transaksi_Reservasi{}, err
	}
	return transaksiReservasi, nil
}

func (tr *reservasiRepository) GetLatestID(ctx context.Context) (int, error) {
	var ID entities.Transaksi
	if err := tr.connection.Table("transaksis").Last(&ID).Error; err != nil {
		return 0, nil
	}
	return ID.ID_Transaksi, nil
}
