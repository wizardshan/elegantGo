package model

import (
	"time"
)

const UserStatusCanceled = 20

const (
	UserLevelDescNormal   = 0
	UserLevelDescSilver   = 10
	UserLevelDescGold     = 20
	UserLevelDescPlatinum = 30
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
