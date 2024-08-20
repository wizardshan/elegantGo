package response

import (
    "elegantGo/chapter-auto-gen/domain"
)

func (respUser *User) Mapper(domUser *domain.User) *User {
    if domUser == nil {
        return nil
    }

    user := new(User)
    user.Mobile = domUser.Mobile
    user.Password = domUser.Password
    user.Level = domUser.Level
    user.Nickname = domUser.Nickname
    user.Avatar = domUser.Avatar
    user.Bio = domUser.Bio
    user.ID = domUser.ID
    user.CreateTime = domUser.CreateTime
    user.UpdateTime = domUser.UpdateTime
    user.Comments = user.Comments.Mapper(domUser.Comments)
    user.Posts = user.Posts.Mapper(domUser.Posts)

    return user
}

func (respUsers Users) Mapper(domUsers domain.Users) Users {
    return mapper(domUsers, func(dom *domain.User) (user *User) {
        return user.Mapper(dom)
    })
}

