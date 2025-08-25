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

func (d *Dal) UpdateScenarioDelay(ctx context.Context, scenarioID int64, newDelay string) error {
	sql := `
		update scenarios set delay = $1 where id = $2
	`
	_, err := d.pool.Exec(ctx, sql, newDelay, scenarioID)

	return err
}
