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

func (r *Repository) DeletePostFromPatient(ctx context.Context, patientID, postID int64) error {
	sql := `
		delete from patient_posts where patient_id = $1 and post_id = $2
	`

	_, err := r.pool.Exec(ctx, sql, patientID, postID)
	return err
}
