package model

import (
	"time"

	"gorm.io/gorm"
)

type NoteModel struct {
	Id       uint      `gorm:"primary_key, AUTO_INCREMENT" json:"id"`
	User     string    `json:"user"`
	Title    string    `json:"title" binding:"required"`
	Note     string    `json:"note" binding:"required"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

func (r *NoteModel) BeforeCreate(tx *gorm.DB) (err error) {
	r.CreateAt = time.Now().UTC()
	r.UpdateAt = time.Now().UTC()
	return
}

func (r *NoteModel) BeforeUpdate(tx *gorm.DB) (err error) {
	r.UpdateAt = time.Now().UTC()
	return
}

func (r NoteModel) CreateNote() (NoteModel, error) {
	result := db.Create(&r)
	if result.Error != nil {
		return NoteModel{}, result.Error
	}
	return r, nil
}

func (r NoteModel) GetAllNote() ([]NoteModel, error) {
	var notes []NoteModel
	result := db.Find(&notes)
	if result.Error != nil {
		return []NoteModel{}, result.Error
	}
	return notes, nil
}

func (r NoteModel) QueryNote() (NoteModel, error) {
	result := db.First(&r)
	if result.Error != nil {
		return NoteModel{}, result.Error
	}
	return r, nil
}

func (r NoteModel) UpdateNote(m map[string]interface{}) (NoteModel, error) {
	_r := r
	result := db.First(&_r)
	if result.Error != nil {
		return NoteModel{}, result.Error
	}

	resultUpdate := db.Model(&_r).Select("*").Updates(m)
	if resultUpdate.Error != nil {
		return _r, resultUpdate.Error
	}
	return _r, nil
}

func (r NoteModel) DeleteNote() error {
	if result := db.First(&r); result.Error != nil {
		return result.Error
	}
	if resultDelete := db.Delete(&r); resultDelete.Error != nil {
		return resultDelete.Error
	}
	return nil
}
