package repository

import (
	"context"
	"elegantGo/orm-crud-1/repository/ent"
)

func (repo *User) Create(ctx context.Context, optionFunc func(opt *ent.UserCreate)) *ent.User {
	builder := repo.db.User.Create()
	optionFunc(builder)
	return builder.SaveX(ctx)
}
