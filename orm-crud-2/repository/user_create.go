package repository

import (
	"context"
	"elegantGo/orm-crud-2/repository/ent"
)

func (repo *User) Create(ctx context.Context, optionFunc func(opt *ent.UserCreate)) *ent.User {
	return repo.create(ctx, repo.db, optionFunc)
}

func (repo *User) create(ctx context.Context, db *ent.Client, optionFunc func(opt *ent.UserCreate)) *ent.User {
	builder := db.User.Create()
	optionFunc(builder)
	return builder.SaveX(ctx)
}
