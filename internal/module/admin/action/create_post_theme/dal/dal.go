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

func (r *Repository) CreatePostTheme(ctx context.Context, name string, isRequired bool) error {
	sql := `
		insert into posts_themes (name, is_required) values ($1, $2)
	`
	_, err := r.pool.Exec(ctx, sql, name, isRequired)
	return err
}
