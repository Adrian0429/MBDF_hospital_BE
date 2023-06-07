package entities

type Perawat struct {
	ID_Perawat    string `gorm:"type:char(4);primaryKey;not null" json:"id_perawat"`
	Nama_Perawat  string `gorm:"type:varchar(50);not null" json:"nama_perawat"`
	Jenis_Kelamin string `gorm:"type:char(1);not null" json:"jenis_kelamin"`
	Tanggal_Lahir string `gorm:"type:date;not null" json:"tanggal_lahir"`
	No_Telepon    string `gorm:"type:varchar(15);not null" json:"no_telepon"`

	Sesi_Dokter      []Sesi_Dokter      `gorm:"foreignKey:PerawatID" json:"sesi_dokter"`
	Sesi_Jaga_Nginap []Sesi_Jaga_Nginap `gorm:"foreignKey:PerawatID" json:"sesi_jaga_nginap"`
}
