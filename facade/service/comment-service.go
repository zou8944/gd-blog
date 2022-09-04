package service

import "gd-blog/repo"

type CommentService interface {
}

type commentService struct {
	commentRepo repo.CommentRepo
}

func NewCommentService(commentRepo repo.CommentRepo) CommentService {
	return commentService{commentRepo: commentRepo}
}
