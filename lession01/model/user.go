package model

import (
	"time"
)

type User struct {
	ID         int
	Level      int
	Mobile     string
	Nickname   string
	Avatar     string
	Balance    int
	Status     int
	LoginTime  time.Time
	CreateTime time.Time
}
