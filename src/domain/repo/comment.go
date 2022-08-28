package repo

import "gd-blog/src/domain/entity"

type CommentRepo interface {
	Select(separateId int, limit int) ([]*entity.Comment, error)
	Save(comment *entity.Comment) (*entity.Comment, error)
	Update(comment *entity.Comment) (*entity.Comment, error)
}
