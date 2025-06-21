package entity

import (
	"time"
)

type Buku struct {
	IDBuku        int64     `json:"id_buku" gorm:"primaryKey;autoIncrement;column:id_buku"`
	Judul         string    `json:"judul" gorm:"column:judul"`
	IDPenulis     int64     `json:"id_penulis" gorm:"column:id_penulis"`
	IDPenerbit    int64     `json:"id_penerbit" gorm:"column:id_penerbit"`
	IDKategori    int64     `json:"id_kategori" gorm:"column:id_kategori"`
	TahunTerbit   int       `json:"tahun_terbit" gorm:"column:tahun_terbit"`
	ISBN          string    `json:"isbn" gorm:"column:isbn"`
	Deskripsi     string    `json:"deskripsi" gorm:"column:deskripsi"`
	LokasiRak     string    `json:"lokasi_rak" gorm:"column:lokasi_rak"`
	Stok          int64     `json:"stok" gorm:"column:stok"`
	CoverBuku     string    `json:"cover_buku" gorm:"column:cover_buku"`
	CreatedAt     time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (Buku) TableName() string {
	return "buku"
} 