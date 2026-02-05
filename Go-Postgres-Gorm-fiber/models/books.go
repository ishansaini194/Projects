package models

import "gorm.io/gorm"

type Book struct {
	ID        uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	Author    *string `json:"author"`
	Title     *string `json:"title"`
	Publisher *string `json:"publisher"`
}

func MigrateBooks(db *gorm.DB) error {
	return db.AutoMigrate(&Book{})
}
