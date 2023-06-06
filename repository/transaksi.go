package repository

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/entities"
	"gorm.io/gorm"
)

type TransaksiRepository interface {
	NewTransaksi(ctx context.Context, Transaksi entities.Transaksi) (entities.Transaksi, error)
	GetAllTransaksi(ctx context.Context) ([]entities.Transaksi, error)
	GetTransaksiByIDPasien(ctx context.Context, PasienID string) ([]entities.Transaksi, error)
}

type transaksiRepository struct {
	connection *gorm.DB
}

func NewTransaksiRepository(db *gorm.DB) TransaksiRepository {
	return &transaksiRepository{
		connection: db,
	}
}

func (tr *transaksiRepository) NewTransaksi(ctx context.Context, transaksi entities.Transaksi) (entities.Transaksi, error) {
	if err := tr.connection.Create(&transaksi).Error; err != nil {
		return entities.Transaksi{}, err
	}
	return transaksi, nil
}

func (tr *transaksiRepository) GetAllTransaksi(ctx context.Context) ([]entities.Transaksi, error) {
	var transaksi []entities.Transaksi
	query := `SELECT * FROM Transaksi`
	if err := tr.connection.Raw(query).Scan(&transaksi).Error; err != nil {
		return nil, err
	}
	return transaksi, nil
}

func (tr *transaksiRepository) GetTransaksiByIDPasien(ctx context.Context, PasienID string) ([]entities.Transaksi, error) {
	var transaksi []entities.Transaksi
	query := `SELECT ID_Transaksi, Tanggal, Total_Harga_DP, Pasien_NIK_Pasien FROM Transaksi WHERE Pasien_NIK_Pasien = ? ORDER BY ID_Transaksi`
	if err := tr.connection.Raw(query, PasienID).Scan(&transaksi).Error; err != nil {
		return nil, err
	}
	return transaksi, nil
}
