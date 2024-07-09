package response

import (
    "elegantGo/chapter-auto-gen/domain"
    "time"
)

type Users []*User

type User struct {
    ID int
    CreateTime time.Time
    UpdateTime time.Time
    HashID string
    Mobile string
    Password string
    Level int
    Nickname string
    Avatar string
    Bio string
}

func MapperUser(domUser *domain.User) *User {
    if domUser == nil {
        return nil
    }

    respUser := new(User)
    respUser.ID = domUser.ID
    respUser.CreateTime = domUser.CreateTime
    respUser.UpdateTime = domUser.UpdateTime
    respUser.HashID = domUser.HashID
    respUser.Mobile = domUser.Mobile
    respUser.Password = domUser.Password
    respUser.Level = domUser.Level
    respUser.Nickname = domUser.Nickname
    respUser.Avatar = domUser.Avatar
    respUser.Bio = domUser.Bio
    return respUser
}

func MapperUsers(domainUsers domain.Users) Users {
    return mapper(domainUsers, MapperUser)
}
