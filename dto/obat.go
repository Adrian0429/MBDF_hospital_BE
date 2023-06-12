package dto

type ObatCreateDTO struct {
	ID_Obat    int    `gorm:"type:int;primaryKey;not null" json:"id_obat"`
	Nama_Obat  string `gorm:"type:varchar(50);not null" json:"nama_obat"`
	Harga_Obat string `gorm:"type:money;not null" json:"harga_obat"`
	Stok       int    `gorm:"type:int;not null" json:"Stok"`
	Umum       bool   `gorm:"type:bool;not null" json:"Umum"`
}

type AllObat struct {
	Nama_Obat string `gorm:"type:varchar(50);not null" json:"nama_obat"`
}
