package dal

import (
	"context"
	"fmt"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dal/dao"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto/user"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type InitBotDal struct {
	pool *pgxpool.Pool
}

func NewDal(pool *pgxpool.Pool) *InitBotDal {
	return &InitBotDal{
		pool: pool,
	}
}
func (d *InitBotDal) GetPatient(ctx context.Context, chatID int64) (user.Patient, error) {
	sql := `
		select p.* 
		from patients p 
		    join patient_doctor pd on p.id = pd.patient_id 
		where pd.chat_id = $1
		`

	var patient dao.Patient
	err := pgxscan.Get(ctx, d.pool, &patient, sql, chatID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return user.Patient{}, nil
		}
		return user.Patient{}, err
	}

	return patient.ToDomain(), nil
}

func (d *InitBotDal) CreatePatient(ctx context.Context, fullName string) (int64, error) {
	sql := `insert into patients (full_name) values ($1) returning id`

	var id int64
	err := d.pool.QueryRow(ctx, sql, fullName).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("ошибка при сохранении доктора %w", err)
	}

	return id, nil
}

func (d *InitBotDal) CreateM2MPatientDoctor(ctx context.Context, chatID, doctorTgID, patientID int64) error {
	sql := `insert into patient_doctor (doctor_tg, patient_id, chat_id) values ($1, $2, $3)`

	args := []any{
		doctorTgID,
		patientID,
		chatID,
	}

	_, err := d.pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("ошибка при вставке m2m доктора и пациента: %w", err)
	}

	return nil
}

func (d *InitBotDal) UpdatePatientSex(ctx context.Context, patientID int64, sex user.Sex) error {
	sql := `update patients set sex = $1 where id = $2`

	_, err := d.pool.Exec(ctx, sql, sex, patientID)
	if err != nil {
		return err
	}

	return nil
}

func (d *InitBotDal) UpdatePatientBirthDate(ctx context.Context, patientID int64, birthDate time.Time) error {
	sql := `update patients set birth_date = $1 where id = $2`

	_, err := d.pool.Exec(ctx, sql, birthDate, patientID)
	if err != nil {
		return err
	}

	return nil
}

func (d *InitBotDal) UpdatePatientMetricsLink(ctx context.Context, patientID int64, metricsLink string) error {
	sql := `update patients set metrics_link = $1 where id = $2`

	_, err := d.pool.Exec(ctx, sql, metricsLink, patientID)
	if err != nil {
		return err
	}

	return nil
}
