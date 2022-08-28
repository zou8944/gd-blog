package repo

import "gd-blog/src/domain/entity"

type BlogRepo interface {
	SelectOne(id int) (*entity.Blog, error)
	Select(separateId int, limit int) ([]*entity.Blog, error)
	Search(keyword string, limit int) ([]*entity.Blog, error)
	Insert(blog *entity.Blog) (*entity.Blog, error)
	Update(blog *entity.Blog) (*entity.Blog, error)
	Delete(id int) (*entity.Blog, error)
}
