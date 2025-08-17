package dal

import (
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dal/dao"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
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

func (d *Dal) GetStepContent(ctx context.Context, scenarioID, stepID int64) (_ dto.Content, err error) {
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
			return dto.Content{}, nil
		}
		return dto.Content{}, err
	}

	return content.ToDomain(), nil
}
