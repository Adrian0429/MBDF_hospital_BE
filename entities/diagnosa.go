package entities

import (

)


type Diagnosa struct {
	Nama_Diagnosa string `gorm:"type:char(16);primaryKey;not null" json:"nik_pasien"`
}
