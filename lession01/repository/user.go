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
		Amount:     1100,
		Mobile:     "13000000001",
		Nickname:   "user1",
		Avatar:     "avatar-default.png",
		CreateTime: time.Now(),
	}

	user2 := &model.User{
		ID:         2,
		Level:      20,
		Amount:     2000,
		Mobile:     "13000000002",
		Nickname:   "user2",
		Avatar:     "avatar-default.png",
		CreateTime: time.Now(),
	}

	var users model.Users
	users = append(users, user1, user2)
	return users
}
