package repository

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"gorm.io/gorm"
)

type DokterRepository interface {
	RegisterDokter(ctx context.Context, dokter entities.Dokter) (entities.Dokter, error)
	GetAllDokter(ctx context.Context) ([]dto.AllDokterDTO, error)
	AllDokterUser(ctx context.Context) ([]dto.AllDokterUserDTO, error)
	GetDokterByID(ctx context.Context, DokterID string) (entities.Dokter, error)
	UpdateDoctor(ctx context.Context, dokter entities.Dokter) error
	Jadwal_Dokter_Admin(ctx context.Context) ([]dto.Jadwal_Dokter_AdminDTO, error)
}

type dokterRepository struct {
	connection *gorm.DB
}

func NewDokterRepository(db *gorm.DB) DokterRepository {
	return &dokterRepository{
		connection: db,
	}
}

func (dr *dokterRepository) RegisterDokter(ctx context.Context, dokter entities.Dokter) (entities.Dokter, error) {
	if err := dr.connection.Table("dokters").Create(&dokter).Error; err != nil {
		return entities.Dokter{}, err
	}
	return dokter, nil
}

func (dr *dokterRepository) GetAllDokter(ctx context.Context) ([]dto.AllDokterDTO, error) {
	var dokter []dto.AllDokterDTO

	CreateView := `CREATE OR REPLACE VIEW view_dokter AS SELECT id_dokter, Nama_dokter, tanggal_lahir, (select nama_poli from polis where polis.kode_poli = dokters.poli_kode_poli) FROM dokters;`
	if err := dr.connection.Raw(CreateView).Scan(&dokter).Error; err != nil {
		return nil, err
	}

	query := `select * from view_dokter`
	if err := dr.connection.Raw(query).Scan(&dokter).Error; err != nil {
		return nil, err
	}
	return dokter, nil
}

func (dr *dokterRepository) AllDokterUser(ctx context.Context) ([]dto.AllDokterUserDTO, error) {
	var dokter []dto.AllDokterUserDTO

	viewCreationQuery := `CREATE OR REPLACE VIEW Dokter_Poli AS 
SELECT ID_Dokter, Nama_Dokter, Jenis_Kelamin, No_Telepon, Harga_Konsultasi, Kode_Poli, Nama_Poli 
FROM Dokters JOIN Polis ON Poli_Kode_Poli = Kode_Poli;`

	if err := dr.connection.Exec(viewCreationQuery).Error; err != nil {
		return nil, err
	}

	query := `select * from Dokter_Poli`
	if err := dr.connection.Raw(query).Scan(&dokter).Error; err != nil {
		return nil, err
	}
	return dokter, nil
}

func (dr *dokterRepository) UpdateDoctor(ctx context.Context, dokter entities.Dokter) error {
	if err := dr.connection.Updates(&dokter).Error; err != nil {
		return err
	}
	return nil
}

func (dr *dokterRepository) GetDokterByID(ctx context.Context, DokterID string) (entities.Dokter, error) {
	var dokter entities.Dokter
	if err := dr.connection.Table("dokters").Where("ID_Dokter = ?", DokterID).First(&dokter).Error; err != nil {
		return entities.Dokter{}, err
	}
	return dokter, nil
}

func (dr *dokterRepository) Jadwal_Dokter_Admin(ctx context.Context) ([]dto.Jadwal_Dokter_AdminDTO, error) {
	var jadwal []dto.Jadwal_Dokter_AdminDTO

	viewCreationQuery := `CREATE OR REPLACE VIEW Jadwal_Dokter_Admin AS 
SELECT Nama_Dokter, Jam_Masuk, Jam_Keluar, Jenis_Kelamin, Id_ruangan
FROM Dokters JOIN Sesi_Dokters ON ID_Dokter = Dokter_ID_Dokter	
JOIN Ruangans ON Sesi_Dokters.Ruangan_Id = Id_ruangan;
`

	if err := dr.connection.Exec(viewCreationQuery).Error; err != nil {
		return nil, err
	}

	query := `select * from Jadwal_Dokter_Admin`
	if err := dr.connection.Raw(query).Scan(&jadwal).Error; err != nil {
		return nil, err
	}
	return jadwal, nil
}
