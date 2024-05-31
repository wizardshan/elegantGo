package repository

import (
	"app/chapter-orm-crud-1/pkg/hashid"
	"app/chapter-orm-crud-1/repository/ent"
	"app/chapter-orm-crud-1/repository/ent/user"
	"context"
)

func (repo *User) Register(ctx context.Context) *ent.User {
	mobile := "13000000003"
	password := "a906449d5769fa7361d7ecc6aa3f6d28"
	level := 30
	nickname := "昵称3"
	avatar := "头像3.png"
	bio := "个人介绍3"

	entUser := repo.Create(ctx, func(builder *ent.UserCreate) {
		builder.SetMobile(mobile).SetPassword(password).SetLevel(level).SetNickname(nickname).SetAvatar(avatar).SetBio(bio)
	})

	hashID := hashid.EncodePostID(entUser.ID)
	repo.Update(ctx, func(builder *ent.UserUpdate) {
		builder.SetHashID(hashID).Where(user.ID(entUser.ID))
	})

	return entUser
}
