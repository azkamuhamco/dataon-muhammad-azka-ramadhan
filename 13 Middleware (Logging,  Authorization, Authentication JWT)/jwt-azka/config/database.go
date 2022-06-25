package config

import (
	"fmt"
	"jwt-azka/model"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	InitDB()
}

func InitDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(ConnString()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func ConnString() string {
	config := model.Config{
		DB_Username: "root",
		DB_Password: "1234567",
		DB_Port:     "3306",
		DB_Host:     "localhost",
		DB_Name:     "crud_go",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	return connectionString
}
