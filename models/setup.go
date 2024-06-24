package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {

	dataBase, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go_restapi_gin"))
	if err != nil {
		panic(err)

	}

	dataBase.AutoMigrate(&Product{})
	DB = dataBase

}
