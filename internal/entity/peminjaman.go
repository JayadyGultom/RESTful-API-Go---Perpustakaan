package entity

import "time"

type Peminjaman struct {
	IDPeminjaman      int       `gorm:"column:id_peminjaman;primaryKey" json:"id_peminjaman"`
	IDAnggota         int       `gorm:"column:id_anggota" json:"id_anggota"`
	IDPetugas         int       `gorm:"column:id_petugas" json:"id_petugas"`
	TanggalPinjam     time.Time `gorm:"column:tanggal_pinjam" json:"tanggal_pinjam"`
	BatasPengembalian time.Time `gorm:"column:batas_pengembalian" json:"batas_pengembalian"`
	DetailPeminjaman  []DetailPeminjaman `gorm:"foreignKey:IDPeminjaman;references:IDPeminjaman" json:"-"`
	Anggota           Anggota   `gorm:"foreignKey:IDAnggota;references:IDAnggota" json:"-"`
	Petugas           Petugas   `gorm:"foreignKey:IDPetugas;references:IDPetugas" json:"-"`
}	

func (Peminjaman) TableName() string {
	return "peminjaman"
}

type PeminjamanResponse struct {
	IDPeminjaman      int       `json:"id_peminjaman"`
	IDAnggota         int       `json:"id_anggota"`
	IDPetugas         int       `json:"id_petugas"`
	TanggalPinjam     time.Time `json:"tanggal_pinjam"`
	BatasPengembalian time.Time `json:"batas_pengembalian"`
}
