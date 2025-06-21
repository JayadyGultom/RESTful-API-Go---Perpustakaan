package entity

type DetailPeminjaman struct {
	IDDetail      int64 `json:"id_detail" gorm:"primaryKey;autoIncrement;column:id_detail"`
	IDPeminjaman  int64 `json:"id_peminjaman"`
	IDBuku        int64 `json:"id_buku"`
	Jumlah        int   `json:"jumlah"`
}

func (DetailPeminjaman) TableName() string {
	return "detail_peminjaman"
} 