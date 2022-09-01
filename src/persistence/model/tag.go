package model

const TableNameLabel = "tag"

type Tag struct {
	Model
	Name string `gorm:"uniqueIndex" json:"name"`
}

func (*Tag) TableName() string {
	return TableNameLabel
}
