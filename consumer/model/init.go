package model

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB(db_name string) {
	_db, err := gorm.Open(sqlite.Open(db_name), &gorm.Config{})
	if err != nil {
		fmt.Printf("failed to connect database")
	}
	_db.AutoMigrate(&NoteHistory{})
	db = _db
}
