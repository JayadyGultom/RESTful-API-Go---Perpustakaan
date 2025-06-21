package entity

import "time"

type Reservasi struct {
	IDReservasi      int64     `json:"id_reservasi" gorm:"primaryKey;autoIncrement;column:id_reservasi"`
	IDAnggota        int64     `json:"id_anggota"`
	IDBuku           int64     `json:"id_buku"`
	TanggalReservasi time.Time `json:"tanggal_reservasi"`
	StatusReservasi  string    `json:"status_reservasi"`
} 