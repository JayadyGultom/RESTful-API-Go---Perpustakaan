package config

import (
    "fmt"
    "log"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := "root:@tcp(127.0.0.1:3306)/perpustakaan?charset=utf8mb4&parseTime=True&loc=Local"

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
        DisableForeignKeyConstraintWhenMigrating: true,
    })
    if err != nil {
        log.Fatal("Gagal koneksi ke database: ", err)
    }

    fmt.Println("Database berhasil terkoneksi.")
    DB = db
}
