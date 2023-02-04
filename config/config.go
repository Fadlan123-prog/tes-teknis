package config

import (
	"project/tes-teknis/structs"

	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:NmEximus_123@tcp(127.0.0.1:3306)/teknis_db?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic("failed ot connect to Database")
	}

	db.AutoMigrate(structs.Product{})
	return db
}
