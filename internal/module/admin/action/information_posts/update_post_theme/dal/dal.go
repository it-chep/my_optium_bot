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

// UpdatePostTheme обновление темы
func (r *Repository) UpdatePostTheme(ctx context.Context, themeID int64, name string, isRequired bool) error {
	sql := `
		update posts_themes set is_required = $2, name = $3 where id = $1;
	`
	_, err := r.pool.Exec(ctx, sql, themeID, isRequired, name)
	return err
}
