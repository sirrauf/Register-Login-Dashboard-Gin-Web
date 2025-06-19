package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

type User struct {
	gorm.Model
	NamaLengkap string
	Alamat      string
	Email       string `gorm:"unique"`
	Password    string
	Role        string
}

func InitDB() {
	dbkonek := "root:@tcp(127.0.0.1:3306)/db_dashboard?charset=utf8mb4&parseTime=True&loc=Local"
	var errors error
	Database, errors = gorm.Open(mysql.Open(dbkonek), &gorm.Config{})
	if errors != nil {
		panic("Gagal koneksi ke database")
	}
	Database.AutoMigrate(&User{})
}
