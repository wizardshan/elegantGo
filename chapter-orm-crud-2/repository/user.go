package repository

import (
	"context"
	"elegantGo/chapter-orm-crud-2/pkg/hashid"
	"elegantGo/chapter-orm-crud-2/repository/ent"
	"elegantGo/chapter-orm-crud-2/repository/ent/user"
)

type User struct {
	repo
}

func NewUser(db *ent.Client) *User {
	repo := new(User)
	repo.db = db
	return repo
}

func (repo *User) Register(ctx context.Context) (*ent.User, error) {
	mobile := "13000000003"
	password := "a906449d5769fa7361d7ecc6aa3f6d28"
	level := 30
	nickname := "昵称3"
	avatar := "头像3.png"
	bio := "个人介绍3"

	var entUser *ent.User
	err := repo.withTx(ctx, repo.db, func(tx *ent.Tx) error {
		db := tx.Client()
		entUser = repo.create(ctx, db, func(opt *ent.UserCreate) {
			opt.SetMobile(mobile).SetPassword(password).SetLevel(level).SetNickname(nickname).SetAvatar(avatar).SetBio(bio)
		})
		var err error
		if entUser.HashID, err = hashid.EncodeUserID(entUser.ID); err != nil {
			return err
		}
		repo.update(ctx, db, func(opt *ent.UserUpdate) {
			opt.SetHashID(entUser.HashID).Where(user.ID(entUser.ID))
		})
		return nil
	})
	return entUser, err
}
