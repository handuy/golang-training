package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type note struct {
	Id        int
	Title     string
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ErrorMessage struct {
	Message string
}

func GetAllNotes(db *gorm.DB) ([]note, error) {
	rows, err := db.Raw("select * from notes").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []note
	var item note
	for rows.Next() {
		db.ScanRows(rows, &item)
		result = append(result, item)
	}

	return result, nil
}

func GetNoteById(db *gorm.DB, noteId int) (note, error) {
	var result note
	err := db.Raw("select * from notes where id = ?", noteId).Scan(&result).Error
	if err != nil {
		return result, err
	}
	if result == (note{}) {
		return result, errors.New("Không tìm thấy note")
	}

	return result, nil
}
