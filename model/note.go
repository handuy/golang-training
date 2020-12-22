package model

import (
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
