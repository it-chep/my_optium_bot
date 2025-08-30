package dal

import (
	"context"

	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto/education"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/logger"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dal/dao"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type Dal struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Dal {
	return &Dal{
		pool: pool,
	}
}

func (d *Dal) GetStepContent(ctx context.Context, scenarioID, stepID int64) (_ education.Post, err error) {
	sql := `
		select * from contents where scenario_id = $1 and step_id = $2;
	`

	args := []any{
		scenarioID,
		stepID,
	}

	var content dao.Content
	err = pgxscan.Get(ctx, d.pool, &content, sql, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return education.Post{}, nil
		}
		return education.Post{}, err
	}

	return content.ToDomain(), nil
}

func (d *Dal) GetDoctorUsername(ctx context.Context, chatID int64) string {
	sql := `
		select d.tg_username from doctors d
		left join patient_doctor pd on d.tg_id = pd.doctor_tg
		where pd.chat_id = $1
	`

	var doctorUsername string
	err := pgxscan.Get(ctx, d.pool, &doctorUsername, sql, chatID)
	if err != nil {
		logger.Error(ctx, "ошибка получения username доктора по ID", err)
		return ""
	}

	return doctorUsername
}
