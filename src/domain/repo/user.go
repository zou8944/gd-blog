package repo

import "gd-blog/src/domain/entity"

type UserRepo interface {
	SelectOne(email string) *entity.User
	Insert(user *entity.User) (*entity.User, error)
}
