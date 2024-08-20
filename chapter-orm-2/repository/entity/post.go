package entity

import (
	"time"
)

type Post struct {
	ID         int
	HashID     string
	UserID     int
	Title      string
	Content    string
	Views      int
	CreateTime time.Time
	UpdateTime time.Time

	UserNickname string
	UserAvatar   string

	Comments []*Comment
}
