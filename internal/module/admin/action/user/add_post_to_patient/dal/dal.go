package dal

import (
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
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

func (r *Repository) AddPostToPatient(ctx context.Context, userID, postID int64) error {
	sqlPatient := `select tg_id from patients where id = $1`
	var tgId int64
	pgxscan.Get(ctx, r.pool, &tgId, sqlPatient, userID)

	sql := `
		insert into patient_posts(patient_id, post_id) values ($1, $2)
	`

	_, err := r.pool.Exec(ctx, sql, tgId, postID)

	return err
}
