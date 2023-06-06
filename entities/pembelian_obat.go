package entities

import (

)


type Pembelian_Obat struct {
	Obat_ID_Obat 			[]Obat `gorm:"type:int;primaryKey;not null" json:"id_obat"`
	Transaksi_ID_Transaksi	[]Transaksi	`gorm:"type:int;;not null" json:"id_transaksi"`
}
