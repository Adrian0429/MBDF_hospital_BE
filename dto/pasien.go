package dto

import "github.com/google/uuid"

type PasienCreateDTO struct {
	Uid                 uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"uid"`
	NIK_pasien          string    `gorm:"type:char(16);primaryKey;not null" json:"nik_pasien"`
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
	Jenis_Kelamin string `json:"jenis_kelamin"`
	Tanggal_Lahir string `json:"tanggal_lahir"`
	No_Telepon    string `json:"no_telepon"`
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
