package model

import "gorm.io/gorm"

const TableNameVisitor = "visitor"

type Visitor struct {
	gorm.Model
	Name  string
	IP    string
	Email string
}

func (*Visitor) TableName() string {
	return TableNameVisitor
}
