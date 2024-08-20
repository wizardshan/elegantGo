package repository

import (
	"context"
	"elegantGo/chapter-orm-crud-1/pkg/hashid"
	"elegantGo/chapter-orm-crud-1/repository/ent"
	"elegantGo/chapter-orm-crud-1/repository/ent/user"
)

func (repo *User) Register(ctx context.Context) (*ent.User, error) {
	mobile := "13000000003"
	password := "a906449d5769fa7361d7ecc6aa3f6d28"
	level := 30
	nickname := "昵称3"
	avatar := "头像3.png"
	bio := "个人介绍3"

	entUser := repo.Create(ctx, func(opt *ent.UserCreate) {
		opt.SetMobile(mobile).SetPassword(password).SetLevel(level).SetNickname(nickname).SetAvatar(avatar).SetBio(bio)
	})

	hashID, err := hashid.EncodeUserID(entUser.ID)
	if err != nil {
		return nil, err
	}
	repo.Update(ctx, func(opt *ent.UserUpdate) {
		opt.SetHashID(hashID).Where(user.ID(entUser.ID))
	})

	return entUser, nil
}
