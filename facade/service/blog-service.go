package service

import (
	"gd-blog/facade/dto"
	"gd-blog/repo"
	"gopkg.in/errgo.v2/errors"
)

type BlogService interface {
	GetSiteStat() (*dto.Statistics, error)
	GetBlog(id int) (*dto.Blog, error)
	ListBlog(pageNo int, pageSize int, cid int) (map[string]interface{}, error)
	SearchBlog(keyword string) ([]dto.Blog, error)
	ListTag() ([]dto.Tag, error)
	ListCategories() ([]dto.Category, error)
}

type blogService struct {
	blogRepo repo.BlogRepo
}

func NewBlogService(blogRepo repo.BlogRepo) BlogService {
	return &blogService{
		blogRepo: blogRepo,
	}
}

func (bs *blogService) GetSiteStat() (*dto.Statistics, error) {
	return bs.blogRepo.SelectStat()
}

func (bs *blogService) GetBlog(id int) (*dto.Blog, error) {
	bm, err := bs.blogRepo.SelectOne(id)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	return dto.ConvertBM2BT(&bm)
}

func (bs *blogService) ListBlog(pageNo int, pageSize int, cid int) (map[string]interface{}, error) {
	bms, err := bs.blogRepo.Select(pageNo, pageSize, cid)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	count, err := bs.blogRepo.Count()
	if err != nil {
		return nil, errors.Wrap(err)
	}
	bts, err := dto.ConvertBMS2BTS(bms)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	return map[string]interface{}{
		"articles":     bts,
		"articleCount": count,
	}, nil
}

func (bs *blogService) SearchBlog(keyword string) ([]dto.Blog, error) {
	bms, err := bs.blogRepo.Search(keyword)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	return dto.ConvertBMS2BTS(bms)
}

func (bs *blogService) ListTag() ([]dto.Tag, error) {
	return bs.blogRepo.SelectTags()
}

func (bs *blogService) ListCategories() ([]dto.Category, error) {
	return bs.blogRepo.SelectCategories()
}
