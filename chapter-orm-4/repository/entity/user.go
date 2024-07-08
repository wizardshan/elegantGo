package entity

import (
	"time"
)

type Users []*User

type User struct {
	ID         int
	Mobile     string
	Password   string
	Nickname   string
	Avatar     string
	Bio        string
	CreateTime time.Time
	UpdateTime time.Time
}
