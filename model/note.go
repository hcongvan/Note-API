package model

import (
	"log"

	"gorm.io/gorm"
)

type NoteModel struct {
	gorm.Model
	User  string `json:"user"`
	Title string `json:"title" binding:"required"`
	Note  string `json:"note" binding:"required"`
}

func (r NoteModel) CreateNote() (NoteModel, bool) {
	result := db.Create(&r)
	if result == nil {
		log.Println("Can't create record")
		return NoteModel{}, true
	}
	return r, false
}

func (r NoteModel) GetAllNote() []NoteModel {
	var notes []NoteModel
	result := db.Find(&notes)
	if result == nil {
		log.Println("Something wrong went get all record")
	}
	return notes
}

func (r NoteModel) QueryNote() NoteModel {
	result := db.First(&r)
	if result == nil {
		log.Println("Can't find the record")
	}
	return r
}

func (r NoteModel) UpdateNotebyId() (NoteModel, bool) {
	_r := r
	result := db.First(&_r)
	if result == nil {
		log.Println("Can't find the record to update")
		return NoteModel{}, true
	}
	if r.User != "" {
		_r.User = r.User
	}
	if r.Title != "" {
		_r.Title = r.Title
	}
	if r.Note != "" {
		_r.Note = r.Note
	}
	db.Save(&_r)
	return _r, false
}

func (r NoteModel) DeleteNotebyId() bool {
	result := db.First(&r)
	if result == nil {
		log.Println("Can't find record to delete")
		return true
	}
	db.Delete(&r)
	return false
}
