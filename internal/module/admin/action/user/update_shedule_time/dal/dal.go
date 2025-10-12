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

func (d *Dal) UpdateNextDelay(ctx context.Context, userID, scenarioID int64, nextDelay time.Time) error {
	sql := `
		UPDATE patient_scenarios ps 
		SET scheduled_time = $1 
		FROM patients p 
		WHERE ps.patient_id = p.tg_id 
		AND p.id = $2 and ps.scenario_id = $3;
	`

	_, err := d.pool.Exec(ctx, sql, nextDelay, userID, scenarioID)
	return err
}
