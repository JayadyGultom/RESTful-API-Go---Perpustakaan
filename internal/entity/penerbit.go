package entity

type Penerbit struct {
	IDPenerbit   int    `gorm:"column:id_penerbit;primaryKey" json:"id_penerbit"`
	NamaPenerbit string `gorm:"column:nama_penerbit" json:"nama_penerbit"`
	Alamat       string `gorm:"column:alamat" json:"alamat"`
}

func (Penerbit) TableName() string {
	return "penerbit"
}
