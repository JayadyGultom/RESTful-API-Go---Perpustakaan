package entity

import (
	"time"
	"gorm.io/gorm"
)

type Anggota struct {
	IDAnggota       int64          `json:"id_anggota" gorm:"primaryKey;autoIncrement;column:id_anggota"`
	NamaLengkap     string         `json:"nama_lengkap"`
	NomorIdentitas  string         `json:"nomor_identitas" gorm:"uniqueIndex"`
	JenisKelamin    string         `json:"jenis_kelamin"`
	Email           string         `json:"email"`
	NomorHP         string         `json:"nomor_hp"`
	Alamat          string         `json:"alamat"`
	Status          string         `json:"status"`
	TanggalBergabung time.Time     `json:"tanggal_bergabung"`
	Foto            string         `json:"foto"`
	KataSandi       string         `json:"kata_sandi"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (Anggota) TableName() string {
	return "anggota"
}