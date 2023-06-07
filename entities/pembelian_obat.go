package entities

type PembelianObat struct {
	ObatID      int `gorm:"type:int;primaryKey;not null" json:"obat_id"`
	TransaksiID int `gorm:"type:int;primaryKey;not null" json:"transaksi_id"`
}
