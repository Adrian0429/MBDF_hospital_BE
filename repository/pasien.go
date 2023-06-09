package repository

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PasienRepository interface {
	RegisterPasien(ctx context.Context, pasien entities.Pasien) (entities.Pasien, error)
	GetAllPasien(ctx context.Context) ([]entities.Pasien, int, error)
	GetPasienByID(ctx context.Context, PasienID uuid.UUID) (entities.Pasien, error)
	GetPasienByEmail(ctx context.Context, Email string) (entities.Pasien, error)
	UpdatePasien(ctx context.Context, pasien entities.Pasien) error
	DeletePasien(ctx context.Context, UserID uuid.UUID) error
	GetLatestReservation(ctx context.Context, NIK string) ([]dto.AmbilTransaksiTerbaru, error)
}

type pasienRepository struct {
	connection *gorm.DB
}

func NewPasienRepository(db *gorm.DB) PasienRepository {
	return &pasienRepository{
		connection: db,
	}
}

func (pr *pasienRepository) RegisterPasien(ctx context.Context, pasien entities.Pasien) (entities.Pasien, error) {
	if err := pr.connection.Table("pasiens").Create(&pasien).Error; err != nil {
		return entities.Pasien{}, err
	}
	return pasien, nil
}

func (pr *pasienRepository) GetAllPasien(ctx context.Context) ([]entities.Pasien, int, error) {
	var pasien []entities.Pasien
	var count int64

	if err := pr.connection.Table("pasiens").Find(&pasien).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return pasien, int(count), nil
}

func (pr *pasienRepository) GetPasienByID(ctx context.Context, PasienID uuid.UUID) (entities.Pasien, error) {
	var pasien entities.Pasien

	if err := pr.connection.Table("pasiens").Where("Uid = ?", PasienID).Take(&pasien).Error; err != nil {
		return entities.Pasien{}, err
	}

	return pasien, nil
}

func (pr *pasienRepository) GetPasienByEmail(ctx context.Context, Email string) (entities.Pasien, error) {
	var pasien entities.Pasien

	if err := pr.connection.Table("pasiens").Where("email = ?", Email).Take(&pasien).Error; err != nil {
		return entities.Pasien{}, err
	}
	return pasien, nil
}

func (pr *pasienRepository) UpdatePasien(ctx context.Context, pasien entities.Pasien) error {
	if err := pr.connection.Updates(&pasien).Error; err != nil {
		return err
	}
	return nil
}

func (pr *pasienRepository) DeletePasien(ctx context.Context, userID uuid.UUID) error {
	if err := pr.connection.Delete(&entities.Pasien{}, &userID).Error; err != nil {
		return err
	}
	return nil
}

func (pr *pasienRepository) GetLatestPembelianObat(ctx context.Context, PasienID string) ([]dto.LatestPembelianObatDTO, error) {
	var Pembelian []dto.LatestPembelianObatDTO

	query := `
		SELECT transaksis.Tanggal, obats.Nama_Obat, pembelian_obats.Jumlah_Obat
		FROM transaksis
		JOIN pembelian_obats ON transaksis.ID_Transaksi = pembelian_obats.transaksi_id
		JOIN obats ON pembelian_obats.obat_id = obats.ID_Obat
		WHERE transaksis.Pasien_NIK_Pasien = ?;
	`

	if err := pr.connection.Raw(query, PasienID).Scan(&Pembelian).Error; err != nil {
		return nil, err
	}

	return Pembelian, nil
}

func (sdr *pasienRepository) GetLatestReservasion(ctx context.Context, NIK string) ([]dto.AmbilTransaksiTerbaru, error) {
	var reservasiTerbaru []dto.AmbilTransaksiTerbaru
	query := "SELECT Tanggal_Reservasi, Nama_Dokter, Diagnosa_Nama_Diagnosa, Nama_Ruangan FROM Pasiens JOIN Transaksis ON NIK_Pasien = Pasien_NIK_Pasien JOIN Transaksi_Reservasis ON ID_Transaksi = Transaksi_ID_Transaksi JOIN Transaksi_Reservasi_Diagnosas ON ID_Medical_Record = Transaksi_Reservasi_ID_Medical_Record JOIN Ruangans ON Ruangan_ID_Ruangan = ID_Ruangan JOIN Sesi_Dokters ON Sesi_Dokter_ID = ID_Sesi JOIN Dokters on Dokter_ID_Dokter = ID_Dokter WHERE NIK_Pasien = ?;"
	// query := "select * from Jadwal_Dokter"
	if err := sdr.connection.Raw(query, NIK).Find(&reservasiTerbaru).Error; err != nil {
		return nil, err
	}
	return reservasiTerbaru, nil
}