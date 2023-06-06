package entities

import (

)

type Obat struct {
	ID_Obat 	int 		`gorm:"type:int;primaryKey;not null" json:"id_obat"`
	Nama_Obat 	string 		`gorm:"type:varchar(50);not null" json:"nama_obat"`
	Harga_Obat 	string 		`gorm:"type:int;not null" json:"harga_obat"`
	Stok		int
	Umum		bool
}
