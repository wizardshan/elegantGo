package repository

import (
	"context"
	"elegantGo/chapter-orm-crud-1/repository/ent"
	"elegantGo/chapter-orm-crud-1/repository/ent/user"
)

func (repo *User) Fetch(ctx context.Context, id int) *ent.User {
	return repo.FetchOne(ctx, func(builder *ent.UserQuery) {
		builder.Where(user.ID(id))
	})
}

func (repo *User) FetchOne(ctx context.Context, optionFunc func(builder *ent.UserQuery)) *ent.User {
	builder := repo.db.User.Query()
	optionFunc(builder)
	return builder.FirstX(ctx)
}

func (repo *User) FetchMany(ctx context.Context, optionFunc func(builder *ent.UserQuery)) ent.Users {
	builder := repo.db.User.Query()
	optionFunc(builder)
	return builder.AllX(ctx)
}

func (repo *User) Count(ctx context.Context, optionFunc func(builder *ent.UserQuery)) int {
	builder := repo.db.User.Query()
	optionFunc(builder)
	return builder.CountX(ctx)
}

func (repo *User) Exist(ctx context.Context, optionFunc func(builder *ent.UserQuery)) bool {
	builder := repo.db.User.Query()
	optionFunc(builder)
	return builder.ExistX(ctx)
}
