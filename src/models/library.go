package models

import (
	_ "github.com/jinzhu/gorm"
)

type Library struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}
