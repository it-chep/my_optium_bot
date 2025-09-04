package dal

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		pool: pool,
	}
}

func (r *Repository) UpdateListName(ctx context.Context, listID int64, name string) error {
	sql := `
		update user_lists
		set name = $2
		where id = $1 
	`

	args := []interface{}{
		listID,
		name,
	}

	_, err := r.pool.Exec(ctx, sql, args...)
	return err
}
