package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "src/repo",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	db, err := gorm.Open(sqlite.Open("sqlite/blog.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	g.UseDB(db)
	g.GenerateModel("blog")
	g.GenerateModel("category")
	g.GenerateModel("label")
	g.GenerateModel("comment")
	g.GenerateModel("visitor")
	g.Execute()
}
