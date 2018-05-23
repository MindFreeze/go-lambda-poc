package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "*secret*")
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
	return db
}
