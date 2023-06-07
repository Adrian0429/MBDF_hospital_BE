package entities

type Sesi_Dokter struct {
	ID_Sesi            int    `gorm:"type:int;primaryKey;not null" json:"id_sesi"`
	Hari               string `gorm:"type:varchar(10);;not null" json:"hari"`
	Jam_Masuk          int    `gorm:"type:int;;not null" json:"jam_masuk"`
	Jam_Keluar         int    `gorm:"type:int;;not null" json:"jam_keluar"`
	Maks_Jumlah_Pasien int    `gorm:"type:int;;not null" json:"maks_Pasien"`

	RuanganID string  `gorm:"type:varchar(5)" json:"ruangan_id" binding:"required"`
	Ruangan   Ruangan `gorm:"foreignKey:RuanganID" json:"ruangan"`

	PerawatID string  `gorm:"type:varchar(4)" json:"id_perawat" binding:"required"`
	Perawat   Perawat `gorm:"foreignKey:PerawatID" json:"perawat"`

	Dokter_ID_Dokter string `gorm:"type:char(8)" json:"dokter_id_dokter"`
	Dokter           Dokter `gorm:"foreignKey:Dokter_ID_Dokter" json:"dokter"`
}
