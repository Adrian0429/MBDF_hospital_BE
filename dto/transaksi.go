package dto

type NewTransaksi struct {
	ID_Transaksi   int    `gorm:"type:int;primaryKey;not null" json:"id_transaksi"`
	Tanggal        string `gorm:"type:date;not null" json:"Tanggal"`
	Total_Harga_DP string `gorm:"type:money;not null" json:"Total_DP"`

	Pasien_NIK_Pasien string `gorm:"type:string" json:"Pasien_NIK_Pasien"`
}
