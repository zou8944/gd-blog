package repoimpl

import (
	"database/sql"
	"gd-blog/src/domain/entity"
)

type UserRepoImpl struct {
	db *sql.DB
}

func NewUserRepoImpl(db *sql.DB) UserRepoImpl {
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
