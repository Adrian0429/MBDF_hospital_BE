package entities

type Jenis_Ruangan struct {
	ID_Jenis_Ruangan string `gorm:"char(4):int;primaryKey;not null" json:"id_jenis_ruangan"`
	Jenis_Ruangan    string `gorm:"type:varchar(30);not null" json:"jenis_ruangan"`

	Ruangan []Ruangan `gorm:"foreignKey:ID_jenis_ruangan" json:"ruangan,omitempty"`
}
