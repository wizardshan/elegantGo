package repository

import (
	"context"
	"elegantGo/chapter-orm-crud-2/repository/ent"
)

func (repo *User) Update(ctx context.Context, optionFunc func(builder *ent.UserUpdate)) int {
	return repo.update(ctx, repo.db, optionFunc)
}

func (repo *User) update(ctx context.Context, db *ent.Client, optionFunc func(builder *ent.UserUpdate)) int {
	builder := db.User.Update()
	optionFunc(builder)
	return builder.SaveX(ctx)
}
