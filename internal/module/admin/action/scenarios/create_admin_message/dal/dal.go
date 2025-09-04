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

func (d *Dal) CreateAdminMessage(ctx context.Context, scenarioID, stepOrder int64, message string) error {
	sql := `
		with step as (select id
					  from scenario_steps
					  where scenario_id = $1
						and step_order = $2)
		insert
		into admin_messages (scenario_id, next_step, message)
		select $1, id, $3
		from step
	`

	_, err := d.pool.Exec(ctx, sql, scenarioID, stepOrder, message)
	return err
}

func (d *Dal) CreateDoctorMessage(ctx context.Context, scenarioID, stepOrder int64, message string) error {
	sql := `
		with step as (select id
					  from scenario_steps
					  where scenario_id = $1
						and step_order = $2)
		insert
		into doctor_messages (scenario_id, next_step, message)
		select $1, id, $3
		from step
	`

	_, err := d.pool.Exec(ctx, sql, scenarioID, stepOrder, message)
	return err
}
