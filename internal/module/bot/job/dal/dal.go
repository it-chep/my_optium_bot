package dal

import (
	"context"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dal/dao"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lib/pq"
	"github.com/samber/lo"
)

type JobDal struct {
	pool *pgxpool.Pool
}

func NewDal(pool *pgxpool.Pool) *JobDal {
	return &JobDal{
		pool: pool,
	}
}

func (d *JobDal) GetAvailableScenarios(ctx context.Context) ([]dto.PatientScenario, error) {
	var (
		scenarios = &dao.PatientScenarios{}
		// достаем тех, у кого нет активных сценариев, но есть доступные
		sql = `select *
					from patient_scenarios
				where scheduled_time < $1
  					and completed_at is null
  					and patient_id not in (
  						select patient_id from patient_scenarios where active is true
  					)
				order by scheduled_time 
		`
	)

	if err := pgxscan.Select(ctx, d.pool, scenarios, sql, time.Now().UTC()); err != nil {
		return nil, err
	}

	return scenarios.ToDomain(), nil
}

func (d *JobDal) MarkScenariosActive(ctx context.Context, scenarios []dto.PatientScenario) error {
	var (
		sql = `update patient_scenarios
					set active = true
			   where id = any($1)
  		`
		ids = lo.Map(scenarios, func(sc dto.PatientScenario, _ int) int64 { return sc.ID })
	)

	_, err := d.pool.Exec(ctx, sql, pq.Array(ids))
	return err
}

func (d *JobDal) GetActiveScheduledScenarios(ctx context.Context) ([]dto.PatientScenario, error) {
	var (
		scenarios = &dao.PatientScenarios{}
		// достаем тех, у кого есть активные сценарии, время которых подошло
		sql = `select *
					from patient_scenarios
				where scheduled_time < $1
  					and completed_at is null
				  	and active is true
				  	and answered is true 
				  	and sent is false
--   					and patient_id in (
--   						select patient_id 
--   							from patient_scenarios 
--   						where active is true and answered is true and sent is false and completed_at is null
--   					) 
		`
	)

	if err := pgxscan.Select(ctx, d.pool, scenarios, sql, time.Now().UTC()); err != nil {
		return nil, err
	}

	return scenarios.ToDomain(), nil
}

func (d *JobDal) GetActiveOldScenarios(ctx context.Context, delay time.Duration) ([]dto.PatientScenario, error) {
	var (
		scenarios = &dao.PatientScenarios{}
		// достаем тех, у кого есть активные сценарии, на которые пользователь не отвечает
		sql = `select *
					from patient_scenarios
				where active=true 
				  and scheduled_time < $1
  					and completed_at is null
					and sent is true
		`
	)

	if err := pgxscan.Select(ctx, d.pool, scenarios, sql, time.Now().UTC().Add(-delay)); err != nil {
		return nil, err
	}

	return scenarios.ToDomain(), nil
}

func (d *JobDal) MoveToFuture(ctx context.Context, patientScenarioID int64, scheduled time.Time) error {
	var (
		// двигаем сценарий в будущее, чтобы могли выполняться другие
		sql = `update patient_scenarios 
			   		set scheduled_time = $2, sent = false, active = false, answered = false
				where id = $1	

		`
	)

	_, err := d.pool.Exec(ctx, sql, patientScenarioID, scheduled)
	return err
}
