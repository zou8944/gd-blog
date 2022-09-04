package main

import (
	"gd-blog/src/repo/model"
	"gopkg.in/errgo.v2/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {
	db, err := gorm.Open(sqlite.Open("sqlite/blog.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	err = db.AutoMigrate(&model.Category{}, &model.Tag{}, &model.Visitor{})
	err = db.AutoMigrate(&model.Blog{}, &model.Comment{})
	if err != nil {
		log.Fatalln(errors.Wrap(err))
	}
}
