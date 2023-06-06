package entities

import (

)


type Tranksaksi_Reservasi_Diagnosa struct {
	Transaksi_Reservasi_ID_Medical_Record	[]Transaksi_Reservasi		`gorm:"type:int;primaryKey;not null" json:"transaksi_Id_Medical_record"`
	Diagnosa_Nama_Diagnosa					[]Diagnosa	`gorm:"type:varchar(50);not null" json:"nama_Diagnosa"`
}
