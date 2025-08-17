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

func (d *Dal) ExitScenario(ctx context.Context, userID int64) error {
	sql := `
		delete from doctors_scenarios where doctor_id = $1 and scenario_id = 11;
	`
	_, err := d.pool.Exec(ctx, sql, userID)
	if err != nil {
		return err
	}

	return nil
}
