package dal

import (
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/user/move_2_step/dto"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lib/pq"
	"github.com/samber/lo"
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

type PatientIds struct {
	PatientID int64 `db:"tg_id"`
	ChatID    int64
}

func (d *Dal) GetPatientIDs(ctx context.Context, userID int64) (patientIDs PatientIds, err error) {
	sql := `
		select p.tg_id, pd.chat_id from patients p join patient_doctor pd on p.id = pd.patient_id where p.id = $1		
	`

	err = pgxscan.Get(ctx, d.pool, &patientIDs, sql, userID)
	return patientIDs, err
}

func (d *Dal) MoveScenarios(ctx context.Context, patientIDs PatientIds, scenarios []dto.Scenario) error {
	var (
		// ne chat gpt

		sql = `
			UPDATE patient_scenarios 
			SET step = 1,
				scheduled_time = data.new_scheduled_time,
				completed_at = null,
				sent = false,
				answered = false
			FROM (
				SELECT 
					unnest($3::bigint[]) as scenario_id,
					unnest($4::timestamp[]) as new_scheduled_time
			) as data
			WHERE patient_scenarios.patient_id = $1 
			  AND patient_scenarios.chat_id = $2 
			  AND patient_scenarios.scenario_id = data.scenario_id;
		`
		args = []interface{}{
			patientIDs.PatientID,
			patientIDs.ChatID,
			pq.Array(lo.Map(scenarios, func(s dto.Scenario, _ int) int64 { return s.ID })),
			pq.Array(lo.Map(scenarios, func(s dto.Scenario, _ int) time.Time { return s.ScheduledTime.UTC() })),
		}
	)

	_, err := d.pool.Exec(ctx, sql, args...)
	return err
}
