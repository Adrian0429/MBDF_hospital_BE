package entities

import (
	// "time"
// 	"github.com/Caknoooo/golang-clean_template/helpers"
// 	"gorm.io/gorm"
)


type Pasien struct {
	NIK_pasien             string    `gorm:"type:char(16);primaryKey;not null" json:"nik_pasien"`
	Nama_Pasien          string    `gorm:"type:varchar(50);not null" json:"nama"`
	Jenis_Kelamin   string    `gorm:"type:char(1);not null" json:"jenis_kelamin"`
	Tanggal_Lahir   string    `gorm:"type:date;not null" json:"tanggal_lahir"`
	No_Telepon      string    `gorm:"type:varchar(15);not null" json:"no_telepon"`
	Email    string    `gorm:"type:varchar(100)" json:"email"`
	Password string    `gorm:"type:varchar(100)" json:"password"`
	Tanggal_Daftar_Akun string `gorm:"type:date" json:"tanggal_daftar_akun"`
}

// func (u *Pasien) BeforeCreate(tx *gorm.DB) error{	
// 	var err error
// 	u.Password, err = helpers.HashPassword(u.Password)
// 	if err != nil {
// 		return err	
// 	}
// 	return nil
// }