package entities

type PembelianObat struct {
	ObatID      int `gorm:"type:int;primaryKey;not null" json:"obat_id"`
	TransaksiID int `gorm:"type:int;primaryKey;not null" json:"transaksi_id"`
	Jumlah_Obat int `gorm:"type:int;not null" json:"jumlah_obat"`
}
