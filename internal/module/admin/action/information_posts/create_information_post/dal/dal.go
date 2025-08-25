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

func (r *Repository) CreateInformationPost(ctx context.Context) error {
	sql := `
		insert into information_posts (name, posts_theme_id, order_in_theme, media_id, content_type_id, post_text)
		values ($1, $2, $3, $4, $5, $6)
	`
	// todo пробросить параметры
	_, err := r.pool.Exec(ctx, sql)
	return err
}
