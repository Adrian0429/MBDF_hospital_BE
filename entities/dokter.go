package entities

// "time"
// 	"github.com/Caknoooo/golang-clean_template/helpers"
// 	"gorm.io/gorm"

type Dokter struct {
	ID_Dokter        string `gorm:"type:char(8);primaryKey;not null" json:"ID_Dokter"`
	Nama_Dokter      string `gorm:"type:varchar(50);not null" json:"nama_dokter"`
	Jenis_Kelamin    string `gorm:"type:char(1);not null" json:"jenis_kelamin"`
	Tanggal_Lahir    string `gorm:"type:date;not null" json:"tanggal_lahir"`
	No_Telepon       string `gorm:"type:varchar(15);not null" json:"no_telepon"`
	Harga_Konsultasi string `gorm:"type:money" json:"harga_Konsul"`
	Poli_Kode_Poli   string `gorm:"type:char(4);not null" json:"poli_kode_poli"`

	Poli        Poli          `gorm:"foreignKey:Poli_Kode_Poli" json:"poli"`
	Sesi_Dokter []Sesi_Dokter `gorm:"foreignKey:Dokter_ID_Dokter" json:"sesi_dokter"`
}

// func (u *Pasien) BeforeCreate(tx *gorm.DB) error{
// 	var err error
// 	u.Password, err = helpers.HashPassword(u.Password)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
