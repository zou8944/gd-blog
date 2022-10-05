package models

type BlogType string

const (
	LongBlog  BlogType = "LongBlog"
	ShortBlog BlogType = "ShortBlog"
	Complain  BlogType = "Complain"
)

type Blog struct {
	BaseModel
	Title        string     `json:"title"`
	Summary      string     `json:"summary"`
	Content      string     `json:"content"`
	Type         BlogType   `json:"type" gorm:"default:LongBlog"`
	LikeCount    int32      `json:"like_count"`
	CollectCount int32      `json:"collect_count"`
	Categories   []Category `json:"categories" gorm:"many2many:blog_categories"`
	Tags         []Tag      `json:"tags" gorm:"many2many:blog_tags"`
}

func ModifyBlogSummary(blogs []Blog) []Blog {
	var newBlogs []Blog
	// blogs得处理
	for i := 0; i < len(blogs); i++ {
		blog := blogs[i]
		content := []rune(blog.Content)
		if len(content) > 100 {
			blog.Summary = string(content[:100]) + "。。。"
		} else {
			blog.Summary = blog.Content
		}
		blog.Content = ""
		newBlogs = append(newBlogs, blog)
	}
	return newBlogs
}

func GetBlogById(id int64) (*Blog, error) {
	var blog Blog
	db.Preload("Categories").Where("id = ?", id).Find(&blog)
	return &blog, db.Error
}

func ListBlog(cid int64, pageNo int, pageSize int) ([]Blog, error) {
	var blogs []Blog
	tx := db.Preload("Categories").Preload("Tags")
	if cid > 0 {
		tx.Raw("select * from blogs b inner join blog_categories bc on b.id = bc.blog_id where bc.category_id = ? order by b.created_at desc limit ? offset ?", cid, pageSize, pageSize*(pageNo-1)).Find(&blogs)
	} else {
		tx.Limit(pageSize).Offset(pageSize * (pageNo - 1)).Order("created_at desc").Find(&blogs)
	}
	if blogs == nil {
		blogs = make([]Blog, 0)
	}
	return blogs, tx.Error
}

func CountBlog(cid int64) (int64, error) {
	var count int64
	if cid > 0 {
		db.Raw("select count(1) from blogs b inner join blog_categories bc on b.id = bc.blog_id where category_id = ?", cid).Find(&count)
	} else {
		db.Model(&Blog{}).Count(&count)
	}
	return count, db.Error
}
