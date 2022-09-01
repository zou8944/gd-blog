package model

const TableNameComment = "comment"

// Comment mapped from table <comment>
type Comment struct {
	Model
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
