package repository

import (
	"context"
	"elegantGo/orm-crud-3/pkg/hashid"
	"elegantGo/orm-crud-3/repository/ent"
	"elegantGo/orm-crud-3/repository/ent/user"
)

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
