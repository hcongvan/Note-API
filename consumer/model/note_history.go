package model

import (
	"time"
)

type NoteHistory struct {
	Id         uint      `json:"id" gorm:"primary_key AUTO_INCREMENT"`
	PreviousId uint      `json:"previous_id"`
	NoteId     uint      `json:"note_id"`
	Status     bool      `json:"status"`
	Type       uint      `json:"type"`
	User       string    `json:"user"`
	Title      string    `json:"title"`
	Note       string    `json:"note"`
	CreateAt   time.Time `json:"create_at"`
}

func (r NoteHistory) CreateNoteHistory() (NoteHistory, error) {
	result := db.Create(&r)
	if result.Error != nil {
		return NoteHistory{}, result.Error
	}
	return r, nil
}

func (r NoteHistory) GetNoteHistory() (NoteHistory, error) {
	result := db.Find(&r)
	if result.Error != nil {
		return NoteHistory{}, result.Error
	}
	return r, nil
}

func (r NoteHistory) ReadAllNoteHistory(offset int, limit int, order bool) []NoteHistory {
	var _r []NoteHistory
	result := db.Where("status = ?", true).Find(&_r)
	if result.Error != nil {
		return []NoteHistory{}
	}
	return _r
}

func (r NoteHistory) ReadNoteHistory(noteid int) []NoteHistory {
	var _r []NoteHistory
	result := db.Where("note_id = ? AND status = ?", noteid, true).Find(&_r)
	if result.Error != nil {
		return []NoteHistory{}
	}
	return _r
}

func (r NoteHistory) UpdateNoteHistory(m map[string]interface{}) (NoteHistory, error) {
	result := db.Find(&r)
	if result.Error != nil {
		return NoteHistory{}, result.Error
	}
	_r := r
	resultUpdate := db.Model(&_r).Select("*").Updates(m)
	if resultUpdate.Error != nil {
		return _r, resultUpdate.Error
	}
	return _r, nil
}

func (r NoteHistory) DeleteNoteHistory() error {
	result := db.Model(&NoteHistory{}).Where("note_id = ?", r.NoteId).Update("status", false)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
