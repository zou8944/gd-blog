package repoimpl

import (
	"gd-blog/src/domain/entity"
	"gorm.io/gorm"
)

type UserRepoImpl struct {
	db *gorm.DB
}

func NewUserRepoImpl(db *gorm.DB) UserRepoImpl {
	return UserRepoImpl{db: db}
}

func (UserRepoImpl) SelectOne(email string) *entity.User {
	//TODO implement me
	panic("implement me")
}

func (UserRepoImpl) Insert(user *entity.User) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}
