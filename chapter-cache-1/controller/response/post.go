package response

import (
    "time"
)

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
    
    Comments Comments `json:",omitempty"`
    User *User `json:",omitempty"`
}