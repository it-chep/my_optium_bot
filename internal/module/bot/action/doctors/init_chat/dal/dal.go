package dal

import (
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto/user"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type InitBotDal struct {
	pool *pgxpool.Pool
}

func NewDal(pool *pgxpool.Pool) *InitBotDal {
	return &InitBotDal{
		pool: pool,
	}
}
func (d *InitBotDal) GetPatient(ctx context.Context, chatID int64) (*user.Patient, error) {
	// todo
}

func (d *InitBotDal) CreatePatient(ctx context.Context, fullName string) error {
	// todo понять бы какой должен быть tg_id скорее пользователя, но тогда надо доктора тут притянуть
	sql := `insert into patients (full_name) values ($1) on conflict (full_name) do nothing`

	_, err := d.pool.Exec(ctx, sql, fullName)
	if err != nil {
		return err
	}

	return nil
}

func (d *InitBotDal) UpdatePatientSex(ctx context.Context, fullName string, sex user.Sex) error {
	sql := `update patients set sex = $1 where full_name = $2`

	_, err := d.pool.Exec(ctx, sql, sex, fullName)
	if err != nil {
		return err
	}

	return nil
}

func (d *InitBotDal) UpdatePatientBirthDate(ctx context.Context, fullName string, birthDate time.Time) error {
	sql := `update patients set birth_date = $1 where full_name = $2`

	_, err := d.pool.Exec(ctx, sql, birthDate, fullName)
	if err != nil {
		return err
	}

	return nil
}

func (d *InitBotDal) UpdatePatientMetricsLink(ctx context.Context, fullName, metricsLink string) error {
	sql := `update patients set metrics_link = $1 where full_name = $2`

	_, err := d.pool.Exec(ctx, sql, metricsLink, fullName)
	if err != nil {
		return err
	}

	return nil
}
