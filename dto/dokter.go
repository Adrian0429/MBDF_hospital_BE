package dto

type DokterCreateDTO struct {
	ID_Dokter        string `gorm:"type:char(8);primaryKey;not null" json:"ID_Dokter"`
	Nama_Dokter      string `gorm:"type:varchar(50);not null" json:"nama_dokter"`
	Jenis_Kelamin    string `gorm:"type:char(1);not null" json:"jenis_kelamin"`
	Tanggal_Lahir    string `gorm:"type:date;not null" json:"tanggal_lahir"`
	No_Telepon       string `gorm:"type:varchar(15);not null" json:"no_telepon"`
	Email            string `gorm:"type:varchar(100)" json:"email"`
	Password         string `gorm:"type:varchar(100)" json:"password"`
	Harga_Konsultasi string `gorm:"type:money" json:"harga_Konsul"`
	Poli_Kode_Poli   string `gorm:"type:char(4);not null" json:"poli_kode_poli"`
}
