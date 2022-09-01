package model

const TableNameBlog = "blog"

type Blog struct {
	Model
	Title        string     `json:"title"`
	Summary      string     `json:"summary"`
	Content      string     `json:"content"`
	LikeCount    int32      `json:"likeCount"`
	CollectCount int32      `json:"collectCount"`
	Categories   []Category `gorm:"many2many:blog_categories" json:"categories"`
	Tags         []Tag      `gorm:"many2many:blog_tags" json:"tags"`
}

func (*Blog) TableName() string {
	return TableNameBlog
}
