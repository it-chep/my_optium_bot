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

// DeletePostTheme удаляем тему
func (r *Repository) DeletePostTheme(ctx context.Context, themeID int64) error {
	sql := `
		delete from posts_themes where id = $1
	`
	_, err := r.pool.Exec(ctx, sql, themeID)
	return err
}

// DropThemeFromPost удаляем тему у поста
func (r *Repository) DropThemeFromPost(ctx context.Context, themeID int64) error {
	sql := `
		update information_posts set posts_theme_id = null where posts_theme_id = $1
	`
	_, err := r.pool.Exec(ctx, sql, themeID)
	return err
}
