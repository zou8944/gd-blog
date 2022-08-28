package service

import (
	"gd-blog/src/domain/entity"
	"gd-blog/src/domain/repo"
)

type BlogDomainService struct {
	blogRepo    repo.BlogRepo
	commentRepo repo.CommentRepo
	userRepo    repo.UserRepo
}

func NewBlogDomainService(blogRepo repo.BlogRepo, commentRepo repo.CommentRepo, userRepo repo.UserRepo) BlogDomainService {
	return BlogDomainService{
		blogRepo:    blogRepo,
		commentRepo: commentRepo,
		userRepo:    userRepo,
	}
}

func (bs *BlogDomainService) GetBlog(id int) (*entity.Blog, error) {
	return bs.blogRepo.SelectOne(id)
}

func (bs *BlogDomainService) ListBlog(sepId int, limit int) ([]*entity.Blog, error) {
	if limit <= 0 {
		limit = 20
	}
	return bs.blogRepo.Select(sepId, limit)
}

func (bs *BlogDomainService) SearchBlog(keyword string, limit int) ([]*entity.Blog, error) {
	return bs.blogRepo.Search(keyword, limit)
}

func (bs *BlogDomainService) CreateBlog(title, content string) (*entity.Blog, error) {
	blog, err := entity.NewBlog(title, content)
	if err != nil {
		return nil, err
	}
	return bs.blogRepo.Insert(blog)
}

func (bs *BlogDomainService) UpdateBlog(id int, newTitle string, newContent string) (*entity.Blog, error) {
	blog, err := bs.blogRepo.SelectOne(id)
	if err != nil {
		return nil, err
	}
	blog.Update(newTitle, newContent)
	return bs.blogRepo.Update(blog)
}

func (bs *BlogDomainService) DeleteBlog(id int) (*entity.Blog, error) {
	return bs.blogRepo.Delete(id)
}

// LikeBlog TODO 并发问题
func (bs *BlogDomainService) LikeBlog(id int) (*entity.Blog, error) {
	blog, err := bs.blogRepo.SelectOne(id)
	if err != nil {
		return nil, err
	}
	blog.Like()
	return bs.blogRepo.Update(blog)
}

func (bs *BlogDomainService) CreateComment(username, email, content string, blogId, replyId int) (*entity.Comment, error) {
	user := entity.NewUser(username, email)
	comment := entity.NewComment(user, blogId, replyId, content)
	return bs.commentRepo.Save(comment)
}

func (bs *BlogDomainService) ListComment(sepId, limit int) ([]*entity.Comment, error) {
	return bs.commentRepo.Select(sepId, limit)
}
