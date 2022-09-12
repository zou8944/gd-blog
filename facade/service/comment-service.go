package service

import (
	"gd-blog/facade/dto"
	"gd-blog/repo"
)

type CommentService interface {
	Create(comment *dto.Comment) (*dto.Comment, error)
	List(blogId uint) ([]dto.Comment, error)
}

type commentService struct {
	commentRepo repo.CommentRepo
}

func NewCommentService(commentRepo repo.CommentRepo) CommentService {
	return commentService{commentRepo: commentRepo}
}

func (cs commentService) Create(comment *dto.Comment) (*dto.Comment, error) {
	cm, err := comment.ToModel()
	if err != nil {
		return nil, err
	}
	cm, err = cs.commentRepo.Create(cm)
	if err != nil {
		return nil, err
	}
	return dto.ConvertCM2CT(cm)
}

func (cs commentService) List(blogId uint) ([]dto.Comment, error) {
	cms, err := cs.commentRepo.List(blogId)
	if err != nil {
		return nil, err
	}
	return dto.ConvertCMS2CTS(cms)
}
