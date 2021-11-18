package model

import (
	"fmt"

	"gorm.io/gorm"
)

type NoteModel struct {
	gorm.Model
	Id    uint
	User  string
	Title string
	Note  string
}

func (r NoteModel) CreateNote(db *gorm.DB) {
	result := db.Create(&r)
	if result != nil {
		fmt.Printf("Can't not create")
	}
}

func (r NoteModel) GetAllNote(db *gorm.DB) {
	result := db.Find(&NoteModel{})
	if result == nil {
		fmt.Printf("tach luon roi")
	}

}

func (r NoteModel) QueryNotebyUser(user string, db *gorm.DB) {

}

func (r NoteModel) QueryNotebyId(id uint) {

}

func (r NoteModel) UpdateNotebyId(db *gorm.DB) {
	result := db.First(&r)
	if result == nil {
		panic("tach cmnr")
	}
	db.Save(&r)
}

func (r NoteModel) DeleteNotebyId(db *gorm.DB) {
	result := db.First(&r)
	if result == nil {
		panic("tach cmnr")
	}
	db.Delete(&r)
}
