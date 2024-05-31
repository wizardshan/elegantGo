package repository

import (
	"app/chapter-orm-crud-1/repository/ent"
	"app/chapter-orm-crud-1/repository/ent/user"
	"context"
)

type User struct {
	db *ent.Client
}

func NewUser(db *ent.Client) *User {
	repo := new(User)
	repo.db = db
	return repo
}

func (repo *User) FetchByID(ctx context.Context, id int) *ent.User {
	return repo.db.User.Query().Where(user.ID(id)).FirstX(ctx)
}

func (repo *User) FetchByMobile(ctx context.Context, mobile string) *ent.User {
	return repo.db.User.Query().Where(user.Mobile(mobile)).FirstX(ctx)
}

func (repo *User) FetchByMobileAndPassword(ctx context.Context, mobile string, password string) *ent.User {
	return repo.db.User.Query().Where(user.Mobile(mobile), user.Password(password)).FirstX(ctx)
}

func (repo *User) FetchByNickname(ctx context.Context, nickname string) ent.Users {
	return repo.db.User.Query().Where(user.Nickname(nickname)).AllX(ctx)
}
