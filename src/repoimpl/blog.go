package repoimpl

import (
	"gd-blog/src/domain/entity"
	"gd-blog/src/repoimpl/model"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type BlogRepoImpl struct {
	db *gorm.DB
}

func NewBlogRepoImpl(db *gorm.DB) BlogRepoImpl {
	return BlogRepoImpl{db: db}
}

func (b BlogRepoImpl) SelectOne(id int) (*entity.Blog, error) {
	var blogModel model.Blog
	b.db.Model(&model.Blog{}).Preload("Categories").First(&blogModel)
	var blogEntity entity.Blog
	err := copier.Copy(blogEntity, blogModel)
	return &blogEntity, err
}

func (b BlogRepoImpl) Select(separateId int, limit int) ([]*entity.Blog, error) {
	var blogModels []model.Blog
	b.db.Model(&model.Blog{}).Find(&blogModels)
	blogEntities := []*entity.Blog{}
	for _, blogModel := range blogModels {
		var blogEntity entity.Blog
		err := copier.Copy(blogEntity, blogModel)
		if err != nil {
			return nil, err
		}
		blogEntities = append(blogEntities, &blogEntity)
	}
	return blogEntities, nil
}

func (b BlogRepoImpl) Search(keyword string, limit int) ([]*entity.Blog, error) {
	//TODO implement me
	panic("implement me")
}

func (b BlogRepoImpl) Insert(blog *entity.Blog) (*entity.Blog, error) {
	//TODO implement me
	panic("implement me")
}

func (b BlogRepoImpl) Update(blog *entity.Blog) (*entity.Blog, error) {
	//TODO implement me
	panic("implement me")
}

func (b BlogRepoImpl) Delete(id int) (*entity.Blog, error) {
	//TODO implement me
	panic("implement me")
}
