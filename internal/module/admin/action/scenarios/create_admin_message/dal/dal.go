package dal

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Dal struct {
	pool *pgxpool.Pool
}

func NewDal(pool *pgxpool.Pool) *Dal {
	return &Dal{
		pool: pool,
	}
}

func (d *Dal) CreateAdminMessage(ctx context.Context, scenarioID, stepID int64, message string) error {
	sql := `
		insert into admin_messages (scenario_id, next_step, message) 
		values ($1, $2, $3)
	`

	_, err := d.pool.Exec(ctx, sql, scenarioID, stepID, message)
	return err
}

func (d *Dal) CreateDoctorMessage(ctx context.Context, scenarioID, stepID int64, message string) error {
	sql := `
		insert into doctor_messages (scenario_id, next_step, message) 
		values ($1, $2, $3)
	`

	_, err := d.pool.Exec(ctx, sql, scenarioID, stepID, message)
	return err
}
