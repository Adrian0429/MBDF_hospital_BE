package repository

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"gorm.io/gorm"
)

type PerawatRepository interface {
	RegisterPerawat(ctx context.Context, perawat entities.Perawat) (entities.Perawat, error)
	GetAllPerawat(ctx context.Context) ([]dto.GetAllPerawat, error)
	GetPerawatByID(ctx context.Context, PerawatID string) (entities.Perawat, error)
	GetJadwalPerawat(ctx context.Context) ([]dto.JadwalPerawatDTO, error)
}

type perawatRepository struct {
	connection *gorm.DB
}

func NewPerawatRepository(db *gorm.DB) PerawatRepository {
	return &perawatRepository{
		connection: db,
	}
}

func (dr *perawatRepository) RegisterPerawat(ctx context.Context, perawat entities.Perawat) (entities.Perawat, error) {
	if err := dr.connection.Table("perawats").Create(&perawat).Error; err != nil {
		return entities.Perawat{}, err
	}
	return perawat, nil
}

func (dr *perawatRepository) GetAllPerawat(ctx context.Context) ([]dto.GetAllPerawat, error) {
	var perawat []dto.GetAllPerawat
	if err := dr.connection.Raw("select * from view_perawat").Scan(&perawat).Error; err != nil {
		return nil, err
	}
	return perawat, nil
}

func (dr *perawatRepository) GetJadwalPerawat(ctx context.Context) ([]dto.JadwalPerawatDTO, error) {
	var jadwal_perawat []dto.JadwalPerawatDTO

	// Create the view separately
	viewCreationQuery := "CREATE OR REPLACE VIEW view_jadwal_perawat AS SELECT hari, (SELECT nama_perawat FROM Perawats WHERE perawats.id_perawat = sj.perawat_id), jam_masuk, jam_keluar, (SELECT nama_ruangan FROM ruangans WHERE ruangans.id_ruangan = sj.ruangan_id) FROM sesi_jaga_nginaps sj"
	if err := dr.connection.Exec(viewCreationQuery).Error; err != nil {
		return nil, err
	}

	// Perform the SELECT query on the view
	selectQuery := "SELECT * FROM view_jadwal_perawat"
	if err := dr.connection.Raw(selectQuery).Scan(&jadwal_perawat).Error; err != nil {
		return nil, err
	}

	return jadwal_perawat, nil
}

func (dr *perawatRepository) GetPerawatByID(ctx context.Context, PerawatID string) (entities.Perawat, error) {
	var perawat entities.Perawat
	if err := dr.connection.Table("perawats").Where("ID_Perawat = ?", PerawatID).First(&perawat).Error; err != nil {
		return entities.Perawat{}, err
	}
	return perawat, nil
}


