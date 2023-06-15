package dto

type PerawatCreateDTO struct {
	ID_Perawat    string `gorm:"type:char(4);primaryKey;not null" json:"id_perawat"`
	Nama_Perawat  string `gorm:"type:varchar(50);not null" json:"nama_perawat"`
	Jenis_Kelamin string `gorm:"type:char(1);not null" json:"jenis_kelamin"`
	Tanggal_Lahir string `gorm:"type:date;not null" json:"tanggal_lahir"`
	No_Telepon    string `gorm:"type:varchar(15);not null" json:"no_telepon"`
	Email         string `gorm:"type:varchar(100)" json:"email"`
	Password      string `gorm:"type:varchar(100)" json:"password"`
}

type GetAllPerawat struct {
	ID_Perawat    string `gorm:"type:char(4);primaryKey;not null" json:"id_perawat"`
	Nama_Perawat  string `gorm:"type:varchar(50);not null" json:"nama_perawat"`
	Tanggal_Lahir string `gorm:"type:date;not null" json:"tanggal_lahir"`
}

type JadwalPerawatDTO struct {
	Hari         string `gorm:"type:varchar(10);not null" json:"hari"`
	Nama_Perawat string `gorm:"type:varchar(50);not null" json:"nama_perawat"`
	Jam_Masuk    int    `gorm:"type:int;not null" json:"jam_masuk"`
	Jam_Keluar   int    `gorm:"type:int;not null" json:"jam_keluar"`
	Nama_Ruangan string `gorm:"not null" json:"Nama_Ruangan"`
}
