package entity

type Penulis struct {
    IDPenulis   int    `gorm:"column:id_penulis;primaryKey"`
    NamaPenulis string `gorm:"column:nama_penulis"`
}

func (Penulis) TableName() string {
    return "penulis"
} 