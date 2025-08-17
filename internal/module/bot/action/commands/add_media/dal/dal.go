package dal

import (
	"context"

	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
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

func (d *Dal) UpdateDoctorStep(ctx context.Context, doctorID int64, step dto.Step) error {
	sql := `insert into doctors_scenarios (doctor_id, scenario_id, step) 
				values ($1, $2, $3)
			on conflict (doctor_id, scenario_id) do update
				set step = excluded.step
			`
	args := []interface{}{
		doctorID,
		step.ScenarioID,
		step.Order,
	}

	_, err := d.pool.Exec(ctx, sql, args...)
	return err
}
