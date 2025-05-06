package entity

type Petugas struct {
    IDPetugas   int    `gorm:"column:id_petugas;primaryKey"`
    Nama string `gorm:"column:nama"`
    Jabatan      string `gorm:"column:jabatan"`
    Username    string `gorm:"column:username"`
    Password    string `gorm:"column:password"`
    
}

func (Petugas) TableName() string {
    return "petugas"
} 