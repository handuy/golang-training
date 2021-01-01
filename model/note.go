package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Note struct {
	Id        int
	Title     string
	Status    bool
	Author    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type StatusMessage struct {
	Message string
}

type NewNote struct {
	Title string
}

type UpdatedNote struct {
	Id        int
	Title     string
	Status    bool
	UpdatedAt time.Time
}

type DeletedNote struct {
	Id int
}

func GetAllNotes(db *gorm.DB) ([]Note, error) {
	rows, err := db.Raw("select * from notes").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Note
	var item Note
	for rows.Next() {
		db.ScanRows(rows, &item)
		result = append(result, item)
	}

	return result, nil
}

func GetNoteById(db *gorm.DB, noteId int) (Note, error) {
	var result Note
	err := db.Raw("select * from notes where id = ?", noteId).Scan(&result).Error
	if err != nil {
		return result, err
	}
	if result == (Note{}) {
		return result, errors.New("Không tìm thấy note")
	}

	return result, nil
}

func InsertNote(db *gorm.DB, newNote NewNote, author string) (Note, error) {
	var result Note
	result.Title = newNote.Title
	result.Status = false
	result.Author = author
	result.CreatedAt = time.Now()

	err := db.Create(&result).Error
	if err != nil {
		return Note{}, err
	}

	return result, nil
}

func UpdateNote(db *gorm.DB, updateNote UpdatedNote) error {
	result := db.Omit("created_at").Updates(&Note{
		Id:        updateNote.Id,
		Title:     updateNote.Title,
		Status:    updateNote.Status,
		UpdatedAt: updateNote.UpdatedAt,
	}).RowsAffected
	if result == 0 {
		return errors.New("Không tìm thấy note")
	}

	return nil
}

func DeleteNote(db *gorm.DB, deletedNote DeletedNote) error {
	result := db.Delete(&Note{
		Id:        deletedNote.Id,
	}).RowsAffected
	if result == 0 {
		return errors.New("Không tìm thấy note")
	}

	return nil
}
