package domain

import "time"

type Posts []*Post

type Post struct {
	HashID string
	UserID int
	Title string
	Content string
	TimesOfRead int
	ID int
	CreateTime time.Time
	UpdateTime time.Time
	
	Comments Comments
	User *User
}