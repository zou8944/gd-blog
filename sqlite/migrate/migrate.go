package main

import (
	model2 "gd-blog/repo/model"
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
	err = db.AutoMigrate(&model2.Category{}, &model2.Tag{}, &model2.Visitor{})
	err = db.AutoMigrate(&model2.Blog{}, &model2.Comment{})
	if err != nil {
		log.Fatalln(errors.Wrap(err))
	}
}
