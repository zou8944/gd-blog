package repo

import (
	"gd-blog/src/facade/dto"
	"gorm.io/gorm"
)

type CommentRepo struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) *CommentRepo {
	return &CommentRepo{db: db}
}

func (cr CommentRepo) Select(separateId int, limit int) ([]*dto.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (cr CommentRepo) Save(comment *dto.Comment) (*dto.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (cr CommentRepo) Update(comment *dto.Comment) (*dto.Comment, error) {
	//TODO implement me
	panic("implement me")
}
