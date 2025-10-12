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

// DeleteNewsletter удаление рассылки
func (r *Repository) DeleteNewsletter(ctx context.Context, newsletterID int64) error {
	sql := `DELETE FROM newsletters WHERE id = $1 AND status_id = 1`
	// удаляем только черновики
	_, err := r.pool.Exec(ctx, sql, newsletterID)
	return err
}
