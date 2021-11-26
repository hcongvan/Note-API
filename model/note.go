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
	errch := PublishMessage(HistoryLog{
		ID:       r.Id,
		Previous: nil,
		Current:  r.ToMap(nil),
		Status:   1,
	})
	if errch != nil {
		return r, errch
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
	result := db.First(&r)
	if result.Error != nil {
		return NoteModel{}, result.Error
	}
	_r := r
	resultUpdate := db.Model(&_r).Select("*").Updates(m)
	if resultUpdate.Error != nil {
		return _r, resultUpdate.Error
	}
	var _sel []string
	for k := range m {
		_sel = append(_sel, k)
	}
	errch := PublishMessage(HistoryLog{
		ID:       r.Id,
		Previous: r.ToMap(_sel),
		Current:  m,
		Status:   2,
	})
	if errch != nil {
		return _r, errch
	}
	return _r, nil
}

func (r NoteModel) DeleteNote() error {
	if result := db.First(&r); result.Error != nil {
		return result.Error
	}
	_r := r
	if resultDelete := db.Delete(&_r); resultDelete.Error != nil {
		return resultDelete.Error
	}
	errch := PublishMessage(HistoryLog{
		ID:       r.Id,
		Previous: r.ToMap(nil),
		Current:  nil,
		Status:   3,
	})
	if errch != nil {
		return errch
	}
	return nil
}

func (r NoteModel) ToMap(sel []string) map[string]interface{} {
	_tmp := map[string]interface{}{
		"user":  r.User,
		"title": r.Title,
		"note":  r.Note,
	}
	if sel != nil {
		var _k []string
		for k := range _tmp {
			for _, v := range sel {
				if k != v {
					_k = append(_k, k)
				}
			}
		}
		for _, idx := range _k {
			delete(_tmp, idx)
		}
	}
	return _tmp
}
