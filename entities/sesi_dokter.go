package entities

import (

)


type Sesi_Dokter struct {
	ID_Sesi 			int 	`gorm:"type:int;primaryKey;not null" json:"id_sesi"`
	Hari				string	`gorm:"type:varchar(10);;not null" json:"hari"`
	Jam_Masuk			int		`gorm:"type:int;;not null" json:"jam_masuk"`
	Jam_Keluar			int		`gorm:"type:int;;not null" json:"jam_keluar"`
	Maks_Jumlah_Pasien	int		`gorm:"type:int;;not null" json:"maks_Pasien"`
	Perawat_ID_Perawat	[]Perawat	`gorm:"type:varchar(6);;not null" json:"id_Perawat"`
	Ruangan_ID_Ruangan	[]Ruangan	`gorm:"type:varchar(6);;not null" json:"id_Ruangan"`
	Dokter_ID_Dokter	[]Dokter	`gorm:"type:varchar(8);;not null" json:"id_Dokter"`
}
