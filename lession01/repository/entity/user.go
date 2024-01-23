package entity

import "time"

type Users []*User

type User struct {
	ID         int
	Level      int
	Mobile     string
	Nickname   string
	Avatar     string
	Amount     int
	LoginTime  time.Time
	CreateTime time.Time
}
