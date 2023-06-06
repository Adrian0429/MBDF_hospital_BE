package entities

type Transaksi_Reservasi_Diagnosa struct {
	Transaksi_Reservasi_ID_Medical_Record int    `gorm:"type:int;primaryKey;not null" json:"transaksi_Id_Medical_record"`
	Diagnosa_Nama_Diagnosa                string `gorm:"type:varchar(50);primaryKey;not null" json:"nama_Diagnosa"`
}
