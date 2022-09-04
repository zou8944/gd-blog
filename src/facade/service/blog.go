package service

import (
	"gd-blog/src/facade/dto"
	"gd-blog/src/repo"
	"gopkg.in/errgo.v2/errors"
	"strconv"
)

type BlogService struct {
	blogRepo    *repo.BlogRepo
	commentRepo *repo.CommentRepo
}

func NewBlogService(blogRepo *repo.BlogRepo, commentRepo *repo.CommentRepo) BlogService {
	return BlogService{
		blogRepo:    blogRepo,
		commentRepo: commentRepo,
	}
}

func (bs *BlogService) GetSiteStat() (*dto.Statistics, error) {
	return bs.blogRepo.SelectStat()
}

func (bs *BlogService) GetBlog(id int) (*dto.Blog, error) {
	bm, err := bs.blogRepo.SelectOne(id)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	return dto.ConvertBM2BT(&bm)
}

func (bs *BlogService) ListBlog(sepId int, limit int) (map[string]interface{}, error) {
	bms, err := bs.blogRepo.Select(sepId, limit)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	var checkpoint string
	if len(bms) > 0 {
		checkpoint = strconv.Itoa(int(bms[len(bms)-1].ID))
	}
	bts, err := dto.ConvertBMS2BTS(bms)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	return map[string]interface{}{
		"data":       bts,
		"checkpoint": checkpoint,
	}, nil
}

func (bs *BlogService) SearchBlog(keyword string) ([]dto.Blog, error) {
	bms, err := bs.blogRepo.Search(keyword)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	return dto.ConvertBMS2BTS(bms)
}
