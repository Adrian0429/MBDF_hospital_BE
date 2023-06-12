package dto

type NewTransaksiReservasiDTO struct {
	Tanggal_Reservasi  string `gorm:"type:date;not null" json:"Tanggal_Reservasi"`
	Ruangan_ID_Ruangan string `gorm:"type:char(6)" json:"Ruangan_ID_Ruangan"`
	Sesi_Dokter_ID     string `gorm:"type:string" json:"Sesi_Dokter_ID_Sesi"`
}
