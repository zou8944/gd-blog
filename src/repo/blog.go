package repo

import (
	"gd-blog/src/facade/dto"
	"gd-blog/src/repo/model"
	"gorm.io/gorm"
)

type BlogRepo struct {
	db *gorm.DB
}

func NewBlogRepo(db *gorm.DB) *BlogRepo {
	return &BlogRepo{db: db}
}

func (br *BlogRepo) SelectStat() (*dto.Statistics, error) {
	var blogCount int64
	var categoryCount int64
	var tagCount int64
	var visitorCount int64
	var viewCount int64
	br.db.Model(&model.Blog{}).Count(&blogCount)
	br.db.Model(&model.Category{}).Count(&categoryCount)
	br.db.Model(&model.Tag{}).Count(&tagCount)
	var viewCountResult map[string]int64
	br.db.Model(&model.Blog{}).Select("sum(view_count) as total").First(&viewCountResult)
	if br.db.Error != nil {
		return nil, br.db.Error
	}
	viewCount = viewCountResult["total"]
	visitorCount = viewCount / 3
	return &dto.Statistics{
		BlogCount:     blogCount,
		CategoryCount: categoryCount,
		TagCount:      tagCount,
		VisitorCount:  visitorCount,
		ViewCount:     viewCount,
	}, nil
}

func (br *BlogRepo) SelectOne(id int) (model.Blog, error) {
	var blogModel model.Blog
	br.db.Model(&model.Blog{}).Preload("Categories").First(&blogModel)
	return blogModel, br.db.Error
}

func (br *BlogRepo) Select(separateId int, limit int) ([]model.Blog, error) {
	blogModels := []model.Blog{}
	dsl := br.db.Preload("Categories").Preload("Tags")
	if separateId != 0 {
		dsl = dsl.Where("created_at < (select created_at from blog where id = ?)", separateId)
	}
	dsl.Limit(limit).Order("created_at desc").Find(&blogModels)
	return blogModels, br.db.Error
}

func (br *BlogRepo) Search(keyword string) ([]model.Blog, error) {
	//TODO implement me
	panic("implement me")
}

func (br *BlogRepo) Insert(blog model.Blog) (model.Blog, error) {
	//TODO implement me
	panic("implement me")
}

func (br *BlogRepo) Update(blog model.Blog) (model.Blog, error) {
	//TODO implement me
	panic("implement me")
}

func (br *BlogRepo) Delete(id int) (model.Blog, error) {
	//TODO implement me
	panic("implement me")
}
