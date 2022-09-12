package dto

import (
	"gd-blog/custom"
	"gd-blog/repo/model"
	"github.com/jinzhu/copier"
)

type Comment struct {
	ID             uint            `json:"id"`
	Visitor        Visitor         `json:"visitor"`
	BlogId         int             `json:"blog_id"`
	ReplyCommentId int             `json:"reply_comment_id"`
	Content        string          `json:"content"`
	CreatedAt      custom.UnixTime `json:"created_at"`
	Approved       bool            `json:"-"`
}

func (c *Comment) ToModel() (*model.Comment, error) {
	var mc model.Comment
	err := copier.Copy(&mc, c)
	return &mc, err
}

func ConvertCM2CT(cm *model.Comment) (*Comment, error) {
	var ct Comment
	err := copier.Copy(&ct, cm)
	return &ct, err
}

func ConvertCMS2CTS(cms []model.Comment) ([]Comment, error) {
	cts := []Comment{}
	for _, cm := range cms {
		var ct Comment
		err := copier.Copy(&ct, &cm)
		if err != nil {
			return nil, err
		}
		cts = append(cts, ct)
	}
	return cts, nil
}
