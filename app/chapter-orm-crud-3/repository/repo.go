package repository

import (
	"app/chapter-orm-crud-3/repository/ent"
	"context"
	"database/sql"
	"github.com/pkg/errors"
)

type repo struct {
	db *ent.Client
}

func (repo *repo) withTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
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

func (repo *repo) QueryMany(ctx context.Context, query string, args []any, optionFunc func(rows *sql.Rows) error) error {
	rows, err := repo.db.DB().QueryContext(ctx, query, args...)
	if err != nil {
		return err
	}

	for rows.Next() {
		err = optionFunc(rows)
		if err != nil {
			break
		}
	}

	// Check for errors during rows "Close".
	if closeErr := rows.Close(); closeErr != nil {
		return closeErr
	}

	// Check for row scan error.
	if err != nil {
		return err
	}

	// Check for errors during row iteration.
	if err = rows.Err(); err != nil {
		return err
	}

	return nil
}

func (repo *repo) QueryOne(ctx context.Context, query string, args []any, optionFunc func(row *sql.Row) error) error {
	row := repo.db.DB().QueryRowContext(ctx, query, args...)
	return optionFunc(row)
}
