package app

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Временное решение для создания бд
func CreateDb(ctx context.Context, pool *pgxpool.Pool, createDBs ...func(context.Context, *pgxpool.Pool) error) error {
	for _, fn := range createDBs {
		if err := fn(ctx, pool); err != nil {
			return err
		}
	}

	return nil
}
