package config

import (
	"first-Go/API/structs"
	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB {
	//Set database, harus dibuat dahulu nama databasenya
	//Format args yang digunakan: username:password@tcp(127.0.0.1:3306)/namadatabase?charset=utf8&parseTime=True
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/godb?charset=utf8&parseTime=True")
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(structs.Person{})
	return db
}
