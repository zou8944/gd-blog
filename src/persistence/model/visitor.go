package model

const TableNameVisitor = "visitor"

type Visitor struct {
	Model
	Name  string `gorm:"uniqueIndex"`
	IP    string
	Email string
}

func (*Visitor) TableName() string {
	return TableNameVisitor
}
