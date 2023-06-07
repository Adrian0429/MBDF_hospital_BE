package entities

type Obat struct {
	ID_Obat             int                    `gorm:"type:int;primaryKey;not null" json:"id_obat"`
	Nama_Obat           string                 `gorm:"type:varchar(50);not null" json:"nama_obat"`
	Harga_Obat          string                 `gorm:"type:money;not null" json:"harga_obat"`
	Stok                int                    `gorm:"type:int;not null" json:"Stok"`
	Umum                bool                   `gorm:"type:bool;not null" json:"Umum"`
	Transaksi           []*Transaksi           `gorm:"many2many:pembelian_obat;" json:"transaksi,omitempty"`
	Transaksi_Reservasi []*Transaksi_Reservasi `gorm:"many2many:resep;" json:"transaksi_reservasi,omitempty"`
}
