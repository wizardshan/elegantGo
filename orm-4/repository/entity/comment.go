package entity

import (
	"time"
)

type Comment struct {
	ID         int
	UserID     int
	PostID     int
	Content    string
	CreateTime time.Time
	UpdateTime time.Time

	User *User
	Post *Post
}
