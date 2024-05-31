package repository

import (
	"app/chapter-orm-crud-2/repository/ent"
	"context"
	"database/sql"
	"github.com/pkg/errors"
)

type repoDB struct {
	db *ent.Client
}

func (repo *repoDB) withTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			err = errors.Wrapf(err, "rolling back transaction: %v", errRollback)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrapf(err, "committing transaction: %v", err)
	}
	return nil
}

func (repo *repoDB) QueryMany(ctx context.Context, query string, args []any, option func(rows *sql.Rows)) error {
	return repo.queryMany(ctx, repo.db, query, args, option)
}

func (repo *repoDB) queryMany(ctx context.Context, db *ent.Client, query string, args []any, option func(rows *sql.Rows)) error {
	rows, err := db.DB().QueryContext(ctx, query, args...)
	if err != nil {
		return err
	}

	for rows.Next() {
		option(rows)
	}

	// Check for errors during rows "Close".
	if closeErr := rows.Close(); closeErr != nil {
		return closeErr
	}

	// Check for errors during row iteration.
	if err = rows.Err(); err != nil {
		return err
	}

	return nil
}

func (repo *repoDB) QueryOne(ctx context.Context, query string, args []any, option func(rows *sql.Rows)) error {
	return repo.queryOne(ctx, repo.db, query, args, option)
}

func (repo *repoDB) queryOne(ctx context.Context, db *ent.Client, query string, args []any, option func(rows *sql.Rows)) error {
	rows, err := db.DB().QueryContext(ctx, query, args...)
	defer rows.Close()
	if err != nil {
		return err
	}

	for rows.Next() {
		option(rows)
	}
	return nil
}
