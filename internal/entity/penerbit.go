package entity

import "time"

type Penerbit struct {
	IDPenerbit   int64     `json:"id_penerbit" gorm:"primaryKey;autoIncrement;column:id_penerbit"`
	NamaPenerbit string    `json:"nama_penerbit" gorm:"column:nama_penerbit"`
	Alamat       string    `json:"alamat" gorm:"column:alamat"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (Penerbit) TableName() string {
	return "penerbit"
} 