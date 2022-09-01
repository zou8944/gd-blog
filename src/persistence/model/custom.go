package model

import (
	"gd-blog/src/custom"
	"gorm.io/gorm"
)

type Model struct {
	ID        uint            `gorm:"primarykey" json:"id"`
	CreatedAt custom.UnixTime `json:"createdAt"`
	UpdatedAt custom.UnixTime `json:"UpdatedAt"`
	DeletedAt gorm.DeletedAt  `gorm:"index" json:"-"`
}
