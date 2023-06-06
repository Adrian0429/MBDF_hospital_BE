package entities

import (

)


type Resep struct {
	Obat_ID_Obat							[]Obat	`gorm:"type:int;primaryKey;not null" json:"id_obat"`
	Transaksi_Reservasi_ID_Medical_Record	[]Transaksi_Reservasi `gorm:"type:int;;not null" json:"transaksi_ID_Medical_Record"`
}
