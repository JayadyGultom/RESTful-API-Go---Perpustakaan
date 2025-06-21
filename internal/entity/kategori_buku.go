package entity

import "time"

type KategoriBuku struct {
	IDKategori   int64     `json:"id_kategori" gorm:"primaryKey;autoIncrement;column:id_kategori"`
	NamaKategori string    `json:"nama_kategori" gorm:"column:nama_kategori"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (KategoriBuku) TableName() string {
	return "kategori_buku"
} 