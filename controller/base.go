package controller

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type env struct {
	db *gorm.DB
}

type dbConfig struct {
	DB_USER    string `mapstructure:"DB_USER"`
	DB_PASS    string `mapstructure:"DB_PASS"`
	DB_ADDRESS string `mapstructure:"DB_ADDRESS"`
	DB_SCHEMA  string `mapstructure:"DB_SCHEMA"`
}

var newLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		LogLevel: logger.Info, // Log level
		Colorful: false,       // Disable color
	},
)

func loadConfig(path string) (config dbConfig, err error) {
	var dbInfo dbConfig
	
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&dbInfo)
	return dbInfo, nil
}

func NewEnv() (*env, error) {
	config, err := loadConfig(".")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local`, 
		config.DB_USER, config.DB_PASS, config.DB_ADDRESS, config.DB_SCHEMA), // auto configure based on currently MySQL version
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
