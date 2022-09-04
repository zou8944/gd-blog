package repo

import (
	"gorm.io/gorm"
)

type CommentRepo interface {
}

type commentRepo struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) CommentRepo {
	return &commentRepo{db: db}
}
