package domain

import "time"

type Users []*User

type User struct {
	ID int
	CreateTime time.Time
	UpdateTime time.Time
	HashID string
	Mobile string
	Password string
	Level int
	Nickname string
	Avatar string
	Bio string
}