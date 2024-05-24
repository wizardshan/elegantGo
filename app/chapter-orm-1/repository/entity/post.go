package entity

import (
	"time"
)

type Posts []*Post

type Post struct {
	ID          int
	HashID      string
	UserID      int
	Title       string
	Content     string
	TimesOfRead int
	CreateTime  time.Time
	UpdateTime  time.Time

	Nickname string
	Avatar   string

	Comments Comments
}
