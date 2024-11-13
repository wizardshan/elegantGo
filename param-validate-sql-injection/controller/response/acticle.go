package response

import "time"

type Article struct {
	ID         int
	Nickname   string
	Bio        string
	CreateTime time.Time
}
