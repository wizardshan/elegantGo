package repository

import (
	"context"
	"elegantGo/chapter-param-complex-validator-3/repository/ent"
)

type User struct {
	repo
}

func NewUser(db *ent.Client) *User {
	repo := new(User)
	repo.db = db
	return repo
}

func (repo *User) FetchMany(ctx context.Context, optionFunc func(builder *ent.UserQuery)) ent.Users {
	return repo.fetchMany(ctx, repo.db, optionFunc)
}

func (repo *User) fetchMany(ctx context.Context, db *ent.Client, optionFunc func(builder *ent.UserQuery)) ent.Users {
	builder := db.User.Query()
	optionFunc(builder)
	return builder.AllX(ctx)
}
