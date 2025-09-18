package dal

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Dal struct {
	pool *pgxpool.Pool
}

func NewDal(pool *pgxpool.Pool) *Dal {
	return &Dal{
		pool: pool,
	}
}

func (d *Dal) UpdateNextDelay(ctx context.Context, userID int64, nextDelay time.Time) error {
	sql := `
		update patient_scenarios set scheduled_time = $1 where patient_id = $2;
	`

	_, err := d.pool.Exec(ctx, sql, nextDelay, userID)
	return err
}
