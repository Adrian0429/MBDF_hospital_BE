package entities

type Sesi_Jaga_Nginap struct {
	ID_Sesi    int    `gorm:"type:int;primaryKey;not null" json:"id_sesi"`
	Hari       string `gorm:"type:varchar(10);not null" json:"hari"`
	Jam_Masuk  int    `gorm:"type:int;not null" json:"jam_masuk"`
	Jam_Keluar int    `gorm:"type:int;not null" json:"jam_keluar"`

	RuanganID string  `gorm:"type:varchar(5)" json:"Ruangan_ID_Ruangan" binding:"required"`
	Ruangan   Ruangan `gorm:"foreignKey:RuanganID" json:"ruangan"`

	PerawatID string  `gorm:"type:varchar(4)" json:"Perawat_ID_Perawat" binding:"required"`
	Perawat   Perawat `gorm:"foreignKey:PerawatID" json:"perawat"`
}
