package dal

import (
	"context"

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

func (d *CommonDal) GetUser(ctx context.Context, id int64) (_ user.User, err error) {
	var (
		usr       = &dao.User{}
		doctorSql = `
			select true as is_doctor, d.tg_id as id, ds.scenario_id, ds.step as step_order
				from doctors d
         			left join doctors_scenarios ds on d.tg_id = ds.doctor_id
			where d.tg_id = $1
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
