package dal

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dal/dao"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
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
	if err := pgxscan.Select(ctx, d.pool, scenario, `select * from scenario_steps where scenario_id = $1 order by step_order`, id); err != nil {
		return dto.Scenario{}, err
	}

	buttons := &dao.Buttons{}
	if err := pgxscan.Select(ctx, d.pool, scenario, `select * from step_buttons where scenario = $1 order by id`, id); err != nil {
		return dto.Scenario{}, err
	}

	return scenario.ToDomain(steps, buttons), nil
}
