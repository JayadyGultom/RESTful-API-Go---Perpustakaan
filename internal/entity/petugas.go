package entity

import "time"

type Petugas struct {
	IDPetugas int64     `json:"id_petugas" gorm:"primaryKey;autoIncrement;column:id_petugas"`
	Nama      string    `json:"nama" gorm:"column:nama"`
	Jabatan   string    `json:"jabatan" gorm:"column:jabatan"`
	Username  string    `json:"username" gorm:"column:username"`
	Password  string    `json:"password" gorm:"column:password"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (Petugas) TableName() string {
	return "petugas"
} 