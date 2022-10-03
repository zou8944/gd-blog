package repo

import (
	"gd-blog/facade/dto"
	"gd-blog/repo/model"
	"gorm.io/gorm"
)

type BlogRepo interface {
	SelectStat() (*dto.Statistics, error)
	SelectOne(id int) (model.Blog, error)
	Select(pageNo int, pageSize int, cid int) ([]model.Blog, error)
	Count() (int, error)
	Search(keyword string) ([]model.Blog, error)
	SelectTags() ([]dto.Tag, error)
	SelectCategories() ([]dto.Category, error)
}

type blogRepo struct {
	db *gorm.DB
}

func NewBlogRepo(db *gorm.DB) BlogRepo {
	return &blogRepo{db: db}
}

func (br *blogRepo) SelectStat() (*dto.Statistics, error) {
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

func (br *blogRepo) SelectOne(id int) (model.Blog, error) {
	var blogModel model.Blog
	br.db.Model(&model.Blog{}).Preload("Categories").Where("id = ?", id).First(&blogModel)
	return blogModel, br.db.Error
}

func (br *blogRepo) Select(pageNo int, pageSize int, cid int) ([]model.Blog, error) {
	blogModels := []model.Blog{}
	db := br.db.Preload("Categories").Preload("Tags")
	if cid > 0 {
		db.Raw("select * from blog b inner join blog_categories bc on b.id = bc.blog_id where bc.category_id = ? order by b.created_at desc limit ? offset ?", cid, pageSize, pageSize*(pageNo-1)).Find(&blogModels)
	} else {
		db.Limit(pageSize).Offset(pageSize * (pageNo - 1)).Order("created_at desc").Find(&blogModels)
	}
	return blogModels, db.Error
}

func (br *blogRepo) SelectTags() ([]dto.Tag, error) {
	var results []map[string]interface{}
	br.db.Model(&model.Tag{}).Select("id, name, (select count(1) from blog_tags where tag_id = tag.id) as blog_count").Order("blog_count desc").Find(&results)
	if br.db.Error != nil {
		return nil, br.db.Error
	}
	var tags []dto.Tag
	for _, r := range results {
		tag := dto.Tag{
			ID:        r["id"].(uint),
			Name:      r["name"].(string),
			BlogCount: r["blog_count"].(int64),
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

func (br *blogRepo) SelectCategories() ([]dto.Category, error) {
	var results []map[string]interface{}
	br.db.Model(&model.Category{}).Select("id, name, (select count(1) from blog_categories where category_id = category.id) as blog_count").Order("blog_count desc").Find(&results)
	if br.db.Error != nil {
		return nil, br.db.Error
	}
	var categories []dto.Category
	for _, r := range results {
		category := dto.Category{
			ID:        r["id"].(uint),
			Name:      r["name"].(string),
			BlogCount: r["blog_count"].(int64),
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (br *blogRepo) Count() (int, error) {
	var count int64
	br.db.Model(&model.Blog{}).Count(&count)
	return int(count), br.db.Error
}

func (br *blogRepo) Search(keyword string) ([]model.Blog, error) {
	panic("not implemented")
}
