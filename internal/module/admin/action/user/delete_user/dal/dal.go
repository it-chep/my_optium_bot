package dal

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		pool: pool,
	}
}

func (r *Repository) GetUserTgID(ctx context.Context, userID int64) (int64, error) {
	sql := `
		select tg_id from patients where id = $1
	`
	var tgID int64
	err := r.pool.QueryRow(ctx, sql, userID).Scan(&tgID)
	if err != nil {
		return 0, errors.Wrap(err, "get tgID")
	}
	return tgID, nil
}

func (r *Repository) DeletePatientScenarios(ctx context.Context, tgID int64) error {
	sql := `delete from patient_scenarios where patient_id = $1`
	_, err := r.pool.Exec(ctx, sql, tgID)
	if err != nil {
		return errors.Wrap(err, "delete DeletePatientScenarios")
	}
	return nil
}

func (r *Repository) DeletePatientPosts(ctx context.Context, tgID int64) error {
	sql := `delete from patient_posts where patient_id = $1`
	_, err := r.pool.Exec(ctx, sql, tgID)
	if err != nil {
		return errors.Wrap(err, "delete DeletePatientPosts")
	}
	return nil
}

func (r *Repository) DeletePatientDoctor(ctx context.Context, userID int64) error {
	sql := `delete from patient_doctor where patient_id = $1`
	_, err := r.pool.Exec(ctx, sql, userID)
	if err != nil {
		return errors.Wrap(err, "delete DeletePatientDoctor")
	}
	return nil
}

func (r *Repository) DeleteUser(ctx context.Context, userID int64) error {
	sql := `delete from patients where id = $1`
	_, err := r.pool.Exec(ctx, sql, userID)
	if err != nil {
		return errors.Wrap(err, "delete DeleteUser")
	}
	return nil
}

func (r *Repository) DeleteLists(ctx context.Context, userID int64) error {
	sql := `delete from users_lists where id = $1`
	_, err := r.pool.Exec(ctx, sql, userID)
	if err != nil {
		return errors.Wrap(err, "delete lists")
	}
	return nil
}

func (r *Repository) DeleteRepetitions(ctx context.Context, tgID int64) error {
	sql := `delete from repetitions where patient_tg_id = $1`
	_, err := r.pool.Exec(ctx, sql, tgID)
	if err != nil {
		return errors.Wrap(err, "delete repetitions")
	}
	return nil
}
