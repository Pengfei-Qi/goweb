package models

import "reflect"

type User struct {
	UserID   int64  `db:"user_id"`
	Email    string `db:"email"`
	Username string `db:"username"`
	Password string `db:"password"`
}

func (u User) IsEmpty() bool {
	return reflect.DeepEqual(u, User{})
}
