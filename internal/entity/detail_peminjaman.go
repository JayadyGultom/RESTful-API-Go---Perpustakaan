package entity

type DetailPeminjaman struct {
	IDDetail     int `gorm:"column:id_detail;primaryKey" json:"id_detail"`
	IDPeminjaman int `gorm:"column:id_peminjaman" json:"id_peminjaman"`
	IDBuku       int `gorm:"column:id_buku" json:"id_buku"`
	JumlahBuku   int `gorm:"column:jumlah" json:"jumlah"`
	Buku         Buku `gorm:"foreignKey:IDBuku;references:IDBuku" json:"-"`
}

func (DetailPeminjaman) TableName() string {
	return "detail_peminjaman"
}
