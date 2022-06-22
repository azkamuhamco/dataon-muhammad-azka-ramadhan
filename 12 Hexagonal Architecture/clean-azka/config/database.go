package config

import (
	"clean-azka/model"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

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
