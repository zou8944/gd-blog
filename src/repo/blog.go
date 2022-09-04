package repo

import (
	"gd-blog/src/repo/model"
	"gorm.io/gorm"
)

type BlogRepo struct {
	db *gorm.DB
}

func NewBlogRepo(db *gorm.DB) *BlogRepo {
	return &BlogRepo{db: db}
}

func (b *BlogRepo) SelectOne(id int) (model.Blog, error) {
	var blogModel model.Blog
	b.db.Model(&model.Blog{}).Preload("Categories").First(&blogModel)
	return blogModel, b.db.Error
}

func (b *BlogRepo) Select(separateId int, limit int) ([]model.Blog, error) {
	blogModels := []model.Blog{}
	dsl := b.db.Preload("Categories").Preload("Tags")
	if separateId != 0 {
		dsl = dsl.Where("created_at < (select created_at from blog where id = ?)", separateId)
	}
	dsl.Limit(limit).Order("created_at desc").Find(&blogModels)
	return blogModels, b.db.Error
}

func (b *BlogRepo) Search(keyword string) ([]model.Blog, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BlogRepo) Insert(blog model.Blog) (model.Blog, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BlogRepo) Update(blog model.Blog) (model.Blog, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BlogRepo) Delete(id int) (model.Blog, error) {
	//TODO implement me
	panic("implement me")
}
