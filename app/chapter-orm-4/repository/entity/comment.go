package entity

import (
	"time"
)

type Comments []*Comment

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
