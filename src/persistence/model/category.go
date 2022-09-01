package model

const TableNameCategory = "category"

type Category struct {
	Model
	Name        string `gorm:"uniqueIndex" json:"name"`
	Description string `json:"description"`
}

func (*Category) TableName() string {
	return TableNameCategory
}
