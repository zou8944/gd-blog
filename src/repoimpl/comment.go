package repoimpl

import (
	"database/sql"
	"gd-blog/src/domain/entity"
)

type CommentRepoImpl struct {
	db *sql.DB
}

func NewCommentRepoImpl(db *sql.DB) CommentRepoImpl {
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
