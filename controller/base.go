package controller

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type env struct {
	db *gorm.DB
}

var newLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		LogLevel: logger.Info, // Log level
		Colorful: false,       // Disable color
	},
)

func NewEnv() (*env, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "duy:123456@tcp(localhost:3306)/todolist?charset=utf8&parseTime=True&loc=Local", // auto configure based on currently MySQL version
	}), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}
	return &env{
		db: db,
	}, nil
}
