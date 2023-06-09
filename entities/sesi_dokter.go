package entities

type Sesi_Dokter struct {
	ID_Sesi    int    `gorm:"type:int;primaryKey;not null" json:"id_sesi"`
	Hari       string `gorm:"type:varchar(10);not null" json:"hari"`
	Jam_Masuk  int    `gorm:"type:int;not null" json:"jam_masuk"`
	Jam_Keluar int    `gorm:"type:int;not null" json:"jam_keluar"`

	RuanganID string  `gorm:"type:varchar(5)" json:"-"`
	Ruangan   Ruangan `gorm:"foreignKey:RuanganID" json:"-"`

	PerawatID string  `gorm:"type:varchar(4)" json:"-"`
	Perawat   Perawat `gorm:"foreignKey:PerawatID" json:"-"`

	Dokter_ID_Dokter string `gorm:"type:char(8)" json:"-"`
	Dokter           Dokter `gorm:"foreignKey:Dokter_ID_Dokter" json:"-"`
}
