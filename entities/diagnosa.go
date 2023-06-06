package entities

type Diagnosa struct {
	Nama_Diagnosa       string                 `gorm:"type:char(16);primaryKey;not null" json:"nik_pasien"`
	Transaksi_Reservasi []*Transaksi_Reservasi `gorm:"many2many:transaksi_reservasi_diagnosa;" json:"transaksi_reservasi,omitempty"`
}
