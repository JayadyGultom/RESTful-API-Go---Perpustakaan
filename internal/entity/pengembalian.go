package entity

import "time"

type Pengembalian struct {
	IDPengembalian   int64     `json:"id_pengembalian" gorm:"primaryKey;autoIncrement;column:id_pengembalian"`
	IDPeminjaman     int64     `json:"id_peminjaman"`
	TanggalKembali   time.Time `json:"tanggal_kembali"`
	KondisiBuku      string    `json:"kondisi_buku"`
} 