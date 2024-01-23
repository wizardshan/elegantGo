package entity

import "time"

type Users []*User

type User struct {
	ID         int
	Level      int
	Mobile     string
	Nickname   string
	Avatar     string
	CreateTime time.Time
}
