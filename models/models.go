package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title	string
	Author     string
}

func DBMigrate(db *gorm.DB) {
	db.AutoMigrate(&Book{})
}
