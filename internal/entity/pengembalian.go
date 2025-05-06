package entity

import "time"

type Pengembalian struct {
    IDPengembalian    int       `gorm:"column:id_pengembalian;primaryKey"`
    IDPeminjaman      int       `gorm:"column:id_peminjaman"`
    TanggalKembali    time.Time `gorm:"column:tanggal_kembali"`
    KondisiBuku        string   `gorm:"column:kondisi_buku"`
    Peminjaman  Peminjaman `gorm:"foreignKey:id_peminjaman;references:id_peminjaman" json:"-"`
}

func (Pengembalian) TableName() string {
    return "pengembalian"
} 