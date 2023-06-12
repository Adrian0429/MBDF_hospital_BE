package dto

import "github.com/google/uuid"

type PasienCreateDTO struct {
	NIK_pasien          string    `gorm:"type:char(16);primaryKey;not null" json:"nik_pasien"`
	Uid                 uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"uid"`
	Nama_Pasien         string    `gorm:"type:varchar(50);not null" json:"nama"`
	Jenis_Kelamin       string    `gorm:"type:char(1);not null" json:"jenis_kelamin"`
	Tanggal_Lahir       string    `gorm:"type:date;not null" json:"tanggal_lahir"`
	No_Telepon          string    `gorm:"type:varchar(15);not null" json:"no_telepon"`
	Email               string    `gorm:"type:varchar(100)" json:"email"`
	Password            string    `gorm:"type:varchar(100)" json:"password"`
	Tanggal_Daftar_Akun string    `gorm:"type:date" json:"tanggal_daftar_akun"`
}

type GetAllPasienDTO struct {
	Count         int    `json:"count"`
	NIK_pasien    string `json:"nik_pasien"`
	Nama_Pasien   string `json:"nama"`
	Tanggal_Lahir string `json:"tanggal_lahir"`
}

type UpdatePasienDTO struct {
	Uid         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"uid"`
	NIK_pasien  string    `gorm:"type:char(16);primaryKey;not null" json:"nik_pasien"`
	Nama_Pasien string    `gorm:"type:varchar(50);not null" json:"nama"`
	No_Telepon  string    `gorm:"type:varchar(15);not null" json:"no_telepon"`
	Email       string    `gorm:"type:varchar(100)" json:"email"`
	Password    string    `gorm:"type:varchar(100)" json:"password"`
}

type PasienLoginDTO struct {
	Email    string `gorm:"type:varchar(100)" json:"email"`
	Password string `gorm:"type:varchar(100)" json:"password"`
}

type LatestPembelianObatDTO struct {
	Tanggal     string `gorm:"type:date;not null" json:"Tanggal"`
	Nama_Obat   string `gorm:"type:varchar(50);not null" json:"nama_obat"`
	Jumlah_Obat int    `gorm:"type:int;not null" json:"jumlah_obat"`
}

type Transaksi_PasienDTO struct {
	ID_Transaksi   int    `gorm:"type:int;primaryKey;not null" json:"id_transaksi"`
	Tanggal        string `gorm:"type:date;not null" json:"Tanggal"`
	Total_Harga_DP string `gorm:"type:money;not null" json:"Total_DP"`
	NIK_pasien     string `gorm:"type:char(16);primaryKey;not null" json:"nik_pasien"`
	Nama_Pasien    string `gorm:"type:varchar(50);not null" json:"nama"`
	Jenis_Kelamin  string `gorm:"type:char(1);not null" json:"jenis_kelamin"`
	No_Telepon     string `gorm:"type:varchar(15);not null" json:"no_telepon"`
}

type Jadwal_Dokter_UserDTO struct {
	Nama_Dokter      string `gorm:"type:varchar(50);not null" json:"nama_dokter"`
	Jam_Masuk        int    `gorm:"type:int;not null" json:"jam_masuk"`
	Jam_Keluar       int    `gorm:"type:int;not null" json:"jam_keluar"`
	Jenis_Kelamin    string `gorm:"type:char(1);not null" json:"jenis_kelamin"`
	Harga_Konsultasi string `gorm:"type:money" json:"harga_Konsul"`
}
