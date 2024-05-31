package repository

import (
	"app/chapter-orm-crud-1/repository/ent"
	"context"
)

func (repo *User) Delete(ctx context.Context, optionFunc func(builder *ent.UserDelete)) int {
	return repo.delete(ctx, optionFunc)
}

func (repo *User) delete(ctx context.Context, optionFunc func(builder *ent.UserDelete)) int {
	builder := repo.db.User.Delete()
	optionFunc(builder)
	return builder.ExecX(ctx)
}
