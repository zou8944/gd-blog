package model

import "gorm.io/gorm"

const TableNameCategory = "category"

type Category struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex" `
	Description string
}

func (*Category) TableName() string {
	return TableNameCategory
}
