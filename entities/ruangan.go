package entities

type Ruangan struct {
	ID_Ruangan          string                `gorm:"type:char(6);primaryKey;not null" json:"id_ruangan"`
	Nama_Ruangan        string                `gorm:"type:varchar(30);not null" json:"nama_ruangan"`
	Kelas               string                `gorm:"type:varchar(20);not null" json:"kelas"`
	Harga_Ruangan       string                `gorm:"type:money;not null" json:"harga_ruangan"`
	ID_jenis_ruangan    string                `gorm:"type:varchar(2)" json:"id_jenis_ruangan" binding:"required"`
	Jenis_Ruangan       Jenis_Ruangan         `gorm:"foreignKey:ID_jenis_ruangan" json:"jenis_ruangan"`
	Kapasitas           int                   `gorm:"type:int;not null" json:"kapasitas"`
	Sesi_Jaga_Nginap    []Sesi_Jaga_Nginap    `gorm:"foreignKey:RuanganID" json:"sesi_jaga_nginap"`
	Sesi_Dokter         []Sesi_Dokter         `gorm:"foreignKey:RuanganID" json:"sesi_dokter"`
	Transaksi_Reservasi []Transaksi_Reservasi `gorm:"foreignKey:RuanganID" json:"transaksi_reservasi"`
}
