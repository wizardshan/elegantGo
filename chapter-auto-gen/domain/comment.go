package domain

import "time"

type Comments []*Comment

type Comment struct {
	UserID int
	PostID int
	Content string
	ID int
	CreateTime time.Time
	UpdateTime time.Time
	Post *Post
	User *User
}