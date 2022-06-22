package config

import (
	"clean-azka/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var guna = model.User{}
var connStr = model.ConnString()

func InitDB() {
	var err error
	DB, err = gorm.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&guna)
}
