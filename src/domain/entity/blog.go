package entity

import "time"

type Blog struct {
	Id           int
	Title        string
	Summary      string
	Content      string
	LikeCount    int
	CollectCount int
	Scores       []string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewBlog(title, content string) (*Blog, error) {
	blog := &Blog{
		Title:        title,
		Summary:      content[:100],
		Content:      content,
		LikeCount:    0,
		CollectCount: 0,
		Scores:       []string{},
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
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
