package entity

type KategoriBuku struct {
	IDKategori   int    `gorm:"column:id_kategori;primaryKey" json:"id_kategori"`
	NamaKategori string `gorm:"column:nama_kategori" json:"nama_kategori"`
}

func (KategoriBuku) TableName() string {
	return "kategori_buku"
}
