package entities

import (

)


type Jenis_Ruangan struct {
	ID_Jenis_Ruangan int `gorm:"type:int;primaryKey;not null" json:"id_jenis_ruangan"`
	Jenis_Ruangan string `gorm:"type:varchar(30);not null" json:"jenis_ruangan"`
}
