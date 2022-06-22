package config

import (
	"clean-azka/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var debe = model.DB
var guna = model.User{}
var connStr = model.ConnString()

func InitDB() {
	var err error
	debe, err = gorm.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	debe.AutoMigrate(&guna)
}
