package entity

type Buku struct {
	IDBuku      int    `gorm:"primaryKey;column:id_buku" json:"id_buku"`
	Judul       string `gorm:"type:varchar(255)" json:"judul"`
	IDPenulis   int    `gorm:"column:id_penulis" json:"id_penulis"`
	IDPenerbit  int    `gorm:"column:id_penerbit" json:"id_penerbit"`
	IDKategori  int    `gorm:"column:id_kategori" json:"id_kategori"`
	TahunTerbit int    `gorm:"type:year(4)" json:"tahun_terbit"`
	ISBN        string `gorm:"type:varchar(20)" json:"isbn"`
	Stok        int    `gorm:"type:int" json:"stok"`
}

func (Buku) TableName() string {
	return "buku"
}
