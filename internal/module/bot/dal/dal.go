package dal

import (
	"context"
	"fmt"
	"time"

	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto/admin"
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

func (d *CommonDal) UpdateDoctorStep(ctx context.Context, doctorID, chatID int64, step dto.Step) error {
	sql := `insert into doctors_scenarios (doctor_id, scenario_id, step, chat_id) 
				values ($1, $2, $3, $4)
			on conflict (doctor_id, scenario_id, chat_id) do update
				set step = excluded.step
			`
	args := []interface{}{
		doctorID,
		step.ScenarioID,
		step.Order,
		chatID,
	}

	_, err := d.pool.Exec(ctx, sql, args...)
	return err
}

func (d *CommonDal) DoctorNextStep(ctx context.Context, usr user.User, chatID int64) (dto.Step, error) {
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
		sql = `update doctors_scenarios set completed_at = now() where doctor_id = $1 and scenario_id = $2 and chat_id = $3`
		_, err := d.pool.Exec(ctx, sql, usr.ID, usr.StepStat.ScenarioID, chatID)
		if err != nil {
			return dto.Step{}, err
		}

		return dto.Step{}, nil
	}

	logger.Message(ctx, fmt.Sprintf("Двигаем пользователя на шаг вперед. Тек Сценарий: %d, шаг: %d", usr.StepStat.ScenarioID, usr.StepStat.StepOrder))
	// Двигаем пользователя на 1 шаг вперед
	sql = `update doctors_scenarios set step = $1 where doctor_id = $2 and chat_id = $3`
	nextStep := usr.StepStat.StepOrder + 1

	_, err := d.pool.Exec(ctx, sql, nextStep, usr.ID, chatID)
	if err != nil {
		return dto.Step{}, err
	}

	return d.GetStep(ctx, usr.StepStat.ScenarioID, nextStep)
}

func (d *CommonDal) GetUser(ctx context.Context, id, chatID int64) (_ user.User, err error) {
	var (
		usr       = &dao.User{}
		doctorSql = `
			select true as is_doctor, d.tg_id as id, ds.scenario_id, ds.step as step_order
				from doctors d
         			left join doctors_scenarios ds on d.tg_id = ds.doctor_id
			where d.tg_id = $1 and ds.completed_at is null and ds.chat_id = $2
		`
		patientSql = `
			select false as is_doctor, p.tg_id, ps.scenario_id, ps.step as step_order
				from patients p
         			left join patient_scenarios ps on p.tg_id = ps.patient_id and ps.chat_id = $2 and active=true
			where p.tg_id = $1
		`
	)

	if err = pgxscan.Get(ctx, d.pool, usr, doctorSql, id, chatID); err == nil {
		return usr.ToDomain(), nil
	}

	if err = pgxscan.Get(ctx, d.pool, usr, patientSql, id, chatID); err == nil {
		return usr.ToDomain(), nil
	}

	return user.User{}, err
}

func (d *CommonDal) GetDoctorAddMedia(ctx context.Context, tgID int64) (user.User, error) {
	sql := `
		select d.tg_id as id, ds.scenario_id, ds.step as step_order
				from doctors d
         			left join doctors_scenarios ds on d.tg_id = ds.doctor_id
			where d.tg_id = $1 and ds.scenario_id = 11;
		`

	var userDao dao.User
	err := pgxscan.Get(ctx, d.pool, &userDao, sql, tgID)
	if err != nil {
		return user.User{}, err
	}

	return userDao.ToDomain(), nil
}

func (d *CommonDal) GetPatient(ctx context.Context, tgID int64) (user.Patient, error) {
	sql := `
		select p.* 
		from patients p  
		where p.tg_id = $1
		`

	var patient dao.Patient
	err := pgxscan.Get(ctx, d.pool, &patient, sql, tgID)
	if err != nil {
		return user.Patient{}, err
	}

	return patient.ToDomain(), nil
}

func (d *CommonDal) AssignScenarios(ctx context.Context, patient, chatID int64, scenarios []dto.Scenario) error {
	var (
		sql = `insert into patient_scenarios (patient_id, step, chat_id, scenario_id, scheduled_time) 
				select 
				    $1,
				    1, 
				    $2,
				    unnest($3::bigint[]) as scenario_id,
				    unnest($4::timestamp[]) as scheduled_time
		`
		args = []interface{}{
			patient,
			chatID,
			pq.Array(lo.Map(scenarios, func(s dto.Scenario, _ int) int64 { return s.ID })),
			pq.Array(lo.Map(scenarios, func(s dto.Scenario, _ int) time.Time { return s.ScheduledTime.UTC() })),
		}
	)

	_, err := d.pool.Exec(ctx, sql, args...)
	return err
}

