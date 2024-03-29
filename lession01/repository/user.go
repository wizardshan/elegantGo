package repository

import (
	"elegantGo/lession01/model"
	"time"
)

type User struct {
}

func NewUser() *User {
	repo := new(User)
	return repo
}

func (repo *User) All() model.Users {
	user1 := &model.User{
		ID:         1,
		Level:      10,
		Balance:    1000,
		Mobile:     "MTMwMDAwMDAwMDE=",
		Nickname:   "user1",
		Avatar:     "avatar-default.png",
		Status:     0,
		CreateTime: time.Now(),
	}

	user2 := &model.User{
		ID:         2,
		Level:      20,
		Balance:    2000,
		Mobile:     "MTMwMDAwMDAwMDI=",
		Nickname:   "user2",
		Avatar:     "avatar-default.png",
		Status:     0,
		CreateTime: time.Now(),
	}

	var users model.Users
	users = append(users, user1, user2)
	return users
}
