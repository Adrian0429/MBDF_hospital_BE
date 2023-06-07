package entities

type Transaksi_Reservasi struct {
	ID_Medical_Record      int         `gorm:"type:int;primaryKey;not null" json:"id_medical_record"`
	Tanggal_Reservasi      string      `gorm:"type:date;not null" json:"Tanggal_reservasi"`
	Gejala                 string      `gorm:"type:text;not null" json:"Gejala"`
	Kehadiran              bool        `gorm:"type:boolean;not null" json:"Kehadiran"`
	Durasi_Menginap        int         `gorm:"type:int;not null" json:"Durasi_Menginap"`
	Transaksi_ID_Transaksi int         `gorm:"type:int" json:"Transaksi_ID_Transaksi"`
	Transaksi              Transaksi   `gorm:"foreignKey:Transaksi_ID_Transaksi" json:"transaksi,omitempty"`
	RuanganID              string      `gorm:"column:Ruangan_ID_Ruangan;type:char(6)" json:"ruangan_id_ruangan"`
	Ruangan                Ruangan     `gorm:"foreignKey:RuanganID" json:"ruangan,omitempty"`
	Sesi_Dokter_ID         string      `gorm:"type:string" json:"Sesi_Dokter_ID_Sesi"`
	Sesi_Dokter            Sesi_Dokter `gorm:"foreignKey:ID_Sesi"`
	Obat                   []*Obat     `gorm:"many2many:resep;" json:"obat,omitempty"`
	Diagnosa               []*Diagnosa `gorm:"many2many:transaksi_reservasi_diagnosa;" json:"diagnosa,omitempty"`
}
