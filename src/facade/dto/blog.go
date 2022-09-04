package dto

import (
	"gd-blog/src/custom"
	"gd-blog/src/repo/model"
	"github.com/jinzhu/copier"
	"gopkg.in/errgo.v2/errors"
)

type Blog struct {
	ID           uint            `json:"id"`
	Title        string          `json:"title"`
	Summary      string          `json:"summary"`
	Content      string          `json:"content"`
	Type         model.BlogType  `json:"type"`
	LikeCount    int             `json:"like_count"`
	ViewCount    int             `json:"view_count"`
	CollectCount int             `json:"collect_count"`
	Categories   []Category      `json:"categories"`
	Tags         []Tag           `json:"tags"`
	CreatedAt    custom.UnixTime `json:"created_at"`
	UpdatedAt    custom.UnixTime `json:"updated_at"`
}

func ConvertBM2BT(bm *model.Blog) (*Blog, error) {
	var bt Blog
	err := copier.Copy(&bt, bm)
	return &bt, err
}

func ConvertBMS2BTS(bms []model.Blog) ([]Blog, error) {
	bts := []Blog{}
	for _, bm := range bms {
		var bt Blog
		if err := copier.Copy(&bt, &bm); err != nil {
			return nil, errors.Wrap(err)
		}
		content := []rune(bt.Content)
		if len(content) > 100 {
			bt.Summary = string(content[:100]) + "。。。"
		} else {
			bt.Summary = bt.Content
		}
		bt.Content = ""
		bts = append(bts, bt)
	}
	return bts, nil
}
