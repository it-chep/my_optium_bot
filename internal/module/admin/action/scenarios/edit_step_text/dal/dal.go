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

func (d *Dal) UpdateStepText(ctx context.Context, stepID int64, newStepText string) error {
	sql := `
		update scenario_steps set content = $1 where id = $2
	`

	_, err := d.pool.Exec(ctx, sql, newStepText, stepID)
	return err
}
