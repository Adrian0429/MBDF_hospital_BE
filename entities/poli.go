package entities

type Poli struct {
	Kode_Poli     string `gorm:"type:char(4);primaryKey;not null" json:"kode_Poli"`
	Nama_Poli     string `gorm:"type:char(50);not null" json:"nama_poli"`
	Id_Ketua_Poli string `gorm:"type:char(8);not null" json:"id_ketua_poli"`

	Dokter []Dokter `gorm:"foreignKey:Poli_Kode_Poli" json:"dokter"`
}
