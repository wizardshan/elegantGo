package response

import "time"

type Article struct {
	ID         int       `json:"id"`
	Nickname   string    `json:"nickname"`
	Bio        string    `json:"bio"`
	CreateTime time.Time `json:"createTime"`
}
