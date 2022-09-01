package model

import "gorm.io/gorm"

const TableNameLabel = "tag"

type Tag struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`
}

func (*Tag) TableName() string {
	return TableNameLabel
}
