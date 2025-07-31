package dal

import (
	"context"
	"fmt"
	"time"

	"github.com/it-chep/my_optium_bot.git/internal/pkg/logger"
	"github.com/lib/pq"
	"github.com/samber/lo"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dal/dao"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto/user"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CommonDal struct {
	pool *pgxpool.Pool
}

func NewDal(pool *pgxpool.Pool) *CommonDal {
	return &CommonDal{
		pool: pool,
	}
}

func (d *CommonDal) GetScenario(ctx context.Context, id int64) (dto.Scenario, error) {
	scenario := &dao.Scenario{}
	if err := pgxscan.Get(ctx, d.pool, scenario, `select * from scenarios where id = $1`, id); err != nil {
		return dto.Scenario{}, err
	}

	steps := &dao.Steps{}
	if err := pgxscan.Select(ctx, d.pool, steps, `select * from scenario_steps where scenario_id = $1 order by step_order`, id); err != nil {
		return dto.Scenario{}, err
	}

	buttons := &dao.Buttons{}
	if err := pgxscan.Select(ctx, d.pool, buttons, `select * from step_buttons where scenario = $1 order by id`, id); err != nil {
		return dto.Scenario{}, err
	}

	return scenario.ToDomain(steps, buttons), nil
}

func (d *CommonDal) GetStep(ctx context.Context, scenarioID, stepID int64) (dto.Step, error) {
	args := []interface{}{scenarioID, stepID}

	sql := `select * from scenario_steps where scenario_id = $1 and step_order = $2`
	var step dao.Step
	if err := pgxscan.Get(ctx, d.pool, &step, sql, args...); err != nil {
		return dto.Step{}, err
	}

	sql = `select * from step_buttons where scenario = $1 and step = $2 order by id`
	var buttons dao.Buttons
	if err := pgxscan.Select(ctx, d.pool, &buttons, sql, args...); err != nil {
		return dto.Step{}, err
	}

	return step.ToDomain(buttons), nil
}

func (d *CommonDal) UpdateDoctorStep(ctx context.Context, doctorID int64, step dto.Step) error {
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

func (d *CommonDal) DoctorNextStep(ctx context.Context, usr user.User) (dto.Step, error) {
	logger.Message(ctx, fmt.Sprintf("Получение шагов для определение финального. Тек Сценарий: %d, шаг: %d", usr.StepStat.ScenarioID, usr.StepStat.StepOrder))
	sql := `select * from scenario_steps where scenario_id = $1 order by step_order`

	var steps dao.Steps
	if err := pgxscan.Select(ctx, d.pool, &steps, sql, usr.StepStat.ScenarioID); err != nil {
		return dto.Step{}, err
	}

	// Проверяем что текущий шаг не финальный, тк если текущий шаг финальный, то надо менять сценарии
	currentIsFinal := false
	for _, step := range steps {
		if int64(step.ID) == usr.StepStat.StepOrder && step.IsFinal.Bool {
			currentIsFinal = true
		}
	}

	if currentIsFinal {
		logger.Message(ctx, "У пользователя финальный шаг, не двигаем его, завершаем сценарий")
		sql = `update doctors_scenarios set completed_at = now() where doctor_id = $1 and scenario_id = $2`
		_, err := d.pool.Exec(ctx, sql, usr.ID, usr.StepStat.ScenarioID)
		if err != nil {
			return dto.Step{}, err
		}

		return dto.Step{}, nil
	}

	logger.Message(ctx, fmt.Sprintf("Двигаем пользователя на шаг вперед. Тек Сценарий: %d, шаг: %d", usr.StepStat.ScenarioID, usr.StepStat.StepOrder))
	// Двигаем пользователя на 1 шаг вперед
	sql = `update doctors_scenarios set step = $1 where doctor_id = $2`
	nextStep := usr.StepStat.StepOrder + 1

	_, err := d.pool.Exec(ctx, sql, nextStep, usr.ID)
	if err != nil {
		return dto.Step{}, err
	}

	return d.GetStep(ctx, usr.StepStat.ScenarioID, nextStep)
}

func (d *CommonDal) GetUser(ctx context.Context, id int64) (_ user.User, err error) {
	var (
		usr       = &dao.User{}
		doctorSql = `
			select true as is_doctor, d.tg_id as id, ds.scenario_id, ds.step as step_order
				from doctors d
         			left join doctors_scenarios ds on d.tg_id = ds.doctor_id
			where d.tg_id = $1 and ds.completed_at is null
		`
		patientSql = `
			select false as is_doctor, p.tg_id, ps.scenario_id, ps.step as step_order
				from patients p
         			left join patient_scenarios ps on p.tg_id = ps.patient_id
			where p.tg_id = $1
		`
	)

	if err = pgxscan.Get(ctx, d.pool, usr, doctorSql, id); err == nil {
		return usr.ToDomain(), nil
	}

	if err = pgxscan.Get(ctx, d.pool, usr, patientSql, id); err == nil {
		return usr.ToDomain(), nil
	}

	return user.User{}, err
}

func (d *CommonDal) AssignScenarios(ctx context.Context, patient int64, scenarios []dto.Scenario) error {
	var (
		sql = `insert into patient_scenarios (patient_id, step, scenario_id, scheduled_time) 
				select 
				    $1,
				    0, 
				    unnest($2::bigint[]) as scenario_id,
				    unnest($3::timestamp[]) as scheduled_time
		`
		args = []interface{}{
			patient,
			pq.Array(lo.Map(scenarios, func(s dto.Scenario, _ int) int64 { return s.ID })),
			pq.Array(lo.Map(scenarios, func(s dto.Scenario, _ int) time.Time { return s.ScheduledTime })),
		}
	)

	_, err := d.pool.Exec(ctx, sql, args...)
	return err
}
