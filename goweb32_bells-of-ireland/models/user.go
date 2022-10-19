package models

import "reflect"

type User struct {
	UserID   int64  `db:"user_id" json:"user_id,string"`
	Email    string `db:"email" json:"email"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
}

func (u User) IsEmpty() bool {
	return reflect.DeepEqual(u, User{})
}
