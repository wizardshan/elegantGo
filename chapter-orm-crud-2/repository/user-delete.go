package repository

import (
	"context"
	"elegantGo/chapter-orm-crud-2/repository/ent"
)

func (repo *User) Delete(ctx context.Context, optionFunc func(builder *ent.UserDelete)) int {
	return repo.delete(ctx, repo.db, optionFunc)
}

func (repo *User) delete(ctx context.Context, db *ent.Client, optionFunc func(builder *ent.UserDelete)) int {
	builder := db.User.Delete()
	optionFunc(builder)
	return builder.ExecX(ctx)
}
