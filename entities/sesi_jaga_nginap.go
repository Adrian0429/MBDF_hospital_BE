package entities

import (

)


type Sesi_Jaga_Nginap struct {
	ID_Sesi 			int 	`gorm:"type:int;primaryKey;not null" json:"id_sesi"`
	Hari				string	`gorm:"type:varchar(10);;not null" json:"hari"`
	Jam_Masuk			int		`gorm:"type:int;;not null" json:"jam_masuk"`
	Jam_Keluar			int		`gorm:"type:int;;not null" json:"jam_keluar"`
	Maks_Jumlah_Pasien	int		`gorm:"type:int;;not null" json:"maks_Pasien"`
	Ruangan_ID_Ruangan 	Ruangan	
	Perawat_ID_Perawat	Perawat	
}
