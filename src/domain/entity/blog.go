package entity

import (
	"gd-blog/src/persistence/model"
)

type Blog model.Blog

func NewBlog(title, content string) (*Blog, error) {
	blog := &Blog{
		Title:        title,
		Summary:      content[:100],
		Content:      content,
		LikeCount:    0,
		CollectCount: 0,
	}
	return blog, nil
}

func (b *Blog) Update(title, content string) {
	b.Title = title
	b.Summary = content[:100]
	b.Content = content
}

func (b *Blog) Like() {
	b.LikeCount += 1
}

func (b *Blog) UnLike() {
	b.LikeCount -= 1
	if b.LikeCount < 0 {
		b.LikeCount = 0
	}
}

func (b *Blog) Collect() {
	b.CollectCount += 1
}
