package dto

type SesiDokterCreateDTO struct {
	ID_Sesi          int    `gorm:"type:int;primaryKey;not null" json:"id_sesi"`
	Hari             string `gorm:"type:varchar(10);;not null" json:"hari"`
	Jam_Masuk        int    `gorm:"type:int;not null" json:"jam_masuk"`
	Jam_Keluar       int    `gorm:"type:int;not null" json:"jam_keluar"`
	RuanganID        string `gorm:"type:varchar(5)" json:"Ruangan_ID_Ruangan" binding:"required"`
	PerawatID        string `gorm:"type:varchar(4)" json:"Perawat_ID_Perawat" binding:"required"`
	Dokter_ID_Dokter string `gorm:"type:char(8)" json:"Dokter_ID_Dokter"`
}

type AmbilSesiDTO struct {
	ID_Dokter        string `gorm:"type:char(5);primaryKey;not null" json:"ID_Dokter"`
	Nama_Dokter      string `gorm:"type:varchar(50);not null" json:"nama_dokter"`
	Hari             string `gorm:"type:varchar(10);;not null" json:"hari"`
	Jam_Masuk        int    `gorm:"type:int;not null" json:"jam_masuk"`
	Jam_Keluar       int    `gorm:"type:int;not null" json:"jam_keluar"`
	Jenis_Kelamin    string `gorm:"type:char(1);not null" json:"jenis_kelamin"`
	No_Telepon       string `gorm:"type:varchar(15);not null" json:"no_telepon"`
	Harga_Konsultasi string `gorm:"type:money" json:"harga_Konsul"`
	RuanganID        string `gorm:"type:varchar(5)" json:"Ruangan_ID_Ruangan" binding:"required"`
}

type AmbilTransaksiTerbaru struct {
	Tanggal_Reservasi      string `gorm:"type:date;not null" json:"Tanggal_reservasi"`
	Nama_Dokter            string `gorm:"type:varchar(50);not null" json:"Nama_dokter"`
	Diagnosa_Nama_Diagnosa string `gorm:"type:varchar(50);primaryKey;not null" json:"Nama_diagnosa"`
	Nama_Ruangan           string `gorm:"type:varchar(40);not null" json:"Nama_Ruangan"`
}
