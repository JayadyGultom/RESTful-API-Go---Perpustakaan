package entity

import (
	"time"
	"gorm.io/gorm"
)

type Denda struct {
	IDDenda         int64          `json:"id_denda" gorm:"primaryKey;autoIncrement;column:id_denda"`
	IDPengembalian  int64          `json:"id_pengembalian"`
	JumlahDenda     float64        `json:"jumlah_denda"`
	StatusPembayaran string        `json:"status_pembayaran"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (Denda) TableName() string {
	return "denda"
} 