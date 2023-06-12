package entities

type Transaksi struct {
	ID_Transaksi   int    `gorm:"type:int;primaryKey;not null" json:"id_transaksi"`
	Tanggal        string `gorm:"type:date;not null" json:"Tanggal"`
	Total_Harga_DP string `gorm:"type:money;not null" json:"Total_DP"`

	Pasien_NIK_Pasien string `gorm:"type:string" json:"Pasien_NIK_Pasien"`
	Pasien            Pasien `gorm:"foreignKey:Pasien_NIK_Pasien" json:"user,omitempty"`

	Transaksi_Reservasi []Transaksi_Reservasi `gorm:"foreignKey:Transaksi_ID_Transaksi" json:"transaksi_reservasi,omitempty"`

}
