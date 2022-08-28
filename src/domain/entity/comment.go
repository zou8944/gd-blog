package entity

type Comment struct {
	Id             int
	User           *User
	BlogId         int
	ReplyCommentId int
	Content        string
	Approved       bool
}

func NewComment(user *User, blogId int, replyId int, content string) *Comment {
	comment := &Comment{
		User:           user,
		BlogId:         blogId,
		ReplyCommentId: replyId,
		Content:        content,
		Approved:       false,
	}
	return comment
}

func (c *Comment) Update(content string) {
	c.Content = content
}

func (c *Comment) Approve() {
	c.Approved = true
}

func (c *Comment) UnApprove() {
	c.Approved = false
}
