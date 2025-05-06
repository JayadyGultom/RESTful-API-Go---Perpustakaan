package entity

type Denda struct {
    IDDenda         int     `gorm:"column:id_denda;primaryKey"`
    IDPengembalian  int     `gorm:"column:id_pengembalian"`
    JumlahDenda     float64 `gorm:"column:jumlah_denda"`
    StatusPembayaran string `gorm:"column:status_pembayaran"`
    Pengembalian          Pengembalian   `gorm:"foreignKey:IDPengembalian;references:IDPengembalian" json:"-"`
}

func (Denda) TableName() string {
    return "denda"
} 