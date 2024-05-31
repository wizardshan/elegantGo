package repository

import (
	"app/chapter-orm-crud-1/repository/ent"
	"app/chapter-orm-crud-1/repository/ent/user"
	"context"
)

func (repo *User) Fetch(ctx context.Context, id int) *ent.User {
	return repo.fetchOne(ctx, func(builder *ent.UserQuery) {
		builder.Where(user.ID(id))
	})
}

func (repo *User) FetchOne(ctx context.Context, optionFunc func(builder *ent.UserQuery)) *ent.User {
	return repo.fetchOne(ctx, optionFunc)
}

func (repo *User) FetchMany(ctx context.Context, optionFunc func(builder *ent.UserQuery)) ent.Users {
	return repo.fetchMany(ctx, optionFunc)
}

func (repo *User) Count(ctx context.Context, optionFunc func(builder *ent.UserQuery)) int {
	return repo.count(ctx, optionFunc)
}

func (repo *User) fetchOne(ctx context.Context, optionFunc func(builder *ent.UserQuery)) *ent.User {
	builder := repo.db.User.Query()
	optionFunc(builder)
	return builder.FirstX(ctx)
}

func (repo *User) fetchMany(ctx context.Context, optionFunc func(builder *ent.UserQuery)) ent.Users {
	builder := repo.db.User.Query()
	optionFunc(builder)
	return builder.AllX(ctx)
}

func (repo *User) count(ctx context.Context, optionFunc func(builder *ent.UserQuery)) int {
	builder := repo.db.User.Query()
	optionFunc(builder)
	return builder.CountX(ctx)
}
