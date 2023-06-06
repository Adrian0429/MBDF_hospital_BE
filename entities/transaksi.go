package entities

type Transaksi struct {
	ID_Transaksi      int    `gorm:"type:int;primaryKey;not null" json:"id_transaksi"`
	Tanggal           string `gorm:"type:date;;not null" json:"Tanggal"`
	Total_Harga_DP    string `gorm:"type:int;;not null" json:"Total_DP"`
	Pasien_NIK_Pasien string `gorm:"type:varchar(16);;not null" json:"jam_keluar"`
}
