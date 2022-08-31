package model

import "gorm.io/gorm"

const TableNameLabel = "tag"

type Tag struct {
	gorm.Model
	Name string
}

func (*Tag) TableName() string {
	return TableNameLabel
}
