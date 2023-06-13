package dto

type NewTransaksiReservasiDTO struct {
	ID_Medical_Record int    `gorm:"type:int;primaryKey;not null" json:"id_medical_record"`
	Tanggal_Reservasi string `gorm:"type:date;not null" json:"Tanggal_reservasi"`
	Gejala            string `gorm:"type:text;" json:"Gejala"`
	Kehadiran         bool   `gorm:"type:boolean;not null" json:"Kehadiran"`
	Durasi_Menginap   int    `gorm:"type:int;not null" json:"Durasi_Menginap"`
	Transaksi_ID_Transaksi int       `gorm:"type:int" json:"Transaksi_ID_Transaksi"`
	Ruangan_ID_Ruangan string  `gorm:"type:char(6)" json:"Ruangan_ID_Ruangan"`
	Sesi_Dokter_ID string `gorm:"type:string" json:"Sesi_Dokter_ID_Sesi"`
}
