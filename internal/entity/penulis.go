package entity

import "time"

type Penulis struct {
	IDPenulis   int64     `json:"id_penulis" gorm:"primaryKey;autoIncrement;column:id_penulis"`
	NamaPenulis string    `json:"nama_penulis" gorm:"column:nama_penulis"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (Penulis) TableName() string {
	return "penulis"
} 