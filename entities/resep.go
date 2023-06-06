package entities

type Resep struct {
	ObatID               int `gorm:"type:int;primaryKey;not null" json:"obat_id"`
	TransaksiReservasiID int `gorm:"type:int;primaryKey;not null" json:"transaksi_reservasi_id"`
}
