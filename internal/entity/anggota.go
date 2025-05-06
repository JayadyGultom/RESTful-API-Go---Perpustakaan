package entity

import "time"

type Anggota struct {
    IDAnggota        int       `gorm:"primaryKey;column:id_anggota" json:"id_anggota"`
    Nama             string    `gorm:"type:varchar(255)" json:"nama"`
    Alamat           string    `gorm:"type:text" json:"alamat"`
    NoTelp           string    `gorm:"type:varchar(20);column:no_telp" json:"no_telp"`
    TanggalBergabung time.Time `gorm:"type:date;column:tanggal_bergabung" json:"tanggal_bergabung"`
    Peminjaman       []Peminjaman `gorm:"foreignKey:IDAnggota;references:IDAnggota" json:"-"`
}

func (Anggota) TableName() string {
    return "anggota"
}
