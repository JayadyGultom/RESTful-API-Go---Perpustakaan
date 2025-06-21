package entity

import "time"

type Peminjaman struct {
	IDPeminjaman      int64     `json:"id_peminjaman" gorm:"primaryKey;autoIncrement;column:id_peminjaman"`
	IDAnggota         int64     `json:"id_anggota"`
	IDPetugas         int64     `json:"id_petugas"`
	TanggalPinjam     time.Time `json:"tanggal_pinjam"`
	BatasPengembalian time.Time `json:"batas_pengembalian"`
	StatusPeminjaman  string    `json:"status_peminjaman"`
}

func (Peminjaman) TableName() string {
	return "peminjaman"
} 