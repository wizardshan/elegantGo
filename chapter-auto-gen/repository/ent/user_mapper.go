package ent

import (
	"elegantGo/chapter-auto-gen/domain"
	"github.com/elliotchance/pie/v2"
)

func (entUser *User) Mapper() *domain.User {
	if entUser == nil {
		return nil
	}

	domUser := new(domain.User)
	domUser.ID = entUser.ID
	domUser.CreateTime = entUser.CreateTime
	domUser.UpdateTime = entUser.UpdateTime
	domUser.HashID = entUser.HashID
	domUser.Mobile = entUser.Mobile
	domUser.Password = entUser.Password
	domUser.Level = entUser.Level
	domUser.Nickname = entUser.Nickname
	domUser.Avatar = entUser.Avatar
	domUser.Bio = entUser.Bio

	return domUser
}

func (entUsers Users) Mapper() domain.Users {
	return pie.Map(entUsers, func(ent *User) *domain.User {
		return ent.Mapper()
	})
}