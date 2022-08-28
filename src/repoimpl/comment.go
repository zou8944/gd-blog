package repoimpl

import (
	"gd-blog/src/domain/entity"
	"gorm.io/gorm"
)

type CommentRepoImpl struct {
	db *gorm.DB
}

func NewCommentRepoImpl(db *gorm.DB) CommentRepoImpl {
	return CommentRepoImpl{db: db}
}

func (cr CommentRepoImpl) Select(separateId int, limit int) ([]*entity.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (cr CommentRepoImpl) Save(comment *entity.Comment) (*entity.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (cr CommentRepoImpl) Update(comment *entity.Comment) (*entity.Comment, error) {
	//TODO implement me
	panic("implement me")
}
