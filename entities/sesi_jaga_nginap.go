package entities

type Sesi_Jaga_Nginap struct {
	ID_Sesi            int    `gorm:"type:int;primaryKey;not null" json:"id_sesi"`
	Hari               string `gorm:"type:varchar(10);not null" json:"hari"`
	Jam_Masuk          int    `gorm:"type:int;not null" json:"jam_masuk"`
	Jam_Keluar         int    `gorm:"type:int;not null" json:"jam_keluar"`
	Maks_Jumlah_Pasien int    `gorm:"type:int;not null" json:"maks_Pasien"`

	RuanganID string  `gorm:"column:Ruangan_ID_Ruangan;type:varchar(5)" json:"id_ruangan" binding:"required"`
	Ruangan   Ruangan `gorm:"foreignKey:RuanganID" json:"ruangan"`

	PerawatID string  `gorm:"column:Perawat_ID_Perawat;type:varchar(4)" json:"id_perawat" binding:"required"`
	Perawat   Perawat `gorm:"foreignKey:PerawatID" json:"perawat"`
}
