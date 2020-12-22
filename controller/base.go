package controller

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type env struct {
	db *gorm.DB
}

func NewEnv() (*env, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "duy:123456@tcp(localhost:3306)/todolist?charset=utf8&parseTime=True&loc=Local",                                                                   // auto configure based on currently MySQL version
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &env{
		db: db,
	}, nil
}
