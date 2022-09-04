package dto

import "gd-blog/src/custom"

type Comment struct {
	ID             uint            `json:"id"`
	Visitor        Visitor         `json:"visitor"`
	BlogId         int             `json:"blog_id"`
	ReplyCommentId int             `json:"reply_comment_id"`
	Content        string          `json:"content"`
	CreatedAt      custom.UnixTime `json:"created_at"`
	Approved       bool            `json:"-"`
}

func NewComment(user Visitor, blogId int, replyId int, content string) *Comment {
	comment := &Comment{
		Visitor:        user,
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
