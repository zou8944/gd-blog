package repo

import (
	"gd-blog/repo/model"
	"gorm.io/gorm"
)

type CommentRepo interface {
	Create(comment *model.Comment) (*model.Comment, error)
	List(blogId uint) ([]model.Comment, error)
}

type commentRepo struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) CommentRepo {
	return &commentRepo{db: db}
}

func (cr commentRepo) Create(comment *model.Comment) (*model.Comment, error) {
	cr.db.Create(comment)
	return comment, cr.db.Error
}

func (cr commentRepo) List(blogId uint) ([]model.Comment, error) {
	var comments []model.Comment
	cr.db.Model(&model.Comment{}).Where("blog_id = ?", blogId).Order("created_at desc").Find(&comments)
	return comments, cr.db.Error
}
