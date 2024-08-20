package response

import (
    "time"
)

type Users []*User

type User struct {
    Mobile string
    Password string
    Level int
    Nickname string
    Avatar string
    Bio string
    ID int
    CreateTime time.Time
    UpdateTime time.Time
    
    Comments Comments `json:",omitempty"`
    
    Posts Posts `json:",omitempty"`
}