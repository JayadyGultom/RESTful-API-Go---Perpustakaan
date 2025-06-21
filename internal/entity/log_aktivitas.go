package entity

import "time"

type LogAktivitas struct {
	IDLog     int64     `json:"id_log" gorm:"primaryKey;autoIncrement;column:id_log"`
	IDPengguna int64    `json:"id_pengguna"`
	Peran     string    `json:"peran"`
	Aksi      string    `json:"aksi"`
	Waktu     time.Time `json:"waktu"`
}

func (LogAktivitas) TableName() string {
	return "log_aktivitas"
} 