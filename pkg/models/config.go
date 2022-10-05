package models

import (
	"gd-blog/configs"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var db *gorm.DB

func Init() error {
	dbLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		LogLevel: logger.Info,
	})
	newDB, err := gorm.Open(sqlite.Open(configs.Database.FilePath), &gorm.Config{Logger: dbLogger})
	db = newDB
	return err
}
