package model

import "gorm.io/gorm"

const TableNameComment = "comment"

// Comment mapped from table <comment>
type Comment struct {
	gorm.Model
	BlogID   int32
	Visitor  Visitor `gorm:"many2many:comment_visitors"`
	ReplyID  int32
	GValue   int32
	Content  string
	Approved int32
}

// TableName Comment's table name
func (*Comment) TableName() string {
	return TableNameComment
}
