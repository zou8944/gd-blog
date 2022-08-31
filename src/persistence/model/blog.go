package model

import "gorm.io/gorm"

const TableNameBlog = "blog"

type Blog struct {
	gorm.Model
	Title        string
	Summary      string
	Content      string
	LikeCount    int32
	CollectCount int32
	Categories   []Category `gorm:"many2many:blog_categories"`
	Tags         []Tag      `gorm:"many2many:blog_tags"`
}

func (*Blog) TableName() string {
	return TableNameBlog
}
