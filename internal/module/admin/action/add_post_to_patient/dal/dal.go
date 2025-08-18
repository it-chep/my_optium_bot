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

func (r *Repository) AddPostToPatient(ctx context.Context, patientID, postID int64) error {
	sql := `
		insert into patient_posts(patient_id, post_id) values ($1, $2)
	`

	_, err := r.pool.Exec(ctx, sql, patientID, postID)

	return err
}
