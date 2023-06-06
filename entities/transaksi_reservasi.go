package entities

import (

)


type Transaksi_Reservasi struct {
	ID_Medical_Record 			int 	`gorm:"type:int;primaryKey;not null" json:"id_medical_record"`
	Tanggal_Reservasi			string	`gorm:"type:date;not null" json:"Tanggal_reservasi"`
	Gejala						string	`gorm:"type:text;not null" json:"Gejala"`
	Kehadiran					bool	`gorm:"type:boolean;not null" json:"Kehadiran"`
	Durasi_Menginap				int		`gorm:"type:int;not null" json:"Durasi_Menginap"`
	
	Sesi_Dokter_ID_Sesi			[]Sesi_Dokter		`gorm:"type:int;not null" json:"id_Sesi_Dokter"`
	Transaksi_ID_Transaksi		[]Transaksi		`gorm:"type:int;not null" json:"id_Transaksi"`
	Ruangan_ID_Ruangan			[]Ruangan	`gorm:"type:char(5);not null" json:"id_Ruangan"`
}


