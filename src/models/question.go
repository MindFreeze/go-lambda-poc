package models

import (
	// "database/sql"
	_ "github.com/jinzhu/gorm"
	"gopkg.in/guregu/null.v3"
	// "github.com/satori/go.uuid"
)

type Question struct {
	ID             uint      `db:"id" gorm:"primary_key" json:"id"`
	Libraryid      uint      `db:"libraryid" json:"libraryid"`
	Questiontypeid uint      `db:"questiontypeid" json:"questiontypeid"`
	Text           string    `db:"text" json:"text"`
	Required       null.Bool `db:"required" json:"required"`
}