type MoveStep struct {
	TgID, ChatID, Scenario int64
	Step, NextStep         int
	Delay                  time.Duration

	// костыль
	Answered bool
	Sent     bool
}

func (d *CommonDal) MoveStepPatient(ctx context.Context, moveStep MoveStep) error {
	sql := `update patient_scenarios set step = $1, answered = $9, scheduled_time = $2, sent = $7, active = $8 
                where patient_id = $3 and chat_id = $4 and scenario_id = $5 and step = $6`
	args := []interface{}{
		moveStep.NextStep,                    // $1
		time.Now().UTC().Add(moveStep.Delay), // $2
		moveStep.TgID,                        // $3
		moveStep.ChatID,                      // $4
		moveStep.Scenario,                    // $5
		moveStep.Step,                        // $6
		moveStep.Sent,                        // $7
		lo.Ternary(moveStep.Delay > time.Hour, false, true),
		moveStep.Answered,
	}
	_, err := d.pool.Exec(ctx, sql, args...)
	return err
}

func (d *CommonDal) CompleteScenario(ctx context.Context, tgID, chatID int64, scenarioID int64) error {
	sql := `update patient_scenarios set completed_at=now(), active = false where patient_id = $1 and chat_id = $2 and scenario_id = $3`
	_, err := d.pool.Exec(ctx, sql, tgID, chatID, scenarioID)
	return err
}

func (d *CommonDal) MarkScenariosSent(ctx context.Context, scenarios ...dto.PatientScenario) error {
	var (
		sql = `update patient_scenarios
					set sent = true
			   where id = any($1)
  		`
		ids = lo.Map(scenarios, func(sc dto.PatientScenario, _ int) int64 { return sc.ID })
	)

	_, err := d.pool.Exec(ctx, sql, pq.Array(ids))
	return err
}

func (d *CommonDal) GetAdminMessages(ctx context.Context, scenario, step int64) (admin.MessageAdmin, error) {
	var (
		chatIDs  []int64
		messages []string

		sqlChats = `select chat_id from admin_chat`
		sqlMess  = `select message 
					from admin_messages 
				   where scenario_id = $1 and next_step = $2`
	)

	if err := pgxscan.Select(ctx, d.pool, &chatIDs, sqlChats); err != nil {
		return admin.MessageAdmin{}, err
	}

	if err := pgxscan.Select(ctx, d.pool, &messages, sqlMess, scenario, step); err != nil {
		return admin.MessageAdmin{}, err
	}

	return admin.MessageAdmin{
		ChatIDs:  chatIDs,
		Messages: messages,
	}, nil
}

func (d *CommonDal) GetDoctorMessages(ctx context.Context, patientTg, scenario, step int64) (admin.MessageDoctor, error) {
	var (
		doctorMessage admin.MessageDoctor

		sqlChats = `select doctor_tg 
						from patient_doctor 
					where patient_tg = $1`
		sqlMess = `select message 
						from doctor_messages 
				   	where scenario_id = $1 and next_step = $2`
	)

	if err := pgxscan.Get(ctx, d.pool, &doctorMessage.DoctorID, sqlChats, patientTg); err != nil {
		return doctorMessage, err
	}

	if err := pgxscan.Select(ctx, d.pool, &doctorMessage.Messages, sqlMess, scenario, step); err != nil {
		return doctorMessage, err
	}

	return doctorMessage, nil
}

func (d *CommonDal) AssignInformationPosts(ctx context.Context, patientTg int64) error {
	sql := `
		insert into patient_posts (patient_id, post_id) 
		select $1, ip.id 
		from information_posts ip 
			join posts_themes pt on ip.posts_theme_id = pt.id 
		where pt.is_required = true	
	`

	_, err := d.pool.Exec(ctx, sql, patientTg)
	return err
}

func (d *CommonDal) UpdateLastCommunication(ctx context.Context, patientID int64) error {
	var (
		sql = `update patients
					set last_communicate = $2
			   where id = $1
  		`
	)

	_, err := d.pool.Exec(ctx, sql, patientID, time.Now().UTC())
	return err
}

func (d *CommonDal) MoveScenario(ctx context.Context, patientScenarioID int64, schedTime time.Time) error {
	var (
		sql = `update patient_scenarios
					set scheduled_time = $2
			   where id = $1
  		`
	)

	_, err := d.pool.Exec(ctx, sql, patientScenarioID, schedTime.UTC())
	return err
}

func (d *CommonDal) ScenarioNotAnswered(ctx context.Context, patientTGID, scenarioID int64) error {
	sql := `update patient_scenarios set answered = false where patient_id = $1 and scenario_id = $2`
	_, err := d.pool.Exec(ctx, sql, patientTGID, scenarioID)
	return err
}
