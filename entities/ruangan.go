package entities

type Ruangan struct {
	ID_Ruangan          string                `gorm:"type:char(6);primaryKey;not null" json:"id_ruangan"`
	Nama_Ruangan        string                `gorm:"type:varchar(40);not null" json:"Nama_Ruangan"`
	Kelas               string                `gorm:"type:varchar(20);" json:"kelas"`
	Harga_Ruangan       string                `gorm:"type:money;" json:"harga_ruangan"`
	ID_jenis_ruangan    string                `gorm:"column:id_jenis_ruangan;type:varchar(2)" json:"Jenis_Ruangan_ID_Jenis_Ruangan" binding:"required"`
	Jenis_Ruangan       Jenis_Ruangan         `gorm:"foreignKey:ID_jenis_ruangan" json:"jenis_ruangan"`
	Kapasitas           int                   `gorm:"type:int;not null" json:"kapasitas"`
	Sesi_Jaga_Nginap    []Sesi_Jaga_Nginap    `gorm:"foreignKey:RuanganID" json:"sesi_jaga_nginap"`
	Sesi_Dokter         []Sesi_Dokter         `gorm:"foreignKey:RuanganID" json:"sesi_dokter"`
	Transaksi_Reservasi []Transaksi_Reservasi `gorm:"foreignKey:Ruangan_ID_Ruangan" json:"transaksi_reservasi"`
}
