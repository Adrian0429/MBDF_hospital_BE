package dto

type DokterCreateDTO struct {
	Nama_Dokter      string `gorm:"type:varchar(50);not null" json:"nama_dokter"`
	Jenis_Kelamin    string `gorm:"type:char(1);not null" json:"jenis_kelamin"`
	Tanggal_Lahir    string `gorm:"type:date;not null" json:"tanggal_lahir"`
	No_Telepon       string `gorm:"type:varchar(15);not null" json:"no_telepon"`
	Harga_Konsultasi string `gorm:"type:money" json:"harga_Konsul"`
	Poli_Kode_Poli   string `gorm:"type:char(4);not null" json:"poli_kode_poli"`
}

type AllDokterUserDTO struct {
	ID_Dokter        string `gorm:"type:char(8);primaryKey;not null" json:"ID_Dokter"`
	Nama_Dokter      string `gorm:"type:varchar(50);not null" json:"nama_dokter"`
	Jenis_Kelamin    string `gorm:"type:char(1);not null" json:"jenis_kelamin"`
	No_Telepon       string `gorm:"type:varchar(15);not null" json:"no_telepon"`
	Harga_Konsultasi string `gorm:"type:money" json:"harga_Konsul"`
	Poli_Kode_Poli   string `gorm:"type:char(4);not null" json:"poli_kode_poli"`
	Nama_Poli        string `gorm:"type:char(50);not null" json:"nama_poli"`
}

type AllDokterDTO struct {
	ID_Dokter     string `gorm:"type:char(8);primaryKey;not null" json:"ID_Dokter"`
	Nama_Dokter   string `gorm:"type:varchar(50);not null" json:"nama_dokter"`
	Tanggal_Lahir string `gorm:"type:date;not null" json:"tanggal_lahir"`
	Nama_Poli     string `gorm:"type:char(50);not null" json:"nama_poli"`
}

type UpdateDokterDTO struct {
	ID_Dokter        string `gorm:"type:char(8);primaryKey;not null" json:"ID_Dokter"`
	Nama_Dokter      string `gorm:"type:varchar(50);not null" json:"nama_dokter"`
	No_Telepon       string `gorm:"type:varchar(15);not null" json:"no_telepon"`
	Harga_Konsultasi string `gorm:"type:money" json:"harga_Konsul"`
}

type Jadwal_Dokter_AdminDTO struct {
	Nama_Dokter   string `gorm:"type:varchar(50);not null" json:"nama_dokter"`
	Jam_Masuk     int    `gorm:"type:int;not null" json:"jam_masuk"`
	Jam_Keluar    int    `gorm:"type:int;not null" json:"jam_keluar"`
	Jenis_Kelamin string `gorm:"type:char(1);not null" json:"jenis_kelamin"`
	Id_Ruangan    string `gorm:"type:varchar(6)" json:"id_ruangan" binding:"required"`
}

type PoliDTO struct {
	Kode_Poli string `gorm:"type:char(4);not null" json:"poli_kode_poli"`
}
