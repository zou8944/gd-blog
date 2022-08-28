package entity

type User struct {
	id    int
	name  string
	email string
}

func NewUser(name, email string) *User {
	return &User{
		name:  name,
		email: email,
	}
}
