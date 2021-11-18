package model

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(db_name string) {
	db, err := gorm.Open(sqlite.Open(db_name), &gorm.Config{})
	if err != nil {
		fmt.Printf("failed to connect database")
	}
	db.AutoMigrate(&NoteModel{})
	DB = db
}
