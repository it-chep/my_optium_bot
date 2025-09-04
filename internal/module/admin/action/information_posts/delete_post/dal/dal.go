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

func (r *Repository) DeletePost(ctx context.Context, postID int64) error {
	sql := `delete from information_posts where id = $1`
	_, err := r.pool.Exec(ctx, sql, postID)
	return err
}

func (r *Repository) DeleteUnsentPost(ctx context.Context, postID int64) error {
	sql := `delete from patient_posts where post_id = $1 and is_received = false`
	_, err := r.pool.Exec(ctx, sql, postID)
	return err
}
